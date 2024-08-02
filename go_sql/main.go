package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@(ip:port)/database")

	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	row := db.QueryRowContext(ctx, "SELECT * FROM t limit 1")

	if row.Err() != nil {
		return
	}
	type user struct {
		Id string
	}
	var u user
	if err = row.Scan(&u); err != nil {
		return
	}
}
