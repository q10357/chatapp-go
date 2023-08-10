package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	DBUser string
	DBPass string
	DBName string
}

func InitDatabase(cfg *Config) (*sql.DB, error) {
	dbCfg := mysql.Config{
		User:   cfg.DBUser,
		Passwd: cfg.DBPass,
		Net:    "tcp",
		Addr:   "rel_db:3306",
		DBName: cfg.DBName,
	}

	var db *sql.DB
	var err error
	
	for i := 0; i < 5; i++ {  // Retry 5 times
		db, err = sql.Open("mysql", dbCfg.FormatDSN())
		if err == nil {
			err = db.Ping()
		}
		if err == nil {
			break
		}

		log.Printf("Failed to connect to database. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, err  // return the error if after 5 attempts you couldn't connect
	}

	fmt.Println("Connected!")

	/*showTablesQuery, err := db.Query("SHOW TABLES")

	defer showTablesQuery.Close()
	for showTablesQuery.Next() {
		var tableName string

		// Get table name
		err = showTablesQuery.Scan(&tableName)

		checkError("Error querying tables", err)

		selectQuery, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))

		defer selectQuery.Close()

		checkError("Error creating the query", err)

		// Get column names from the given table
		columns, err := selectQuery.Columns()
		if err != nil {
			checkError(fmt.Sprintf("Error getting columns from table %s", tableName), err)
		}

		fmt.Printf("table name: %s -- columns: %v\n", tableName, strings.Join(columns, ", "))
	}*/
	return db, nil
}

// helper function to handle errors
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
