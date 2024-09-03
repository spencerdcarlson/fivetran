package api

import (
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
