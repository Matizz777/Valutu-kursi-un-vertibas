package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	dsn := url.URL{
		Scheme: "valutaKurss",
		Host:   "localhost:3306",
		User:   url.UserPassword("user", "password"),
		Path:   "valutaKurss",
	}
	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("valutaKurss", dsn.String())
	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}
	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()
	if err := db.PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return
	}

	row := db.QueryRowContext(context.Background(), "SELECT title, description, pubDate from channel where pubdate like %09 Nov 2022%")

	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return
	}
	var atjaunots time.Time

	if err := row.Scan(&atjaunots); err != nil {
		fmt.Println("row.Scan", err)
		return
	}
	fmt.Println("atjaunots", atjaunots)
}
