package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

var db *bun.DB

func ConnectDatabase(isProduction bool) {
	var sqldb *sql.DB
	var err error
	if isProduction {
		sqldb, err = sql.Open(sqliteshim.ShimName, "file:/var/www/inkdrop_quan_usr/data/www/inkdrop.quanthai.net/web-design/inkdrop/database.db?cache=shared&mode=rwc")
	} else {
		sqldb, err = sql.Open(sqliteshim.ShimName, "file:database.db?cache=shared&mode=rwc")
	}
	if err != nil {
		panic(err)
	}
	if err = sqldb.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	fmt.Println("Database connected. We're cooking.")
	db = bun.NewDB(sqldb, sqlitedialect.New())
}

func GetDB() *bun.DB {
	return db
}
