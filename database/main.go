package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"
)

func GetSQLite() *sql.DB {
	db, err := sql.Open("sqlite3", "./dev.db")

	if err != nil {
		panic(err)
	}

	return db
}

func GetRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
