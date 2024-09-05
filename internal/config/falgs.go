package config

import "sync"

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

// ConfigFlags composes the set of values necessary
// for obtaining a REST client config
type ConfigFlags struct {
	CacheDir   *string
	KubeConfig *string

	// config flags
	ClusterName        *string
	AuthInfoName       *string
	Context            *string
	Namespace          *string
	APIServer          *string
	TLSServerName      *string
	Insecure           *bool
	CertFile           *string
	KeyFile            *string
	CAFile             *string
	BearerToken        *string
	Impersonate        *string
	ImpersonateUID     *string
	ImpersonateGroup   *[]string
	Username           *string
	Password           *string
	Timeout            *string
	DisableCompression *bool
	// If non-nil, wrap config function can transform the Config
	// before it is returned in ToRESTConfig function.
	//WrapConfigFn func(*rest.Config) *rest.Config

	//clientConfig     clientcmd.ClientConfig
	clientConfigLock sync.Mutex

	//restMapper     meta.RESTMapper
	restMapperLock sync.Mutex

	//discoveryClient     discovery.CachedDiscoveryInterface
	discoveryClientLock sync.Mutex

	// If set to true, will use persistent client config, rest mapper, discovery client, and
	// propagate them to the places that need them, rather than
	// instantiating them multiple times.
	usePersistentConfig bool
	// Allows increasing burst used for discovery, this is useful
	// in clusters with many registered resources
	discoveryBurst int
	// Allows increasing qps used for discovery, this is useful
	// in clusters with many registered resources
	discoveryQPS float32
	// Allows all possible warnings are printed in a standardized
	// format.
	//warningPrinter *printers.WarningPrinter
}
