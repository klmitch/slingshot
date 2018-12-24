package slingshot

// SlingshotInit is the name of the plugin initialization function
// that will be looked up.
const SlingshotInit = "SlingshotInit"

// reg is the single registry.
var reg Registry = &registry{namespaces: map[string]Namespace{}}

// Get gets a specified namespace from the registry.  If the namespace
// doesn't have any entries and create is false, the second value will
// be false.
func Get(namespace string, create bool) (Namespace, bool) {
	return reg.Get(namespace, create)
}

// GetPlugin gets a specified plugin from the designated namespace of
// the registry.  If the namespace doesn't have any entries for the
// designated key, the second value will be false.
func GetPlugin(namespace, key string) (*PluginMeta, bool) {
	return reg.GetPlugin(namespace, key)
}

// GetAllPlugins gets all the plugin descriptors for the designated
// namespace of the registry.  If the namespace doesn't have any
// entries for the designated key, the second value will be false.
func GetAllPlugins(namespace, key string) ([]*PluginMeta, bool) {
	return reg.GetAllPlugins(namespace, key)
}

// Register is for registering a "core" plugin--that is, a plugin that
// is implemented within the code of the application, rather than one
// loaded from an external file using the plugin package.
func Register(namespace, key string, plugin interface{}, opts ...PluginOption) {
	reg.Register(namespace, key, plugin, opts...)
}

// Load loads a plugin and instructs it to register its plugin points
// with the Slingshot registry.
func Load(path string, params map[string]interface{}) error {
	return reg.Load(path, params)
}
