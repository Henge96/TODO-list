package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(cfg Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s username=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
	logrus.Info(dsn)
	db, err := sqlx.Connect(cfg.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open: %w", err)
	}

	return db, nil
}
