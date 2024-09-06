package cmd

import (
	"fivetran/internal/color"
	"fivetran/internal/config"
	"fivetran/internal/ui"
	"fmt"
	"github.com/spf13/cobra"
)

func printLogo(c color.Paint) {
	for _, l := range ui.LogoSmall {
		fmt.Fprintln(out, color.Colorize(l, c))
	}
	fmt.Fprintln(out)
}

func infoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "List configurations info",
		RunE:  printInfo,
	}
}

func printInfo(cmd *cobra.Command, args []string) error {
	if err := config.InitLocs(); err != nil {
		return err
	}

	const fmat = "%-27s %s\n"
	printLogo(color.Cyan)
	printTuple(fmat, "Version", version, color.Cyan)
	//printTuple(fmat, "Config", config.AppConfigFile, color.Cyan)
	//printTuple(fmat, "Custom Views", config.AppViewsFile, color.Cyan)
	//printTuple(fmat, "Plugins", config.AppPluginsFile, color.Cyan)
	//printTuple(fmat, "Hotkeys", config.AppHotKeysFile, color.Cyan)
	//printTuple(fmat, "Aliases", config.AppAliasesFile, color.Cyan)
	//printTuple(fmat, "Skins", config.AppSkinsDir, color.Cyan)
	//printTuple(fmat, "Context Configs", config.AppContextsDir, color.Cyan)
	//printTuple(fmat, "Logs", config.AppLogFile, color.Cyan)
	//printTuple(fmat, "Benchmarks", config.AppBenchmarksDir, color.Cyan)
	//printTuple(fmat, "ScreenDumps", getScreenDumpDirForInfo(), color.Cyan)

	return nil
}
