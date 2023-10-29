package main

import (
	"database/sql"
)

type DB struct {
	conn *sql.DB
}

type Key struct {
	Value string
}

type AlertResponse struct {
	AlertID string `json:"alert_id"`
	Error   string `json:"error"`
}

type ServiceResponse struct {
	ServiceID   string         `json:"service_id"`
	ServiceName string         `json:"service_name"`
	Alerts      []ServiceAlert `json:"alerts"`
}

type Alert struct {
	AlertID     string `json:"alert_id"`
	ServiceID   string `json:"service_id"`
	ServiceName string `json:"service_name,omitempty"`
	Model       string `json:"model"`
	AlertType   string `json:"alert_type"`
	AlertTS     int64  `json:"alert_ts"`
	Severity    string `json:"severity"`
	TeamSlack   string `json:"team_slack"`
}

type ServiceAlert struct {
	AlertID   string `json:"alert_id"`
	ServiceID string `json:"service_id"`
	Model     string `json:"model"`
	AlertType string `json:"alert_type"`
	AlertTS   int64  `json:"alert_ts"`
	Severity  string `json:"severity"`
	TeamSlack string `json:"team_slack"`
}
