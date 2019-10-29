package handler

import (
	"context"
	"fmt"
	"github.com/ademuanthony/bitenvoy/airtime/proto/airtime"
	"github.com/gofrs/uuid"
)

// DataStore provider functions for database operation
type DataStore interface {
	AddCountry(ctx context.Context, name string, id string) error
	Countries(ctx context.Context) ([]*airtime.Country, error)
	AddProvider(ctx context.Context, provider airtime.Provider) error
	AddProviderCountry(ctx context.Context, providerId string, countryId string) error
	Providers(ctx context.Context, countryId string) ([]*airtime.Provider, error)
	Histories(ctx context.Context, username string, offset int32, maxResultCount int32) ([]*airtime.History, int64, error)
	HistoryCount(ctx context.Context, username string) (int64, error)
}

type airtimeHandler struct{
	store DataStore
}

func NewAirtime(store DataStore) *airtimeHandler {
	return &airtimeHandler{
		store: store,
	}
}

func (a airtimeHandler) AddCountry(ctx context.Context, req *airtime.AddCountryRequest, res *airtime.AddCountryResponse) error {
	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("cannot generate uniqe if for country, %s", err.Error())
	}
	err = a.store.AddCountry(ctx, req.Name, id.String())
	if err != nil {
		return err
	}

	res.Id = id.String()
	return nil
}

func (a airtimeHandler) Countries(ctx context.Context, _ *airtime.GetCountriesRequest, res *airtime.GetCountriesResponse) error {
	countries, err := a.store.Countries(ctx)
	res.Countries = countries

	return err
}

func (a airtimeHandler) AddProviderToCountry(ctx context.Context, req *airtime.AddProviderToCountyRequest, res *airtime.AddProviderToCountryResponse) error {
	if err := a.store.AddProviderCountry(ctx, req.ProviderId, req.CountryId); err != nil {
		return err
	}

	return nil
}

func (a airtimeHandler) AddProvider(ctx context.Context, req *airtime.AddProviderRequest, res *airtime.AddProviderResponse) error {
	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("cannot generate uuid for new provider, %s", req.Name)
	}
	provider := airtime.Provider{
		Id:                   id.String(),
		CountryId:            req.CountryId,
		Name:                 req.Name,
		Logo:                 req.Logo,
		Status:               req.Status,
		Discount:             req.Discount,
	}

	if err = a.store.AddProvider(ctx, provider); err != nil {
		return err
	}

	res.ProviderId = id.String()
	return nil
}

func (a airtimeHandler) Providers(ctx context.Context, req *airtime.ProvidersRequest, res *airtime.ProvidersResponse) error {
	providers, err := a.store.Providers(ctx, req.CountryId)
	if err != nil {
		return err
	}

	res.Providers = providers
	return nil
}

func (a airtimeHandler) SendAirtime(context.Context, *airtime.SendAirtimeRequest, *airtime.SendAirtimeResponse) error {
	return nil
}

func (a airtimeHandler) History(ctx context.Context, req *airtime.HistoryRequest, resp *airtime.HistoryResponse) error {
	histories, totalCount, err := a.store.Histories(ctx, req.Username, req.SkipCount, req.MaxResultCount)
	resp.Histories = histories
	resp.TotalCount = totalCount
	return err
}

func (a airtimeHandler) HistoryCount(ctx context.Context, req *airtime.HistoryCountRequest, resp *airtime.HistoryCountResponse) error {
	count, err := a.store.HistoryCount()
	resp.Count = count
	return err
}

