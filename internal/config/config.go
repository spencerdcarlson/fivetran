package config

import (
	"errors"
	"fivetran/internal/config/data"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
)

const (
	DefaultLoggerTailCount = 100
	MaxLogThreshold        = 5000
	DefaultSinceSeconds    = -1 // tail logs by default
)

type FivetranConfig struct {
	DisablePodCounting bool   `json:"disablePodCounting" yaml:"disablePodCounting"`
	Logger             Logger `json:"logger" yaml:"logger"`
	clientSettings     data.ClientSettings
}

type Config struct {
	Fivetran       *FivetranConfig `yaml:"k9s" json:"k9s"`
	clientSettings data.ClientSettings
}

func NewConfig(c data.ClientSettings) *Config {
	return &Config{
		clientSettings: c,
		Fivetran:       NewFivetran(c),
	}
}

func NewFivetran(c data.ClientSettings) *FivetranConfig {
	return &FivetranConfig{
		Logger:         NewLogger(),
		clientSettings: c,
	}
}

func NewLogger() Logger {
	return Logger{
		TailCount:    DefaultLoggerTailCount,
		BufferSize:   MaxLogThreshold,
		SinceSeconds: DefaultSinceSeconds,
	}
}

type Logger struct {
	TailCount    int64 `json:"tail" yaml:"tail"`
	BufferSize   int   `json:"buffer" yaml:"buffer"`
	SinceSeconds int64 `json:"sinceSeconds" yaml:"sinceSeconds"`
	TextWrap     bool  `json:"textWrap" yaml:"textWrap"`
	ShowTime     bool  `json:"showTime" yaml:"showTime"`
}

func (c *Config) Load(path string, force bool) error {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		if err := c.Save(force); err != nil {
			return err
		}
	}
	bb, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var errs error
	//if err := data.JSONValidator.Validate(json.K9sSchema, bb); err != nil {
	//	errs = errors.Join(errs, fmt.Errorf("k9s config file %q load failed:\n%w", path, err))
	//}

	var cfg Config
	if err := yaml.Unmarshal(bb, &cfg); err != nil {
		errs = errors.Join(errs, fmt.Errorf("main config.yaml load failed: %w", err))
	}
	//c.Merge(&cfg)

	return errs
}

func (c *Config) Save(force bool) error {
	//c.Validate()
	if err := c.Fivetran.Save(force); err != nil {
		return err
	}
	if _, err := os.Stat(AppConfigFile); errors.Is(err, fs.ErrNotExist) {
		//return c.SaveFile(AppConfigFile)
		return nil
	}

	return nil
}

func (f *FivetranConfig) Save(force bool) error {
	//if f.getActiveConfig() == nil {
	//	log.Warn().Msgf("Save failed. no active config detected")
	//	return nil
	//}
	//path := filepath.Join(
	//	AppContextsDir,
	//	data.SanitizeContextSubpath(f.activeConfig.Context.GetClusterName(), f.getActiveContextName()),
	//	data.MainConfigFile,
	//)
	//path := "./"
	//if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) || force {
	//	return f.dir.Save(path, f.getActiveConfig())
	//}

	return nil
}

//func (f *FivetranConfig) getActiveConfig() *data.Config {
//	//f.mx.RLock()
//	//defer f.mx.RUnlock()
//
//	return f.activeConfig
//}
