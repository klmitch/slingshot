// Copyright 2018 Kevin L. Mitchell
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
