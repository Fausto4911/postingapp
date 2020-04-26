package dbconnection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // here
)

/*
func OpenDB(c driver.Connector) *DB

db, err := sql.Open(driver, dataSourceName)

to test the connection
if err := db.PingContext(ctx); err != nil {
  log.Fatal(err)
}

var conninfo string = ""
        db, err := sql.Open("postgres", conninfo)
        if err != nil {
            panic(err)
        }

*/
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "POSTING"
)

func GetPostingDb() string {

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

	fmt.Println("Successfully connected!")

	// fmt.Println("Getting posting db connection...")
	// // Connect to the DB, panic if failed
	// db, err := sql.Open("postgres", "postgres://main:postgres@localhost:5432/POSTING?sslmode=disable")
	// if err != nil {
	// 	fmt.Println(`Could not connect to db`)
	// 	panic(err)
	// }
	// defer db.Close()

	return "cz"
}
