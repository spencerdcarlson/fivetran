package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func GroupItemByName(data GroupsData, name string) (*GroupItem, error) {
	for _, item := range data.Items {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("fivetran: '%s' GroupItem not found", name))
}

func ConnectorsByService(response ConnectorsResponse, service string) []ConnectorItem {
	var items []ConnectorItem
	for _, item := range response.Data.Items {
		if item.Service == service {
			items = append(items, item)
		}
	}
	return items
}

func ConnectorBySheetURL(responses []ConnectorResponse, substr string) (*Connector, error) {
	for _, connector := range responses {
		if strings.Contains(connector.Data.Config.SheetId, substr) {
			return &connector.Data, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("fivetran: '%s' Connector not found", substr))
}

func ToResource(resource string) (*Resource, error) {
	r := strings.TrimSpace(strings.ToLower(resource))
	for i := range Resources {
		if r == Resources[i].Singular || r == Resources[i].Plural {
			return &Resources[i], nil
		}
	}
	return nil, errors.New("no resource found")
}

func (r GroupsResponse) GetBy(attr string, value string) interface{} {
	var result GroupItem
	for item := range r.Data.Items {
		if strings.ToLower(strings.TrimSpace(attr)) == "name" {
			if r.Data.Items[item].Name == value {
				result = r.Data.Items[item]
			}
		}
	}
	return result
}

func (c *CodeType) IsError() bool {
	if code, ok := validCodes[*c]; ok {
		return code.Error
	}
	return false
}

func (r GroupsResponse) IsError() bool {
	return r.Code.IsError()
}

func (r ConnectorsResponse) IsError() bool {
	return r.Code.IsError()
}

func (r ConnectorResponse) IsError() bool {
	return r.Code.IsError()
}

func (r GenericResponse) IsError() bool {
	return r.Code.IsError()
}

func (r GroupsResponse) GetResponse() interface{} {
	return r
}

func (r ConnectorResponse) GetResponse() interface{} {
	return r
}

func (r ConnectorsResponse) GetResponse() interface{} {
	return r
}

func ListResource(key string, resources []string) (FiveTranResponse, error) {
	auth, err := BuildAuth(key)
	if err != nil {
		return nil, err
	}

	// Replace with plural if exists, otherwise keep input
	for i, item := range resources {
		if r, err := ToResource(item); err == nil {
			resources[i] = r.Plural
		}
	}

	body, err := List(auth, resources)
	if err != nil {
		return nil, err
	}

	// If the last item on the resource path is not a valid resource then it is an ID
	// and this is a single resource request
	isSingle := false
	if _, err := ToResource(resources[len(resources)-1]); err != nil {
		isSingle = true
	}

	// Get last valid resource in list of resources
	var resource *Resource
	for i := len(resources) - 1; i >= 0; i-- {
		if r, err := ToResource(resources[i]); err == nil {
			resource = r
			fmt.Printf("Found resource %v from resource path %v\n", resource, resources)
			break
		}
	}

	if resource != nil {
		switch resource.Plural {
		case "groups":
			var response GroupsResponse
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			if response.IsError() {
				return nil, errors.New(fmt.Sprintf("fivetran: '%s' Groups not found", response.Code))
			}
			return response, nil
		case "connectors":
			if isSingle {
				var response ConnectorResponse
				err = json.Unmarshal(body, &response)
				if err != nil {
					return nil, err
				}
				if response.IsError() {
					return nil, errors.New(fmt.Sprintf("fivetran: '%s' Connector not found", response.Code))
				}
				return response, nil
			}
			var response ConnectorsResponse
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			if response.IsError() {
				return nil, errors.New(fmt.Sprintf("fivetran: '%s' Connector not found", response.Code))
			}
			return response, nil
		default:
			println(fmt.Sprintf("fivetran: '%s' Resource not found", resource))
		}
	}
	return nil, errors.New(fmt.Sprintf("The %v resource path is not supported. response: %s", resources, string(body)))
}

func Refresh(apiKey string, groupName string, urlPart string) {
	// Get all Groups
	if response, err := ListResource(apiKey, []string{"groups"}); err == nil {
		if groups, ok := response.(GroupsResponse); ok {
			if item, err := GroupItemByName(groups.Data, groupName); err == nil {
				if jsonBytes, err := json.MarshalIndent(item, "\t", "\t"); err == nil {
					fmt.Printf("Found Fivetran Group (Sink) from name:\n\tName: %s\n\tGroup: %s\n", groupName, string(jsonBytes))
				}
				if response, err := ListResource(apiKey, []string{"groups", item.Id, "connectors"}); err == nil {
					if connectors, ok := response.(ConnectorsResponse); ok {
						fmt.Printf("Found %d total connectors.\n", len(connectors.Data.Items))
						googleSheetConnectors := ConnectorsByService(connectors, "google_sheets")
						fmt.Printf("Found %d 'google_sheets' connectors.\n", len(googleSheetConnectors))

						//for _, connector := range googleSheetConnectors {
						var connectorResults []ConnectorResponse
						connector := googleSheetConnectors[0]
						fmt.Println("Checking google sheet connector")
						if response, err := ListResource(apiKey, []string{"connectors", connector.Id}); err == nil {
							if connector, ok := response.(ConnectorResponse); ok {
								connectorResults = append(connectorResults, connector)
							}
						}

						if foundConnector, err := ConnectorBySheetURL(connectorResults, urlPart); err == nil {
							fmt.Printf("Found connector %v\n", foundConnector)
							if jsonBytes, err := json.MarshalIndent(foundConnector, "\t", "\t"); err == nil {
								fmt.Printf("Found Connector from Google Sheet URL\n\tURL Part: %s\n\tConnector: %s\n", urlPart, string(jsonBytes))
							}
							fmt.Printf("Attempting to trigger sync.\n")
						}

						//}
					}
				}
			}
		}
	}
}
