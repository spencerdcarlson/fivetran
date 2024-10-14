package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

const baseURL = "https://api.fivetran.com/v1"

func List(auth *Auth, resources []string) ([]byte, error) {
	path := strings.Join(resources, "/")
	url := fmt.Sprintf("%s/%s", baseURL, path)
	fmt.Printf("URL: %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	//println(fmt.Sprintf("Key: %s, Secret: %s", auth.APIKey, auth.APISecret))
	req.SetBasicAuth(auth.APIKey, auth.APISecret)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}
	var response GenericResponse
	_ = json.Unmarshal(body, &response)
	if response.IsError() {
		return nil, errors.New(fmt.Sprintf("API returned a generic error. %+v", response))
	}
	return body, nil
}

func SyncConnectorData(auth *Auth, connectorId string) ([]byte, error) {
	url := fmt.Sprintf("%s/connectors/%s/sync", baseURL, connectorId)
	requestData := ConnectorSyncRequest{
		Force: false,
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error preparing request body: %w", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(auth.APIKey, auth.APISecret)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	var response GenericResponse
	_ = json.Unmarshal(body, &response)
	if response.IsError() {
		return nil, errors.New(fmt.Sprintf("API returned a generic error. %+v", response))
	}
	return body, nil
}
