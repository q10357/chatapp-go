package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	DBUser string
	DBPass string
	DBName string
}

func InitDatabase(cfg *Config) (*sql.DB, error) {
	dbCfg := mysql.Config{
		User:   cfg.DBUser,
		Passwd: cfg.DBPass,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: cfg.DBName,
	}

	db, err := sql.Open("mysql", dbCfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected!")
	return db, nil
}
