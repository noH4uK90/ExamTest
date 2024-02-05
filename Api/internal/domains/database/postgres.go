package database

import (
	"Api/internal/config"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgresConnection(config *config.PostgresConfig) (*sqlx.DB, error) {
	cfg := pgx.ConnConfig{
		PreferSimpleProtocol: true,
		Host:                 config.Host,
		Port:                 config.Port,
		User:                 config.User,
		Password:             config.Password,
		Database:             config.Database,
	}

	db := stdlib.OpenDB(cfg)
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return sqlx.NewDb(db, "pgx"), nil
}
