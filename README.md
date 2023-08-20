# GoClickHouseApp

## Table of Contents
- [Scope](#Scope)
- [Technology Stack](#Technology-Stack)
- [Quick Start & Demo](#Quick-Start-&-Demo)
- [Full Development setup](#Full-Development-setup)
- [Screenshots](#Screenshots)
- [Next Steps](#Next-Steps)

## Scope
- Setup a Clickhouse instance. Populate NIFTY tickdata for a trading day
- Develop backend Api using GoLang & GoFiber
- In web page render NIFTY Tick Data at 5min, 15min or 1 Hour interval fetching data from Clickhouse Datastore and massaging data in GoLang Api

## Technology Stack
- Clickhouse as Repository
- Go with GoFiber as Api Stack
- Angular as UI framework and Apache ECharts as Charting Component

## Quick Start & Demo
- Ensure you have Go Installed
- Goto the folder `cmd` and execute the command `go run main.go`
- Configure a Data Source of your choice at `pkg/config/serverconfig.go`
  - (Default) Mock Data from `pkg/repository/mock/mockquery.go`
  - Run a local docker image. Refer to [Repository Setup](DB.md)
  - Point to existing Click house instance
- Access the application at http://localhost:3000

## Full Development setup
- Refer to [Development Setup](DEVELOPMENT.md)

## Screenshots

### 1 Min Timeframe
![1 Min TimeFrame.](/assets/1.jpg)

### 5 Min Timeframe
![5 Min TimeFrame.](/assets/5.jpg)

### 15 Min Timeframe
![15 Min TimeFrame.](/assets/15.jpg)

### 1 Hr Timeframe
![1 Hr TimeFrame.](/assets/60.jpg)

## Next Steps
