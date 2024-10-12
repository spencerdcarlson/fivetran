package main

import (
	"fivetran/api"
	"fivetran/cmd"
	"os"
)

func main() {
	api.Refresh("", "Warehouse")
	cmd.Execute()
	os.Exit(0)

	//fmt.Printf("Found %d total connectors.\n", len(connectors.Data.Items))
	//
	//googleSheetConnectors := api.ConnectorsByService(connectors, "google_sheets")
	//
	//fmt.Printf("Found %d 'google_sheets' connectors.\n", len(googleSheetConnectors))
	//
	//// Get Connector details for all google_sheet connectors
	//var connectorDetails []api.ConnectorResponse
	//var wg sync.WaitGroup
	//concurrencyLimit := 10
	//sem := make(chan struct{}, concurrencyLimit)
	//var mu sync.Mutex
	//
	//for _, connector := range googleSheetConnectors {
	//	wg.Add(1)
	//
	//	// Acquire a slot in the semaphore
	//	sem <- struct{}{}
	//
	//	go func(connectorID string) {
	//		defer wg.Done()
	//		defer func() { <-sem }() // Release the slot in the semaphore
	//
	//		url := fmt.Sprintf("https://api.fivetran.com/v1/connectors/%s", connectorID)
	//		req, err := http.NewRequest("GET", url, nil)
	//		if err != nil {
	//			fmt.Println("Error creating request:", err)
	//			return
	//		}
	//		req.SetBasicAuth(values.APIKey, values.APISecret)
	//
	//		resp, err := client.Do(req)
	//		if err != nil {
	//			fmt.Println("Error making request:", err)
	//			return
	//		}
	//		defer resp.Body.Close()
	//
	//		body, err := io.ReadAll(resp.Body)
	//		if err != nil {
	//			fmt.Println("Error reading response body:", err)
	//			return
	//		}
	//
	//		var connectorDetail api.ConnectorResponse
	//		err = json.Unmarshal(body, &connectorDetail)
	//		if err != nil {
	//			fmt.Println("Error unmarshaling response:", err)
	//			return
	//		}
	//
	//		// Safely append the result to the slice
	//		mu.Lock()
	//		connectorDetails = append(connectorDetails, connectorDetail)
	//		mu.Unlock()
	//	}(connector.Id)
	//}
	//// Wait for all goroutines to finish
	//wg.Wait()
	//
	//connector, err := api.ConnectorBySheetURL(connectorDetails, values.URLPart)
	//if err != nil {
	//	fmt.Println("Error finding Connector:", err)
	//	return
	//}
	//
	//jsonBytes, err = json.MarshalIndent(connector, "\t\t", "\t")
	//if err != nil {
	//	fmt.Println("Error marshaling to JSON:", err)
	//	return
	//}
	//
	//fmt.Printf("Found Connector from Google Sheet URL\n\tURL Part: %s\n\tConnector: %s\n", values.URLPart, string(jsonBytes))
	//fmt.Printf("Attempting to trigger sync.\n")
	//
	//url = fmt.Sprintf("https://api.fivetran.com/v1/connectors/%s/sync", connector.Id)
	//requestData := api.ConnectorSyncRequest{
	//	Force: false,
	//}
	//
	//jsonData, err := json.Marshal(requestData)
	//if err != nil {
	//	fmt.Println("Error marshaling JSON:", err)
	//	return
	//}
	//
	//req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	//if err != nil {
	//	fmt.Println("Error creating request:", err)
	//	return
	//}
	//req.Header.Set("Content-Type", "application/json")
	//req.SetBasicAuth(values.APIKey, values.APISecret)
	//
	//resp, err = client.Do(req)
	//if err != nil {
	//	fmt.Println("Error making request:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//body, err = io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error reading response body:", err)
	//	return
	//}
	//
	//var syncResponse api.ConnectorSyncResponse
	//err = json.Unmarshal(body, &syncResponse)
	//if err != nil {
	//	fmt.Println("Error making request:", err)
	//	return
	//}
	//
	//if strings.EqualFold(strings.TrimSpace(syncResponse.Code), "Success") {
	//	fmt.Printf("Sync request was successful.\n")
	//} else {
	//	fmt.Printf("Sync request failed. %+v\n", syncResponse)
	//}
}
