package slingshot

// PluginIter describes a plugin iterator.
type PluginIter interface {
	Next() *PluginMeta
}

// pluginIter is an implementation of PluginIter.
type pluginIter struct {
	plugins []*PluginMeta // The plugins to iterate over
	idx     int           // The next index to return
}

// IterPlugins is a constructor for a PluginIter.  Given a list of
// plugins, as returned by Namespace.GetAll, it returns an objects
// whose Next method will return each plugin in turn.
func IterPlugins(plugins []*PluginMeta) PluginIter {
	return &pluginIter{
		plugins: plugins,
		idx:     0,
	}
}

// Next returns the next plugin from the iterator.  If there are no
// more plugins in the iterator, it returns nil.  There is no
// wrap-around with this functionality.
func (it *pluginIter) Next() *PluginMeta {
	// Return nil if we've gone off the end of the list
	if it.idx >= len(it.plugins) {
		return nil
	}

	// Get the next plugin
	next := it.plugins[it.idx]

	// Increment the index
	it.idx++

	return next
}
