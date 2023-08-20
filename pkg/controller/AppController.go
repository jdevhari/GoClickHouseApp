package controller

import "github.com/jdevhari/GoClickHouseApp/pkg/model"
import "github.com/jdevhari/GoClickHouseApp/pkg/config"
import "github.com/jdevhari/GoClickHouseApp/pkg/repository"
import "github.com/jdevhari/GoClickHouseApp/pkg/repository/mock"

func GetData(queryParams map[string]string) (model.TickData, error) {
	if config.MOCK_DATA {
		return mock.Query(queryParams)
	} else {
		return repository.Query(queryParams)
	}
}

func CreateTables() {
	repository.CreateTables()
}

func InsertData() {
	repository.InsertData()
}

func DropTables() {
	repository.DropTables()
}


