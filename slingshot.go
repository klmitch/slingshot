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
