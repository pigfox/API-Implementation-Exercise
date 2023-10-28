package main

import (
	"database/sql"
	"time"
)

type DB struct {
	conn *sql.DB
}

type Key struct {
	Value string
}

type alert struct {
	AlertID     string    `json:"alert_id"`
	ServiceID   string    `json:"service_id"`
	ServiceName string    `json:"service_name"`
	Model       string    `json:"model"`
	AlertType   string    `json:"alert_type"`
	AlertTS     time.Time `json:"alert_ts"`
	Severity    string    `json:"warning"`
	TeamSlack   string    `json:"team_slack"`
}
