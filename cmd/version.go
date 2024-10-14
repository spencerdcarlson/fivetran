package cmd

import (
	"encoding/json"
	"errors"
	"fivetran/api"
	"fivetran/internal/color"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func parseFilters(filters []string) map[string]string {
	filterMap := make(map[string]string)
	for _, filter := range filters {
		keyValue := strings.SplitN(filter, "=", 2)
		if len(keyValue) == 2 {
			filterMap[keyValue[0]] = keyValue[1]
		} else {
			fmt.Printf("Ignoring invalid filter: %s\n", filter)
		}
	}
	return filterMap
}

func listCmd() *cobra.Command {
	var key string
	var filters []string
	command := cobra.Command{
		Use:   "list [id] [resource]",
		Short: "List resource",
		Long:  "List API resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			filterMap := parseFilters(filters)
			println(fmt.Sprintf("filters: %v, map: %v", filters, filterMap))
			if resource, ok := cmd.Parent().Annotations["resource"]; ok {
				resources := append([]string{resource}, args...)
				println(fmt.Sprintf("resources: %v", resources))
				response, err := api.ListResource(key, resources, filterMap)
				if err != nil {
					return err
				}
				jsonBytes, err := json.MarshalIndent(response, "", "\t")
				if err != nil {
					return err
				}
				fmt.Printf("%s\n", string(jsonBytes))
				return nil
			} else {
				return errors.New("no resource")
			}
		},
	}
	command.PersistentFlags().StringVar(&key, "key", "", "API key")
	command.Flags().StringSliceVar(&filters, "filter", []string{}, "Filter in key=value format")

	return &command
}

func groupCmd() *cobra.Command {
	var short bool

	command := cobra.Command{
		Use: "group [command]",
		//Annotations: map[string]string{"resource": "causeError"},
		//Annotations: map[string]string{"resource": "connectors"},
		Annotations: map[string]string{"resource": "groups"},
		Short:       "Group resource",
		Long:        "Group resource root command",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion(short)
		},
	}

	command.AddCommand(listCmd())

	// Sub command or flags
	//command.PersistentFlags().BoolVarP(&short, "short", "s", false, "Prints version info in short format")

	return &command
}

func versionCmd() *cobra.Command {
	var short bool

	command := cobra.Command{
		Use:   "version",
		Short: "Print version/build info",
		Long:  "Print version/build information",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion(short)
		},
	}

	command.PersistentFlags().BoolVarP(&short, "short", "s", false, "Prints version info in short format")

	return &command
}

func printVersion(short bool) {
	const fmat = "%-20s %s\n"
	var outputColor color.Paint

	if short {
		outputColor = -1
	} else {
		outputColor = color.Cyan
		printLogo(outputColor)
	}
	printTuple(fmat, "Version", version, outputColor)
	printTuple(fmat, "Commit", commit, outputColor)
	printTuple(fmat, "Date", date, outputColor)
}

func printTuple(fmat, section, value string, outputColor color.Paint) {
	if outputColor != -1 {
		fmt.Fprintf(out, fmat, color.Colorize(section+":", outputColor), value)
		return
	}
	fmt.Fprintf(out, fmat, section, value)
}
