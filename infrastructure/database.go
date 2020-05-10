package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	ctx context.Context
	db  *sql.DB
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "POSTING"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	id := 47
	result, err := db.ExecContext(ctx, "UPDATE balances SET balance = balance + 10 WHERE user_id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rows != 1 {
		fmt.Println("expected to affect 1 row, affected %d", rows)
	}

	fmt.Println("Successfully connected!")
}
