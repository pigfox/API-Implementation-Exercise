package main

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func getID() string {
	uuid := uuid.New().String()
	return strings.ReplaceAll(uuid, "-", "")
}

func getTimeStamp() int64 {
	return time.Now().UTC().UnixNano()
}

/*
This function is required in order to create the desired data scrutcture.
Where the field service_name is not in an alert, but in the document root.
*/
func removeServiceNameFromAlerts(alerts []Alert) []ServiceAlert {
	var sa []ServiceAlert
	for k, _ := range alerts {
		var serviceAlert ServiceAlert
		serviceAlert.AlertID = alerts[k].AlertID
		serviceAlert.Model = alerts[k].Model
		serviceAlert.AlertType = alerts[k].AlertType
		serviceAlert.AlertTS = alerts[k].AlertTS
		serviceAlert.Severity = alerts[k].Severity
		serviceAlert.TeamSlack = alerts[k].TeamSlack
		sa = append(sa, serviceAlert)
	}

	return sa
}
