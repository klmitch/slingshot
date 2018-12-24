package slingshot

// Slingshot is used to carry additional data through the plugin's
// initialization routine to the Register method.
type Slingshot interface {
	Register(namespace, key string, plugin interface{}, opts ...PluginOption)
}

// slingshot is an implementation of Slingshot which contains the key
// bits of data that need to be carried through.
type slingshot struct {
	registry Registry // The Slingshot registry
	path     string   // Full path to the plugin
	filename string   // Basename of the plugin
}

// Register is for registering a plugin extension point.
func (sling *slingshot) Register(namespace, key string, plugin interface{}, opts ...PluginOption) {
	// First, get (or create) the namespace
	ns, _ := sling.registry.Get(namespace, true)

	// Next, construct the plugin metadata
	meta := newPluginMeta(sling.path, sling.filename, namespace, key, plugin, opts...)

	// Finally, add the plugin metadata to the namespace
	ns.Add(key, meta)
}
