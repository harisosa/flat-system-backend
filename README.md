# flat-system-backend

User Guide

# Dependencies
## 1. Install Go Runtime
    https://go.dev/doc/install
## 2. Install Postgresql


## Setting Environtment
    copy the file `copy.env` value and make `.env` file

    | Parameter     | Explanation                                                                               | example          |
    | --------------| ------------------------------------------------------------------------------------------|------------------|
    | DB_HOST       | where the database hosted                                                                 | localhost        |
    | DB_PORT       | database port                                                                             | 5432             |
    | DB_NAME       | database name                                                                             | flat-system      |
    | DB_USERNAME   | database credential - username                                                            | root             |
    | DB_PASSWORD   | database credential - password                                                            | root             |
    | DEBUG         | debug will activate the log and will seeding data into database, for initiation purpose   | true             |

## How to run the Project
1. Go Mod vendor
2. Go run main.go
