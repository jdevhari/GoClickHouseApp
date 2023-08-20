# Dev Machine Setup

[Back To Main](README.md)

## Pre Requisites
- Install Go
- Install Docker
- Install NVM, Node, npm

## Setup
- Checkout the Repo from https://github.com/jdevhari/GoClickHouseApp.git
- cd ui
- Execute the following command
  ``` npm install ```
  
## Backend Setup
- cd cmd folder and execute ```go run main.go ```
- Access the application at http://localhost:3000
- By default the packaged UI layer in ui/dist folder is served

## UI development Setup
- cd to ui folder and execute ```npm run start ```
- For backend api, the UI layer connects to http://localhost:3000
- Access the Reloadable UI application at http://localhost:4200
- Access the PreBuilt   UI application at http://localhost:3000
