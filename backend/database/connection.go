package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


func NewDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", "db", "postgres", "password", "golang_postgres", "5432")
	return sql.Open("postgres", dsn)
}
