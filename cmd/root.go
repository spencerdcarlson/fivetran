package cmd

import (
	"errors"
	"fivetran/internal/client"
	"fivetran/internal/color"
	"fivetran/internal/config"
	"fivetran/internal/config/data"
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"runtime/debug"
)

const (
	appName      = config.AppName
	shortAppDesc = "A graphical CLI for fivetran management."
	longAppDesc  = "FivetranConfig is a CLI to view and manage your fivetran account."
)

var (
	version, commit, date = "dev", "dev", client.NA
	appFlags              *config.Flags
	clientFlags           *client.Flags

	rootCmd = &cobra.Command{
		Use:   appName,
		Short: shortAppDesc,
		Long:  longAppDesc,
		RunE:  run,
	}

	out = colorable.NewColorableStdout()
)

type flagError struct {
	err error
}

func (e flagError) Error() string { return e.err.Error() }

func init() {
	if err := config.InitLogLoc(); err != nil {
		fmt.Printf("Fail to init logs location %s\n", err)
	}

	rootCmd.SetFlagErrorFunc(func(command *cobra.Command, err error) error {
		return flagError{err: err}
	})

	// Add commands
	rootCmd.AddCommand(versionCmd(), infoCmd(), groupCmd())

	// Add flags
	initAppFlags()
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
	if err := config.InitFileLocations(); err != nil {
		return err
	}
	file, err := os.OpenFile(
		*appFlags.LogFile,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		data.DefaultFileMod,
	)
	if err != nil {
		return fmt.Errorf("log file %q init failed: %w", *appFlags.LogFile, err)
	}
	defer func() {
		if file != nil {
			_ = file.Close()
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("Boom! %v", err)
			log.Error().Msg(string(debug.Stack()))
			printLogo(color.Red)
			fmt.Printf("%s", color.Colorize("Boom!! ", color.Red))
			fmt.Printf("%v.\n", err)
		}
	}()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: file})
	zerolog.SetGlobalLevel(parseLevel(*appFlags.LogLevel))

	//log.Info().Msg("starting up...")

	loadConfiguration()
	//cfg, err := loadConfiguration()
	//if err != nil {
	//	log.Error().Err(err).Msgf("Fail to load global/context configuration")
	//}
	//app := view.NewApp(cfg)
	//if err := app.Init(version, *appFlags.RefreshRate); err != nil {
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

// func loadConfiguration() (*config.Config, error) {
func loadConfiguration() (*config.Config, error) {
	log.Info().Msg("starting up...")
	clientCfg := client.NewConfig(clientFlags)
	fmt.Printf("k8sCfg %+v\n", clientCfg)
	cfg := config.NewConfig(clientCfg)
	var errs error
	// This is where the yaml or json config file gets loaded!
	if err := cfg.Load(config.AppConfigFile, false); err != nil {
		errs = errors.Join(errs, err)
	}
	//k9sCfg.FivetranConfig.Override(k9sFlags)
	//if err := k9sCfg.Refine(clientFlags, k9sFlags, k8sCfg); err != nil {
	//	log.Error().Err(err).Msgf("config refine failed")
	//	errs = errors.Join(errs, err)
	//}
	//// Try to access server version if that fail. Connectivity issue?
	//if !conn.CheckConnectivity() {
	//	errs = errors.Join(errs, fmt.Errorf("cannot connect to context: %s", k9sCfg.FivetranConfig.ActiveContextName()))
	//}
	//if !conn.ConnectionOK() {
	//	errs = errors.Join(errs, fmt.Errorf("k8s connection failed for context: %s", k9sCfg.FivetranConfig.ActiveContextName()))
	//}
	//
	//log.Info().Msg("✅ Kubernetes connectivity")
	//if err := k9sCfg.Save(false); err != nil {
	//	log.Error().Err(err).Msg("Config save")
	//	errs = errors.Join(errs, err)
	//}
	//
	//return k9sCfg, errs
	return nil, nil
}

func parseLevel(level string) zerolog.Level {
	switch level {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}

func initAppFlags() {
	appFlags = config.NewFlags()
	rootCmd.Flags().StringVarP(
		appFlags.LogLevel,
		"logLevel", "l",
		config.DefaultLogLevel,
		"Specify a log level (info, warn, debug, trace, error)",
	)
	rootCmd.Flags().StringVarP(
		appFlags.LogFile,
		"logFile", "",
		config.AppLogFile,
		"Specify the log file",
	)
	rootCmd.Flags()
}
