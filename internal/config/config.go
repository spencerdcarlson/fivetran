package config

import (
	"fivetran/internal/client"
	"fivetran/internal/config/data"
	"sync"
)

const (
	defaultRefreshRate     = 2
	defaultMaxConnRetry    = 5
	DefaultLoggerTailCount = 100
	MaxLogThreshold        = 5000
	DefaultSinceSeconds    = -1 // tail logs by default
)

type K9s struct {
	LiveViewAutoRefresh bool   `json:"liveViewAutoRefresh" yaml:"liveViewAutoRefresh"`
	ScreenDumpDir       string `json:"screenDumpDir" yaml:"screenDumpDir,omitempty"`
	RefreshRate         int    `json:"refreshRate" yaml:"refreshRate"`
	MaxConnRetry        int    `json:"maxConnRetry" yaml:"maxConnRetry"`
	ReadOnly            bool   `json:"readOnly" yaml:"readOnly"`
	NoExitOnCtrlC       bool   `json:"noExitOnCtrlC" yaml:"noExitOnCtrlC"`
	SkipLatestRevCheck  bool   `json:"skipLatestRevCheck" yaml:"skipLatestRevCheck"`
	DisablePodCounting  bool   `json:"disablePodCounting" yaml:"disablePodCounting"`
	//ShellPod            ShellPod   `json:"shellPod" yaml:"shellPod"`
	//ImageScans          ImageScans `json:"imageScans" yaml:"imageScans"`
	Logger Logger `json:"logger" yaml:"logger"`
	//Thresholds          Threshold  `json:"thresholds" yaml:"thresholds"`
	manualRefreshRate   int
	manualHeadless      *bool
	manualLogoless      *bool
	manualCrumbsless    *bool
	manualReadOnly      *bool
	manualCommand       *string
	manualScreenDumpDir *string
	//dir                 *data.Dir
	activeContextName string
	//activeConfig        *data.Config
	//conn                client.Connection
	ks data.KubeSettings
	mx sync.RWMutex
}

type Config struct {
	K9s      *K9s `yaml:"k9s" json:"k9s"`
	conn     client.Connection
	settings data.KubeSettings
}

// NewConfig creates a new default config.
func NewConfig(ks data.KubeSettings) *Config {
	return &Config{
		settings: ks,
		K9s:      NewK9s(nil, ks),
	}
}

func NewK9s(conn client.Connection, ks data.KubeSettings) *K9s {
	return &K9s{
		RefreshRate:   defaultRefreshRate,
		MaxConnRetry:  defaultMaxConnRetry,
		ScreenDumpDir: AppDumpsDir,
		Logger:        NewLogger(),
		//Thresholds:    NewThreshold(),
		//ShellPod:      NewShellPod(),
		//ImageScans:    NewImageScans(),
		//dir:           data.NewDir(AppContextsDir),
		//conn:          conn,
		ks: ks,
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
