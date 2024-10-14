package api

import (
	"encoding/json"
	"errors"
	"net/url"
)

type Coded interface {
	IsError() bool
}

type FiveTranResponse interface {
	GetResponse() interface{}
}

type GroupItem struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type GroupsData struct {
	Items []GroupItem `json:"items"` // Changed to export the field and added JSON tag
}

type GroupsResponse struct {
	Code CodeType   `json:"code"`
	Data GroupsData `json:"data"`
}

type Filterable interface {
	GetBy(string, string) interface{}
}

type ConnectorItem struct {
	Id      string `json:"id"`
	Service string `json:"service"`
	Schema  string `json:"schema"`
}

type ConnectorsData struct {
	Items []ConnectorItem `json:"items"` // Changed to export the field and added JSON tag
}

type ConnectorsResponse struct {
	Code CodeType       `json:"code"`
	Data ConnectorsData `json:"data"`
}

type ConnectorConfig struct {
	AuthType   string `json:"auth_type"`
	SheetId    string `json:"sheet_id"`
	URL        *url.URL
	NamedRange string `json:"named_range"`
}

type ConnectorSyncDetails struct {
	LastSynced string `json:"last_synced"`
}

type Connector struct {
	Id                string               `json:"id"`
	Service           string               `json:"service"`
	Schema            string               `json:"schema"`
	SourceSyncDetails ConnectorSyncDetails `json:"source_sync_details"`
	Config            ConnectorConfig      `json:"config"`
}

type ConnectorResponse struct {
	Code CodeType  `json:"code"`
	Data Connector `json:"data"`
}

type ConnectorSyncRequest struct {
	Force bool `json:"force"`
}

type ConnectorSyncResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GenericResponse struct {
	Code    CodeType `json:"code"`
	Message string   `json:"message"`
}

type CodeType string

const (
	NotFound   CodeType = "NotFound"
	AuthFailed CodeType = "AuthFailed"
	Success    CodeType = "Success"
)

type CodeTypeInfo struct {
	Error bool
}

var validCodes = map[CodeType]CodeTypeInfo{
	NotFound:   {Error: true},
	AuthFailed: {Error: true},
	Success:    {Error: false},
}

type Resource struct {
	Singular string `json:"singular"`
	Plural   string `json:"plural"`
}

var Resources = [...]Resource{
	{Singular: "group", Plural: "groups"},
	{Singular: "connector", Plural: "connectors"},
}

func (c *CodeType) UnmarshalJSON(b []byte) error {
	var codeStr string
	if err := json.Unmarshal(b, &codeStr); err != nil {
		return err
	}

	codeType := CodeType(codeStr)
	if _, ok := validCodes[codeType]; !ok {
		return errors.New("invalid code type: " + codeStr)
	}

	*c = codeType
	return nil
}

func (cc *ConnectorConfig) UnmarshalJSON(data []byte) error {
	type Alias ConnectorConfig

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(cc),
	}

	// Unmarshal into the auxiliary struct.
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Attempt to parse the SheetId as a URL.
	parsedURL, err := url.Parse(cc.SheetId)
	if err != nil {
		// If parsing fails, set URL to nil.
		cc.URL = nil
	} else {
		// If parsing succeeds, set the parsed URL.
		cc.URL = parsedURL
	}

	return nil
}
