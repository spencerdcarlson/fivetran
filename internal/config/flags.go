package config

const (
	// DefaultRefreshRate represents the refresh interval.
	DefaultRefreshRate = 2 // secs

	// DefaultLogLevel represents the default log level.
	DefaultLogLevel = "info"

	// DefaultCommand represents the default command to run.
	DefaultCommand = ""
)

var (
	// AppConfigDir tracks main k9s config home directory.
	AppConfigDir string

	// AppSkinsDir tracks skins data directory.
	AppSkinsDir string

	// AppBenchmarksDir tracks benchmarks results directory.
	AppBenchmarksDir string

	// AppDumpsDir tracks screen dumps data directory.
	AppDumpsDir string

	// AppContextsDir tracks contexts data directory.
	AppContextsDir string

	// AppConfigFile tracks k9s config file.
	AppConfigFile string

	// AppViewsFile tracks custom views config file.
	AppViewsFile string

	// AppAliasesFile tracks aliases config file.
	AppAliasesFile string

	// AppPluginsFile tracks plugins config file.
	AppPluginsFile string

	// AppHotKeysFile tracks hotkeys config file.
	AppHotKeysFile string
)

type Flags struct {
	RefreshRate   *int
	LogLevel      *string
	LogFile       *string
	Headless      *bool
	Logoless      *bool
	Command       *string
	AllNamespaces *bool
	ReadOnly      *bool
	Write         *bool
	Crumbsless    *bool
	ScreenDumpDir *string
}

func NewFlags() *Flags {
	return &Flags{
		RefreshRate:   intPtr(DefaultRefreshRate),
		LogLevel:      strPtr(DefaultLogLevel),
		LogFile:       strPtr(AppLogFile),
		Headless:      boolPtr(false),
		Logoless:      boolPtr(false),
		Command:       strPtr(DefaultCommand),
		AllNamespaces: boolPtr(false),
		ReadOnly:      boolPtr(false),
		Write:         boolPtr(false),
		Crumbsless:    boolPtr(false),
		ScreenDumpDir: strPtr(AppDumpsDir),
	}
}

func boolPtr(b bool) *bool {
	return &b
}

func intPtr(i int) *int {
	return &i
}

func strPtr(s string) *string {
	return &s
}
