package cmd

import (
	"errors"
	"fivetran/internal/client"
	"fivetran/internal/config"
	"github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
)

const (
	appName      = config.AppName
	shortAppDesc = "A graphical CLI for fivetran management."
	longAppDesc  = "Fivetran is a CLI to view and manage your fivetran account."
)

var (
	version, commit, date = "dev", "dev", client.NA
	k9sFlags              *config.Flags
	k8sFlags              *config.ConfigFlags

	rootCmd = &cobra.Command{
		Use:   appName,
		Short: shortAppDesc,
		Long:  longAppDesc,
		RunE:  run,
	}

	out = colorable.NewColorableStdout()
)

type flagError struct{ err error }

func (e flagError) Error() string { return e.err.Error() }

func init() {
	//if err := config.InitLogLoc(); err != nil {
	//	fmt.Printf("Fail to init k9s logs location %s\n", err)
	//}

	rootCmd.SetFlagErrorFunc(func(command *cobra.Command, err error) error {
		return flagError{err: err}
	})

	rootCmd.AddCommand(versionCmd(), infoCmd())
	//initK9sFlags()
	//initK8sFlags()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if !errors.As(err, &flagError{}) {
			panic(err)
		}
	}
	// fmt.Printf("appName: %s\ndesc: %s\ndesc:%s\n", appName, shortAppDesc, longAppDesc)
}

func run(cmd *cobra.Command, args []string) error {
	//if err := config.InitLocs(); err != nil {
	//	return err
	//}
	//file, err := os.OpenFile(
	//	*k9sFlags.LogFile,
	//	os.O_CREATE|os.O_APPEND|os.O_WRONLY,
	//	data.DefaultFileMod,
	//)
	//if err != nil {
	//	return fmt.Errorf("Log file %q init failed: %w", *k9sFlags.LogFile, err)
	//}
	//defer func() {
	//	if file != nil {
	//		_ = file.Close()
	//	}
	//}()
	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Error().Msgf("Boom! %v", err)
	//		log.Error().Msg(string(debug.Stack()))
	//		printLogo(color.Red)
	//		fmt.Printf("%s", color.Colorize("Boom!! ", color.Red))
	//		fmt.Printf("%v.\n", err)
	//	}
	//}()
	//
	//log.Logger = log.Output(zerolog.ConsoleWriter{Out: file})
	//zerolog.SetGlobalLevel(parseLevel(*k9sFlags.LogLevel))
	//
	//cfg, err := loadConfiguration()
	//if err != nil {
	//	log.Error().Err(err).Msgf("Fail to load global/context configuration")
	//}
	//app := view.NewApp(cfg)
	//if err := app.Init(version, *k9sFlags.RefreshRate); err != nil {
	//	return err
	//}
	//if err := app.Run(); err != nil {
	//	return err
	//}
	//if view.ExitStatus != "" {
	//	return fmt.Errorf("view exit status %s", view.ExitStatus)
	//}

	return nil
}
