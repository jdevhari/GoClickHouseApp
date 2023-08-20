package repository

import (
	"database/sql"
	"log"

	"github.com/jdevhari/GoClickHouseApp/pkg/config"
	_ "github.com/mailru/go-clickhouse/v2"
)

func DropTables() {
	connect, err := sql.Open("chhttp", config.DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = connect.Exec(`
		DROP TABLE IF EXISTS ticks
	`)

	_, err = connect.Exec(`
		DROP TABLE IF EXISTS ticks_5min
	`)

	_, err = connect.Exec(`
		DROP TABLE IF EXISTS ticks_15min
	`)

	_, err = connect.Exec(`
		DROP TABLE IF EXISTS ticks_60min
	`)

	_, err = connect.Exec(`
		DROP VIEW IF EXISTS ticks_mv5
	`)

	_, err = connect.Exec(`
		DROP VIEW IF EXISTS ticks_mv15
	`)

	_, err = connect.Exec(`
		DROP VIEW IF EXISTS ticks_mv60
	`)

	
	if err != nil {
		log.Fatal(err)
	}


}