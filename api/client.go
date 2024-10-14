package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func List(auth *Auth, resources []string) ([]byte, error) {
	path := strings.Join(resources, "/")
	url := fmt.Sprintf("https://api.fivetran.com/v1/%s", path)
	fmt.Printf("URL: %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating request"))
	}
	//println(fmt.Sprintf("Key: %s, Secret: %s", auth.APIKey, auth.APISecret))
	req.SetBasicAuth(auth.APIKey, auth.APISecret)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error making request"))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, errors.New(fmt.Sprintf("Error reading response body"))
	}
	var response GenericResponse
	_ = json.Unmarshal(body, &response)
	if response.IsError() {
		return nil, errors.New(fmt.Sprintf("API returned a generic error. %+v", response))
	}
	return body, nil
}
