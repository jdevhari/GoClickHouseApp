# Dev Machine Setup

[Back To Main](README.md)

## Pre Requisites
- Install [Go](https://go.dev/dl/)
- Install [Docker](https://docs.docker.com/get-docker/)
- Install NVM, Node, npm

## Checkout
- Checkout the Repo from https://github.com/jdevhari/GoClickHouseApp.git
  
## API Layer
### Setup
- `cd cmd` and execute ```go run main.go ```
- Access the application at http://localhost:3000
- By default the packaged UI layer in ui/dist folder is served

### Build
- To Build the API app, execute `go build main.go`.
- To run the API app in dev mode, execute `go run main.go`.

## UI Layer
### Setup
- `cd GoClickHouseApp` and then `cd ui`
- Execute the following command `npm install ` to install the JS dependencies
- Goto ui folder and execute `npm run start `. The Reloadable UI app is served at `http://localhost:4200`
- For backend api, the UI app connects to apis at `http://localhost:3000`, served by GoFiber
- Alternatively, the PreBuilt UI application can be accessed at `http://localhost:3000`

### Build
- To Build the UI app, execute `npm run build`.
