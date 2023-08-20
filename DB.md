# Database

[Back To Main](README.md)

## Pre Requisites
- Install Docker

## Docker Setup
- Execute the following command
  ``` docker pull clickhouse/clickhouse-server ```
- Start the Docker image. Make sure ports 8123, 9000 are exposed.
- From browser hit http://localhost:8123 If you get ok, then the setup is fine

## Data Setup
 - Create tables by accessing http://localhost:3000/createTables
 - Delete tables by accessing http://localhost:3000/dropTables
 - Populate tables by accessing http://localhost:3000/insertData

## Click House Client Setup
- From browser hit [Tabix Login](http://ui.tabix.io/#!/login)
- login with username as default and no password
- Once logged on to console, use the same for firing sql
