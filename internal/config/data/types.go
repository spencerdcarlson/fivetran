package data

type KubeSettings interface {
	// CurrentContextName returns the name of the current context.
	CurrentContextName() (string, error)

	// CurrentClusterName returns the name of the current cluster.
	CurrentClusterName() (string, error)

	// CurrentNamespaceName returns the name of the current namespace.
	CurrentNamespaceName() (string, error)

	// ContextNames returns all available context names.
	ContextNames() (map[string]struct{}, error)
}
