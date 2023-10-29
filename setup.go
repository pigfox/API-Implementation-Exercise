package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func setUp() {
	db = initDB()
	if db.conn != nil {
		fmt.Println("DB initialized")
	}

	godotenv.Load(".env")
	key.Value = os.Getenv("API_KEY")

	if key.Value != "" {
		fmt.Println("API key set")
	} else {
		log.Fatal("Quitting: API key Not Set")
	}
}

func initDB() DB {
	var err error
	db.conn, err = sql.Open("sqlite3", "Alerts.db")
	if err != nil {
		log.Fatal("DB initialization failed: " + err.Error())
	}
	//defer db.conn.Close()
	createTable()

	return db
}

func createTable() {
	sql := `
CREATE TABLE IF NOT EXISTS alerts (
	alert_id TEXT PRIMARY KEY,
	service_id TEXT,
	service_name TEXT,
	model TEXT,
	alert_type TEXT,
	alert_ts INTEGER,
	severity TEXT,
	team_slack TEXT
);
`
	_, err := db.conn.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}
