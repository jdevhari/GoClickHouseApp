package repository

import (
	"database/sql"
	"log"

	"github.com/jdevhari/GoClickHouseApp/pkg/config"
	_ "github.com/mailru/go-clickhouse/v2"
)

func CreateTables() {
	connect, err := sql.Open("chhttp", config.DB_URL)
	if err != nil {
		log.Print(err)
	}
	if err := connect.Ping(); err != nil {
		log.Print(err)
	}

	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS ticks
		(
			timestamp DateTime64,
			symbol LowCardinality(String),
			open Nullable(Float64),
			high Nullable(Float64),
			low Nullable(Float64),
			close Nullable(Float64)
		)
			ENGINE = MergeTree
			ORDER BY (symbol, timestamp)
	`)

	if err != nil {
		log.Print(err)
	}
	
	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS ticks_5min
		(
			timestamp DateTime64 CODEC(Delta, Default),
			symbol LowCardinality(String),
			open Float64,
			high Float64,
			low Float64,
			close Float64
		)
		ENGINE = AggregatingMergeTree
		ORDER BY (symbol, timestamp)`)

	if err != nil {
		log.Print(err)
	}

	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS ticks_15min
		(
			timestamp DateTime64 CODEC(Delta, Default),
			symbol LowCardinality(String),
			open Float64,
			high Float64,
			low Float64,
			close Float64
		)
		ENGINE = AggregatingMergeTree
		ORDER BY (symbol, timestamp)`)

	if err != nil {
		log.Print(err)
	}

	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS ticks_60min
		(
			timestamp DateTime64 CODEC(Delta, Default),
			symbol LowCardinality(String),
			open Float64,
			high Float64,
			low Float64,
			close Float64
		)
		ENGINE = AggregatingMergeTree
		ORDER BY (symbol, timestamp)`)

	if err != nil {
		log.Print(err)
	}

	log.Print("Table(s) created")	
}

