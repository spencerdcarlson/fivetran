package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func GroupItemByName(response GroupsResponse, name string) (*GroupItem, error) {
	for _, item := range response.Data.Items {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("fivetran: '%s' GroupItem not found", name))
}

func ConnectorsByService(response GroupConnectorsResponse, service string) []ConnectorItem {
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

func IsGenericError(response GenericResponse) bool {
	info, ok := validCodes[response.Code]
	if !ok {
		return false
	}
	return info.IsError
}

func ListResource(key string, resource string) (*Response, error) {
	auth, err := BuildAuth(key)
	if err != nil {
		return nil, err
	}
	body, err := List(auth, resource)
	if err != nil {
		return nil, err
	}
	if strings.ToLower(strings.TrimSpace(resource)) == "groups" {
		var response GroupsResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &Response{Type: GroupsResponseType, GroupsResponse: response}, nil
	}
	return nil, errors.New("fivetran: resource not supported")
}
