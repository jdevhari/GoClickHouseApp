package repository

import (
	"database/sql"
	"log"
	"time"
	"math/rand"

	"github.com/jdevhari/GoClickHouseApp/pkg/config"
	_ "github.com/mailru/go-clickhouse/v2"
)

func InsertData() {
	connect, err := sql.Open("chhttp", config.DB_URL)
	if err != nil {
		log.Print(err)
	}
	if err := connect.Ping(); err != nil {
		log.Print(err)
	}

	tx, err := connect.Begin()
	if err != nil {
		log.Print(err)
	}
	stmt, err := tx.Prepare(`
		INSERT INTO default.ticks (
			timestamp, symbol, open, high, low, close
		) VALUES ( ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Print(err)
	}

	tickTimeStamp := time.Date(2023, time.Month(8), 1, 9, 15, 0, 0, time.UTC)
	var open = 19000
	for i := 0; i < 375; i++ {
		var low = open - rand.Intn(200)
		var high = low + rand.Intn(500)
		var close = 0;
		if high%2 == 0{
			close = 19000 + rand.Intn(700)
		} else {
			close = 19000 - rand.Intn(700)
		}
		if _, err := stmt.Exec(
			tickTimeStamp,
			"NIFTY",
			open,
			high,
			low,
			close,
		); err != nil {
			log.Print(err)
		}
		tickTimeStamp = tickTimeStamp.Add(time.Minute)
		open = close
	}

	if err := tx.Commit(); err != nil {
		log.Print("Commit Error")
		log.Print(err)
	}
	log.Print("Data populated in ticks table")

	_, err = connect.Exec(`
		INSERT INTO default.ticks_5min
		SELECT
		toStartOfFiveMinute(timestamp) AS timestamp,
		symbol,
		argMin(open, ticks.timestamp) as open,
		argMax(high, high) as high,
		argMin(low, low) as low,
		argMax(close, ticks.timestamp) as close
		FROM ticks
		GROUP BY symbol, timestamp
		ORDER BY symbol, timestamp`)

	if err != nil {
		log.Print(err)
	}
	log.Print("Data populated in ticks5 table")

	_, err = connect.Exec(`
	INSERT INTO  default.ticks_15min
		SELECT
		toStartOfFifteenMinutes(timestamp) AS timestamp,
		symbol,
		argMin(open, ticks.timestamp) as open,
		argMax(high, high) as high,
		argMin(low, low) as low,
		argMax(close, ticks.timestamp) as close
		FROM ticks
		GROUP BY symbol, timestamp
		ORDER BY symbol, timestamp`)

	if err != nil {
		log.Print(err)
	}
	log.Print("Data populated in ticks15 table")

	_, err = connect.Exec(`
	INSERT INTO default.ticks_60min
		SELECT
		toStartOfInterval(timestamp, INTERVAL 60 minute) AS timestamp,
		symbol,
		argMin(open, ticks.timestamp) as open,
		argMax(high, high) as high,
		argMin(low, low) as low,
		argMax(close, ticks.timestamp) as close
		FROM ticks
		GROUP BY symbol, timestamp
		ORDER BY symbol, timestamp`)

	if err != nil {
		log.Print(err)
	}
	log.Print("Data populated in ticks60 table")
}

