package repository

import (
	"database/sql"
	"log"
	"time"
	"fmt"

	"github.com/jdevhari/GoClickHouseApp/pkg/model"
	"github.com/jdevhari/GoClickHouseApp/pkg/config"
	_ "github.com/mailru/go-clickhouse/v2"
)

func Query(queryParams map[string]string) (model.TickData,error) {
	var tickData model.TickData
	connect, err := sql.Open("chhttp", config.DB_URL)
	if err != nil {
		log.Print(err)
		return tickData, err
	}
	if err := connect.Ping(); err != nil {
		log.Print(err)
		return tickData, err
	}

	tickData.TickData = make([][]int64, 0) 
	var timeFrame = queryParams["timeFrame"]
	var tableName = "ticks"
	if timeFrame == "5" {
		tableName = "ticks_5min"
	} else if timeFrame == "15" {
		tableName = "ticks_15min"
	} else if timeFrame == "60" {
		tableName = "ticks_60min"
	}
	var query = `
	SELECT timestamp,
		symbol,
		open,
		high,
		low,
		close
	FROM
		%s`
	rows, err := connect.Query(fmt.Sprintf(query, tableName))

	if err != nil {
		log.Print(err)
		return tickData, err
	}

	tickData.InstrumentSymbol = "NIFTY"
	for rows.Next() {
		var (
			timestamp               time.Time
			symbol           		string
			open, high, low, close 	int64
		)
		if err := rows.Scan(
			&timestamp,
			&symbol,
			&open,
			&high,
			&low,
			&close,
		); err != nil {
			log.Print(err)
			return tickData, err
		}

		hours, minutes, _ := timestamp.Clock()
    	hourMinString := fmt.Sprintf("%d:%02d", hours, minutes)

		tickData.DateTime = append(tickData.DateTime, hourMinString)

        sliceDataIdx := make([]int64, 0)
		sliceDataIdx = append(sliceDataIdx, close, open, low, high)

        tickData.TickData = append(tickData.TickData, [][]int64{sliceDataIdx}...)
	}
	return tickData, nil
}