package model

type TickData struct {
	InstrumentSymbol string `json:"instrumentSymbol`
	DateTime []string `json:"dateTime"`
	TickData [][]int64 `json:"tickData"`
  }
  