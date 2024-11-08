package config

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq" // Enable PostgreSQL driver if needed
	_ "modernc.org/sqlite"
)

// openDb opens the database
//
// Business logic:
//   - opens the database based on the driver name
//   - each driver has its own set of parameters
//
// Parameters:
// - driverName: the driver name
// - dbHost: the database host
// - dbPort: the database port
// - dbName: the database name
// - dbUser: the database user
// - dbPass: the database password
//
// Returns:
// - *sql.DB: the database connection
// - error: the error if any
func openDb(driverName string, dbHost string, dbPort string, dbName string, dbUser string, dbPass string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	if driverName == "sqlite" {
		dsn := dbName
		db, err = sql.Open("sqlite", dsn)
	}

	if driverName == "mysql" {
		dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=UTC"
		db, err = sql.Open("mysql", dsn)
		// Maximum Idle Connections
		db.SetMaxIdleConns(5)
		// Maximum Open Connections
		db.SetMaxOpenConns(5)
		// Idle Connection Timeout
		db.SetConnMaxIdleTime(5 * time.Second)
		// Connection Lifetime
		db.SetConnMaxLifetime(30 * time.Second)
	}

	if driverName == "postgres" {
		dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=UTC"
		db, err = sql.Open("postgres", dsn)
	}

	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, errors.New("database for driver " + driverName + " could not be intialized")
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatalln("Db Ping Failed: ", errPing.Error())
	}

	// if err := VerifyConnection(db); err != nil {
	// 	log.Fatalln("[Error] Db Ping Failed:", err)
	// }

	return db, nil
}
