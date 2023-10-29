package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func serviceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get(apiKeyName) != key.Value {
		w.WriteHeader(400)
		return
	}

	serviceID := req.URL.Query().Get("service_id")
	startTS := req.URL.Query().Get("start_ts")
	endTS := req.URL.Query().Get("end_ts")

	query := `
		SELECT alert_id, service_id, service_name, model, alert_type, alert_ts, severity, team_slack
		FROM alerts
		WHERE service_id = ? AND alert_ts >= ? AND alert_ts <= ?
	`

	rows, err := db.conn.Query(query, serviceID, startTS, endTS)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	data := ServiceResponse{ServiceID: serviceID}
	var alerts []Alert
	for rows.Next() {
		var alert Alert
		if err := rows.Scan(&alert.AlertID, &alert.ServiceID, &alert.ServiceName, &alert.Model, &alert.AlertType, &alert.AlertTS, &alert.Severity, &alert.TeamSlack); err != nil {
			log.Fatal(err)
		}
		alerts = append(alerts, alert)
	}

	data.ServiceName = alerts[0].ServiceName
	data.Alerts = removeServiceNameFromAlerts(alerts)
	resp, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	if err == nil {
		w.WriteHeader(200)
		fmt.Fprintf(w, string(resp))
	} else {
		var alertResponse AlertResponse
		alertResponse.Error = err.Error()
		resp, _ = json.Marshal(alertResponse)
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(resp))
}

func listHandler(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get(apiKeyName) != key.Value {
		w.WriteHeader(400)
		return
	}

	rows, err := db.conn.Query("SELECT alert_id, service_id, service_name, model, alert_type, alert_ts, severity, team_slack FROM alerts")
	if err != nil {
		log.Fatal(err)
	}
	var alert Alert
	var alerts []Alert
	for rows.Next() {
		rows.Scan(&alert.AlertID, &alert.ServiceID, &alert.ServiceName, &alert.Model, &alert.AlertType, &alert.AlertTS, &alert.Severity, &alert.TeamSlack)
		alerts = append(alerts, alert)
	}

	resp, err := json.Marshal(alerts)
	if err != nil {
		log.Fatal(err)
	}

	if err == nil {
		w.WriteHeader(200)
	} else {
		var alertResponse AlertResponse
		alertResponse.Error = err.Error()
		resp, _ = json.Marshal(alertResponse)
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(resp))
}

func createHandler(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get(apiKeyName) != key.Value {
		w.WriteHeader(400)
		return
	}

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	var alert Alert
	err = json.Unmarshal(reqBody, &alert)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	alert.AlertID = getID()
	alert.AlertTS = getTimeStamp()

	statement, err := db.conn.Prepare("INSERT INTO alerts (alert_id, service_id, service_name, model, alert_type, alert_ts,severity, team_slack) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	res, err := statement.Exec(alert.AlertID, alert.ServiceID, alert.ServiceName, alert.Model, alert.AlertType, alert.AlertTS, alert.Severity, alert.TeamSlack)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	var alertResponse AlertResponse
	alertResponse.AlertID = alert.AlertID
	if rowsAffected == 1 {
		w.WriteHeader(200)
		alertResponse.Error = ""
	} else {
		w.WriteHeader(500)
		alertResponse.Error = err.Error()
	}

	resp, _ := json.Marshal(alertResponse)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(resp))
}
