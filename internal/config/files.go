package config

import (
	"fivetran/internal/config/data"
	"github.com/adrg/xdg"
	"github.com/rs/zerolog/log"
	"os"
	"os/user"
	"path/filepath"
)

const (
	FivetranEnvConfigDir = "FIVETRAN_CONFIG_DIR"
	FivetranEnvLogsDir   = "FIVETRAN_LOGS_DIR"
	AppName              = "fivetran"
	FivetranLogsFile     = "fivetran.log"
)

var (
	AppLogFile string
)

func InitLogLoc() error {
	var appLogDir string
	switch {
	case isEnvSet(FivetranEnvLogsDir):
		appLogDir = os.Getenv(FivetranEnvLogsDir)
	case isEnvSet(FivetranEnvConfigDir):
		tmpDir, err := UserTmpDir()
		if err != nil {
			return err
		}
		appLogDir = tmpDir
	default:
		var err error
		appLogDir, err = xdg.StateFile(AppName)
		if err != nil {
			return err
		}
	}
	if err := data.EnsureFullPath(appLogDir, data.DefaultDirMod); err != nil {
		return err
	}
	AppLogFile = filepath.Join(appLogDir, FivetranLogsFile)

	return nil
}

func InitLocs() error {
	if isEnvSet(FivetranEnvConfigDir) {
		return initAppEnvLocs()
	}

	return initXDGLocs()
}

func initAppEnvLocs() error {
	AppConfigDir = os.Getenv(FivetranEnvConfigDir)
	if err := data.EnsureFullPath(AppConfigDir, data.DefaultDirMod); err != nil {
		return err
	}

	AppDumpsDir = filepath.Join(AppConfigDir, "screen-dumps")
	if err := data.EnsureFullPath(AppDumpsDir, data.DefaultDirMod); err != nil {
		log.Warn().Err(err).Msgf("Unable to create screen-dumps dir: %s", AppDumpsDir)
	}
	AppBenchmarksDir = filepath.Join(AppConfigDir, "benchmarks")
	if err := data.EnsureFullPath(AppBenchmarksDir, data.DefaultDirMod); err != nil {
		log.Warn().Err(err).Msgf("Unable to create benchmarks dir: %s", AppBenchmarksDir)
	}
	AppSkinsDir = filepath.Join(AppConfigDir, "skins")
	if err := data.EnsureFullPath(AppSkinsDir, data.DefaultDirMod); err != nil {
		log.Warn().Err(err).Msgf("Unable to create skins dir: %s", AppSkinsDir)
	}
	AppContextsDir = filepath.Join(AppConfigDir, "clusters")
	if err := data.EnsureFullPath(AppContextsDir, data.DefaultDirMod); err != nil {
		log.Warn().Err(err).Msgf("Unable to create clusters dir: %s", AppContextsDir)
	}

	AppConfigFile = filepath.Join(AppConfigDir, data.MainConfigFile)
	AppHotKeysFile = filepath.Join(AppConfigDir, "hotkeys.yaml")
	AppAliasesFile = filepath.Join(AppConfigDir, "aliases.yaml")
	AppPluginsFile = filepath.Join(AppConfigDir, "plugins.yaml")
	AppViewsFile = filepath.Join(AppConfigDir, "views.yaml")

	return nil
}

func initXDGLocs() error {
	var err error

	AppConfigDir, err = xdg.ConfigFile(AppName)
	if err != nil {
		return err
	}

	AppConfigFile, err = xdg.ConfigFile(filepath.Join(AppName, data.MainConfigFile))
	if err != nil {
		return err
	}

	AppHotKeysFile = filepath.Join(AppConfigDir, "hotkeys.yaml")
	AppAliasesFile = filepath.Join(AppConfigDir, "aliases.yaml")
	AppPluginsFile = filepath.Join(AppConfigDir, "plugins.yaml")
	AppViewsFile = filepath.Join(AppConfigDir, "views.yaml")

	AppSkinsDir = filepath.Join(AppConfigDir, "skins")
	if err := data.EnsureFullPath(AppSkinsDir, data.DefaultDirMod); err != nil {
		log.Warn().Err(err).Msgf("No skins dir detected")
	}

	AppDumpsDir, err = xdg.StateFile(filepath.Join(AppName, "screen-dumps"))
	if err != nil {
		return err
	}

	AppBenchmarksDir, err = xdg.StateFile(filepath.Join(AppName, "benchmarks"))
	if err != nil {
		log.Warn().Err(err).Msgf("No benchmarks dir detected")
	}

	dataDir, err := xdg.DataFile(AppName)
	if err != nil {
		return err
	}
	AppContextsDir = filepath.Join(dataDir, "clusters")
	if err := data.EnsureFullPath(AppContextsDir, data.DefaultDirMod); err != nil {
		log.Warn().Err(err).Msgf("No context dir detected")
	}

	return nil
}

func isEnvSet(env string) bool {
	return os.Getenv(env) != ""
}

func UserTmpDir() (string, error) {
	u, err := user.Current()
	if err != nil {

		return "", err
	}

	dir := filepath.Join(os.TempDir(), u.Username, AppName)

	return dir, nil
}
