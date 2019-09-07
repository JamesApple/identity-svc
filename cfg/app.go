package cfg

import (
	"log"
	"os"
	"root/data"
	"root/data/postgres"
)

type App struct {
	Config      *Config
	AccountRepo data.AccountRepo
	Migrate     func() error
	Logger      *log.Logger
}

func NewApp(cfg *Config) (*App, error) {
	app := &App{}

	db, err := postgres.Connect(cfg.postgresURL)
	if err != nil {
		return nil, err
	}

	app.Logger = log.New(os.Stdout, "App", 0)

	app.AccountRepo = &postgres.AccountRepo{
		InstrumentedDB: &postgres.InstrumentedDB{
			DB:     db,
			Logger: app.Logger,
		},
	}

	app.Migrate = func() error {
		return postgres.Migrate(db)
	}

	return app, nil
}
