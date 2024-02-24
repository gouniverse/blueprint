package config

// import (
// 	"database/sql"
// 	"time"

// 	"github.com/cenkalti/backoff/v4"
// )

// // VerifyConnection pings the database to verify a connection is established. If the connection cannot be established,
// // it will retry with an exponential back off.
// func verifyConnection(c *sql.DB) error {
// 	pingDB := func() error {
// 		return c.Ping()
// 	}

// 	expBackoff := backoff.NewExponentialBackOff()
// 	expBackoff.MaxElapsedTime = time.Duration(30) * time.Second

// 	err := backoff.Retry(pingDB, expBackoff)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
