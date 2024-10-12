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

func IsError(code CodeType) bool {
	info, ok := validCodes[code]
	if !ok {
		return false
	}
	return info.IsError
}

func IsArrayEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func NormalizeResource(resource string) (*string, *string) {
	// Look up plural from singular
	r := strings.TrimSpace(strings.ToLower(resource))
	if _, ok := Resources[r]; ok {
		plural := Resources[r]
		return &r, &plural
	}
	// Look up plural from plurals
	for singular, plural := range Resources {
		if plural == r {
			return &singular, &plural
		}
	}
	// No match
	return nil, nil
}

func ListResource(key string, resources []string) (*Response, error) {
	auth, err := BuildAuth(key)
	if err != nil {
		return nil, err
	}

	// Replace with plural if exists, otherwise keep input
	for i, item := range resources {
		_, plural := NormalizeResource(item)
		if plural != nil {
			resources[i] = *plural
		}
	}

	println(fmt.Sprintf("Resources %v", resources))

	body, err := List(auth, resources)
	if err != nil {
		return nil, err
	}

	// Get last resource in list of resources
	resource := resources[len(resources)-1]
	_, plural := NormalizeResource(resource)
	if plural != nil {
		r := *plural
		if r == "groups" {
			var response GroupsResponse
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			if IsError(response.Code) {
				return nil, errors.New(fmt.Sprintf("fivetran: '%s' Groups not found", response.Code))
			}
			return &Response{Type: GroupsResponseType, GroupsResponse: response.Data}, nil
		} else if r == "connectors" {
			var response ConnectorsResponse
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			if IsError(response.Code) {
				return nil, errors.New(fmt.Sprintf("fivetran: '%s' Connector not found", response.Code))
			}
			return &Response{Type: ConnectorsResponseType, ConnectorsResponse: response.Data}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("The '%s' resource is not supported. response: %s", resource, string(body)))
}

func Refresh(apiKey string, groupName string) {
	// Get all Groups
	response, err := ListResource(apiKey, []string{"groups"})
	if err != nil {
		//return err
	}
	if response.Type == GroupsResponseType {
		// Find group that matches groupName
		item, err := GroupItemByName(response.GroupsResponse, groupName)
		if err != nil {

		}

		jsonBytes, err := json.MarshalIndent(item, "\t", "\t")
		if err != nil {
			fmt.Println("Error marshaling to JSON:", err)
			return
		}
		fmt.Printf("Found Fivetran Group (Sink) from name:\n\tName: %s\n\tGroup: %s\n", groupName, string(jsonBytes))

		// Get Group connectors for that Group
		response, err := ListResource(apiKey, []string{"groups", item.Id, "connectors"})
		if err != nil {
			//return err
		}
		if response.Type == ConnectorsResponseType {
			if err != nil {
				//return err
			}
			//println(fmt.Sprintf("Group connectors %v", response))
			fmt.Printf("Found %d total connectors.\n", len(response.ConnectorsResponse.Items))
		}

	}
}
