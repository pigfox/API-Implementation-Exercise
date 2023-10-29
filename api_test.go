package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func getENVKey() string {
	godotenv.Load(".env")
	return os.Getenv("API_KEY")
}

func TestPost(t *testing.T) {
	req, _ := http.NewRequest("POST", "http://localhost:8888/alerts", bytes.NewBuffer([]byte(getPostPayload())))
	req.Header.Add(apiKeyName, getENVKey())
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	var alertResponse AlertResponse

	err = json.Unmarshal([]byte(resBody), &alertResponse)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	expectedLenth := 32
	length := len(alertResponse.AlertID)
	if length != expectedLenth {
		t.Errorf("Expected an AlertID of "+strconv.Itoa(expectedLenth)+". Got %d", length)
	}
}

func TestGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8888//alerts?service_id=my_test_service_id&start_ts=1698521301269568032&end_ts=1698521306322457739", nil)
	req.Header.Add(apiKeyName, getENVKey())
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	if body := string(resBody); body == "" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func getPostPayload() string {
	return `{
		"service_id": "my_test_service_id",
		"service_name": "my_test_service",
		"model": "my_test_model",
		"alert_type": "anomaly",
		"severity": "warning",
		"team_slack": "slack_ch"
		}`
}
