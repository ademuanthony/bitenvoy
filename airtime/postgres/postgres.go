package postgres

//go:generate sqlboiler --wipe psql --no-hooks --no-auto-timestamps

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ademuanthony/bitenvoy/utils"
	"github.com/micro/go-micro/util/log"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"

	"github.com/ademuanthony/bitenvoy/airtime/postgres/models"
	"github.com/ademuanthony/bitenvoy/airtime/proto/airtime"
	"github.com/volatiletech/sqlboiler/boil"
)

type PgDb struct {
	db *sql.DB
	queryTimeout time.Duration
}

func NewPgDb(host, port, user, pass, dbname string) (*PgDb, error) {
	db, err := utils.PgConnect(host, port, user, pass, dbname)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(5)
	return &PgDb{
		db: db,
		queryTimeout: time.Second * 30,
	}, nil
}

func (pg *PgDb) Close() error {
	log.Trace("Closing postgresql connection")
	return pg.db.Close()
}

func (pg *PgDb) AddCountry(ctx context.Context, name string, id string) error {
	country := models.Country{
		ID:   id,
		Name: name,
	}

	if exists, _ := models.Countries(models.CountryWhere.Name.EQ(name)).Exists(ctx, pg.db); exists {
		return fmt.Errorf("a country with the name %s already exists", name)
	}

	if err := country.Insert(ctx, pg.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (pg *PgDb) Countries(ctx context.Context) ([]*airtime.Country, error) {
	countries, err := models.Countries().All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var result []*airtime.Country
	for _, country := range countries {
		result = append(result, &airtime.Country{
			Id:                   country.ID,
			Name:                 country.Name,
		})
	}

	return result, nil
}

func (pg *PgDb) AddProvider(ctx context.Context, provider airtime.Provider) error {
	tx, err := pg.db.Begin()
	if err != nil {
		return fmt.Errorf("cannot begin a new trasaction for provider insertion, %s", err.Error())
	}

	if exist, _ := models.Providers(models.ProviderWhere.Name.EQ(provider.Name)).Exists(ctx, tx); exist {
		return fmt.Errorf("a provider with the name, %s exists", provider.Name)
	}

	providerModels := models.Provider{
		ID:       provider.Id,
		Name:     provider.Name,
		Logo:     provider.Logo,
		Status:   int(airtime.ProviderStatus_value[provider.Status.String()]),
		Discount: provider.Discount,
	}

	if err := providerModels.Insert(ctx, tx, boil.Infer()); err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			err = fmt.Errorf("cannot add provider because, %s and cannot rollback because, %s", err.Error(), rollBackErr.Error())
		}
		return err
	}

	countryProvider := models.ProviderCountry{
		ProviderID: providerModels.ID,
		CountryID:  provider.CountryId,
	}

	if err := countryProvider.Insert(ctx, tx, boil.Infer()); err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			err = fmt.Errorf("cannot add provider because, %s and cannot rollback because, %s", err.Error(), rollBackErr.Error())
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("cannot commit the transaction for adding provider, %s", err.Error())
	}

	return nil

}

func (pg *PgDb) AddProviderCountry(ctx context.Context, providerId string, countryId string) error {
	countryProvider := models.ProviderCountry{
		ProviderID: providerId,
		CountryID:  countryId,
	}

	if err := countryProvider.Insert(ctx, pg.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (pg *PgDb) Providers(ctx context.Context, countryId string) ([]*airtime.Provider, error) {
	providerSlice, err := models.Providers(/*models.ProviderWhere.ID.IN([]string{})*/).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var providers []*airtime.Provider
	for _, provider := range providerSlice {
		providers = append(providers, &airtime.Provider{
			Id:                   provider.ID,
			CountryId:            provider.ID,
			Name:                 provider.Name,
			Logo:                 provider.Logo,
			Status:               airtime.ProviderStatus(provider.Status),
			Discount:             provider.Discount,
		})
	}

	return providers, nil
}

func (pg *PgDb) Histories(ctx context.Context, username string, offset int32, maxResultCount int32) ([]*airtime.History, int64, error) {
	var queries []qm.QueryMod
	if username != "" {
		queries = append(queries, models.HistoryWhere.Username.EQ(username))
	}

	totalCount, err := models.Histories(queries...).Count(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	queries = append(queries, qm.Offset(int(offset)))
	queries = append(queries, qm.Limit(int(maxResultCount)))

	historySlice, err := models.Histories(queries...).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	var histories []*airtime.History
	for _, historyModel := range historySlice {
		histories = append(histories, &airtime.History{
			Username:    historyModel.Username,
			PhoneNumber: historyModel.PhoneNumber,
			Network:     historyModel.Network,
			Amount:      historyModel.Amount,
			Date:        int64(historyModel.Date),
		})
	}

	return histories, totalCount, nil
}

func (pg *PgDb) HistoryCount(ctx context.Context, username string) (int64, error) {
	var queries []qm.QueryMod
	if username != "" {
		queries = append(queries, models.HistoryWhere.Username.EQ(username))
	}

	return models.Histories(queries...).Count(ctx, pg.db)
}
