package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/diyliv/interweb/config"
)

func ConnPostgres(cfg *config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.Login,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DB)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Minute * time.Duration(cfg.Postgres.ConnMaxLifeTime))
	conn.SetMaxOpenConns(cfg.Postgres.MaxOpenConn)
	conn.SetMaxIdleConns(cfg.Postgres.MaxIdleConn)

	return conn, nil
}
