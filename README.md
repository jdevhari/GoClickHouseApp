# GoClickHouseApp

## Table of Contents
- [Scope](#Scope)
- [Technology Stack](#Technology-Stack)
- [Quick Start & Demo](#Quick-Start-&-Demo)
- [Full Development setup](#Full-Development-setup)
- [Screenshots](#Screenshots)
- [Next Steps](#Next-Steps)
- [References](#References)

## Scope
- Setup a Clickhouse repository. Populate NIFTY tickdata for a trading day.
- Develop backend Apis using GoLang & GoFiber
- Render NIFTY Tick Data as Candlesticks at 5min, 15min or 1 Hour interval fetching data from Clickhouse Datastore and massaging data in GoLang Api

## Technology Stack
- [Clickhouse](https://clickhouse.com/) as Repository
- [Go](https://go.dev/) with [GoFiber](https://gofiber.io/) as Api Stack
- [Angular](https://angular.io/) as UI framework and [Apache ECharts](https://echarts.apache.org/) as Charting Component

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
[!NOTE]
The NIFTY tickdata data has been generated randomly. Dont be alarmed by the Volatality :blush:

### 1 Min Timeframe
![1 Min TimeFrame.](/assets/1.jpg)

### 5 Min Timeframe
![5 Min TimeFrame.](/assets/5.jpg)

### 15 Min Timeframe
![15 Min TimeFrame.](/assets/15.jpg)

### 1 Hr Timeframe
![1 Hr TimeFrame.](/assets/60.jpg)

## Next Steps
- Deploy to AWS

## References
- [Using ClickHouse for financial market data - Christoph Wurm (ClickHouse)](https://www.youtube.com/watch?v=Ojv6LPXKy2U&ab_channel=ClickHouse)
- [ClickHouse for Financial Analytics - Baudouin Giard (Bloomberg)](https://www.youtube.com/watch?v=HmJTIrGyVls&t=188s&ab_channel=ClickHouse)
- [Introduction to High-Velocity Analytics Using ClickHouse Arrays | ClickHouse Webinar | Altinity](https://www.youtube.com/watch?v=hAzrhKhNqds&ab_channel=Altinity)
