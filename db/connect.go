package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.con/binsabi/go-blog/config"
)

func ConnectToDB(storageConf config.Storage) *sql.DB {
	db, err := sql.Open(storageConf.DBDriver, storageConf.DSN)
	if err != nil {
		log.Fatalf("couldnot connect to db: %v", err)
	}

	db.SetMaxIdleConns(storageConf.MaxIdleConns)
	db.SetMaxOpenConns(storageConf.MaxOpenConns)
	db.SetConnMaxIdleTime(storageConf.MaxIdleTime)
	db.SetConnMaxLifetime(storageConf.MaxConnLife)

	if err := db.Ping(); err != nil {
		log.Fatalf("couldnot ping the db: %v", err)
	}
	return db
}
