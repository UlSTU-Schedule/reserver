package repository

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository struct {
	config *Config
	db     *sqlx.DB
}

// New ...
func New(config *Config) *Repository {
	return &Repository{
		config: config,
	}
}

// Open ...
func (r *Repository) Open() error {
	db, err := sqlx.Open("pgx", r.config.DatabaseURL)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	r.db = db
	return nil
}

// Close ...
func (r *Repository) Close() error {
	return r.db.Close()
}
