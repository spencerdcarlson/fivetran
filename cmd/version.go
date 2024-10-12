package cmd

import (
	"encoding/json"
	"errors"
	"fivetran/api"
	"fivetran/internal/color"
	"fmt"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	var key string
	command := cobra.Command{
		Use:   "list",
		Short: "List resource",
		Long:  "List API resources",
		RunE: func(cmd *cobra.Command, args []string) error {

			if resource, ok := cmd.Parent().Annotations["resource"]; ok {
				resources := append([]string{resource}, args...)
				println(fmt.Sprintf("resources: %v", resources))
				response, err := api.ListResource(key, resources)
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
