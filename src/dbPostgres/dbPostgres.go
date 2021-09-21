package dbPostgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
)

const (
	driverName = "postgres"
	dbUser     = "bond"
	dbPassword = "password"
	dbPort     = 5432
	dbName     = "tokens_db"
	dbSSLmode  = "disable"
)

var (
	driver   = driverName
	user     = dbUser
	password = dbPassword
	port     = dbPort
	name     = dbName
	sslmode  = dbSSLmode
)

func init() {
	dataSourceName := fmt.Sprintf("user=%s port=%d dbname=%s password=%s sslmode=%s",
		user, port, name, password, sslmode)
	var err error
	Client, err = sql.Open(driver, dataSourceName)
	if err != nil {
		fmt.Println("Database Connectivity Failed")
		panic(err)
	}
	pingErr := Client.Ping()
	if pingErr != nil {
		fmt.Println("ping to database failed")
		panic(pingErr)
	}
	fmt.Println("Database Connected Successfully")
}
