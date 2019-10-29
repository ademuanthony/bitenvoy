package main

import (
	"fmt"
	"github.com/ademuanthony/bitenvoy/orders/handler"
	"os"

	"github.com/ademuanthony/bitenvoy/orders/postgres"
	orders "github.com/ademuanthony/bitenvoy/orders/proto/orders"
	"github.com/jessevdk/go-flags"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	cfg, _, err := LoadConfig()
	if err != nil{
		log.Fatal(err)
	}

	db, err := postgres.NewPgDb(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// create tables
	if !db.OrderTableExists() {
		if err = db.CreateOrderTable(); err != nil {
			log.Fatal(err)
		}
	}

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.orders"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	if err := orders.RegisterOrdersHandler(service.Server(), handler.NewOrderHandler(db)); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
const (
	DefaultConfigFilename      = "airtime.conf"
	Hint                       = `Run dcrextdata < --http > to start http server or dcrextdata < --help > for help.`
	defaultDbHost              = "localhost"
	defaultDbPort              = "5432"
	defaultDbUser              = "postgres"
	defaultDbPass              = "dbpass"
	defaultDbName              = "orders"
	defaultLogLevel            = "debug"
)

func defaultConfig() Config {
	cfg := Config{
		DBHost:          defaultDbHost,
		DBPort:          defaultDbPort,
		DBUser:          defaultDbUser,
		DBPass:          defaultDbPass,
		DBName:          defaultDbName,
		DebugLevel:      defaultLogLevel,
	}

	return cfg
}

type Config struct {
	DebugLevel string `short:"d" long:"debuglevel" description:"Logging level {trace, debug, info, warn, error, critical}"`
	Quiet      bool   `short:"q" long:"quiet" description:"Easy way to set debuglevel to error"`

	// Postgresql Configuration
	DBHost string `long:"dbhost" description:"Database host"`
	DBPort string `long:"dbport" description:"Database port"`
	DBUser string `long:"dbuser" description:"Database username"`
	DBPass string `long:"dbpass" description:"Database password"`
	DBName string `long:"dbname" description:"Database name"`

}

func LoadConfig() (*Config, []string, error) {
	cfg := defaultConfig()
	parser := flags.NewParser(&cfg, flags.IgnoreUnknown)
	err := flags.NewIniParser(parser).ParseFile(DefaultConfigFilename)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			fmt.Printf("Missing Config file %s in current directory\n", DefaultConfigFilename)
		} else {
			return nil, nil, err
		}
	}

	unknownArg, err := parser.Parse()
	if err != nil {
		e, ok := err.(*flags.Error)
		if ok && e.Type == flags.ErrHelp {
			os.Exit(0)
		}
		return nil, nil, err
	}

	return &cfg, unknownArg, nil
}
