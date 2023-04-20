package dbrepo

import (
	"database/sql"

	"github.com/gauravpatil28/booking/internal/config"
	"github.com/gauravpatil28/booking/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {

	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}

}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {

	return &testDBRepo{
		App: a,
	}
}

