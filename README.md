Stock Trades API

A RESTful API built using Go, PostgreSQL, and JWT for managing stock trades. This project was developed for the Liquide BE Hackathon by Ruchir Nayak

Prerequisites
Go (1.16 or later)
PostgreSQL
pgAdmin 4
Postman (for testing)

bash
git clone https://github.com/ruch1029/stock-trades-api.git
cd stock-trades-api
go mod init stock-trades-api

Install dependencies:

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v4

Database Setup
Open pgAdmin 4 and create a new database. You can name it stock_trades_db.

Set up your PostgreSQL database connection and change the values in config.go

You can use the following quearies in your db 

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE trades (
    id SERIAL PRIMARY KEY,
    type VARCHAR(10) NOT NULL,
    user_id INT NOT NULL REFERENCES users(id),
    symbol VARCHAR(10) NOT NULL,
    shares INT NOT NULL,
    price INT NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL
);
