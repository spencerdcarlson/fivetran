package client

const (
	NA = "n/a"
)

type Connection interface {
	//Authorizer

	// Config returns current config.
	//Config() *Config

	// ConnectionOK checks api server connection status.
	ConnectionOK() bool

	// Dial connects to api server.
	//Dial() (kubernetes.Interface, error)

	// DialLogs connects to api server for logs.
	//DialLogs() (kubernetes.Interface, error)

	// SwitchContext switches cluster based on context.
	SwitchContext(ctx string) error

	// CachedDiscovery connects to discovery client.
	//CachedDiscovery() (*disk.CachedDiscoveryClient, error)

	// RestConfig connects to rest client.
	//RestConfig() (*restclient.Config, error)

	// MXDial connects to metrics server.
	//MXDial() (*versioned.Clientset, error)

	// DynDial connects to dynamic client.
	//DynDial() (dynamic.Interface, error)

	// HasMetrics checks if metrics server is available.
	HasMetrics() bool

	// ValidNamespaceNames returns all available namespace names.
	//ValidNamespaceNames() (NamespaceNames, error)

	// IsValidNamespace checks if given namespace is known.
	IsValidNamespace(string) bool

	// ServerVersion returns current server version.
	//ServerVersion() (*version.Info, error)

	// CheckConnectivity checks if api server connection is happy or not.
	CheckConnectivity() bool

	// ActiveContext returns the current context name.
	ActiveContext() string

	// ActiveNamespace returns the current namespace.
	ActiveNamespace() string

	// IsActiveNamespace checks if given ns is active.
	IsActiveNamespace(string) bool
}
