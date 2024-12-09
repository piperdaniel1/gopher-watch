package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/piperdaniel1/gopher-watch/server/config"
)

func GetDBConnection(cfg config.DBConfig) (*sql.DB, error) {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.GetDBHost(), cfg.GetDBPort(), cfg.GetDBUser(), cfg.GetDBPassword(), cfg.GetDBName(),
	)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Connected to the database!")
	return db, nil
}
