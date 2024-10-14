package main

import (
	"fivetran/api"
	"os"
)

func main() {
	api.Refresh("", "Warehouse", "1sIoTnItnQRuOFrL9L1rRpfw8valt_1faaz4V9-1nIrg")
	//cmd.Execute()
	os.Exit(0)

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
