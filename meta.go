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

// PluginMeta contains the metadata for a registered plugin.  This
// includes such things as the filename, the namespace (string), the
// key, the plugin version, the API version, or anything else the
// plugin may register.
type PluginMeta struct {
	Path       string                 // Path to the plugin
	Filename   string                 // Basename of the plugin
	Namespace  string                 // Namespace the plugin exists in
	Key        string                 // The key the plugin belongs in
	Plugin     interface{}            // The actual plugin.
	Name       string                 // The name the plugin identifies as
	Version    string                 // The version of the plugin
	License    string                 // Text describing the plugin license
	Docs       string                 // Documentation for the plugin
	APIVersion int                    // The API version of the plugin
	Meta       map[string]interface{} // Additional metadata
}

// PluginOption is an option function that can be passed to the
// Plugin.Register method.
type PluginOption func(meta *PluginMeta)

// Name sets the name of the plugin, which may be distinct from the
// filename or key depending on the needs of the application.
func Name(name string) PluginOption {
	return func(meta *PluginMeta) {
		meta.Name = name
	}
}

// Version sets the version of the plugin.
func Version(version string) PluginOption {
	return func(meta *PluginMeta) {
		meta.Version = version
	}
}

// License sets the license descriptor for the plugin.
func License(license string) PluginOption {
	return func(meta *PluginMeta) {
		meta.License = license
	}
}

// Docs sets the documentation string for the plugin.
func Docs(docs string) PluginOption {
	return func(meta *PluginMeta) {
		meta.Docs = docs
	}
}

// APIVersion sets the version of the plugin API.  This is a number
// that can be used by the application to determine how to call the
// plugin, if the plugin interface has been evolved.
func APIVersion(version int) PluginOption {
	return func(meta *PluginMeta) {
		meta.APIVersion = version
	}
}

// Meta allows setting any number of other pieces of metadata, which
// can be used by the application as desired.
func Meta(key string, value interface{}) PluginOption {
	return func(meta *PluginMeta) {
		meta.Meta[key] = value
	}
}

// newPluginMeta constructs a new PluginMeta instance with all the
// passed-in data.
func newPluginMeta(path, fname, namespace, key string, plugin interface{}, opts ...PluginOption) *PluginMeta {
	// Construct the basic metadata
	meta := &PluginMeta{
		Path:      path,
		Filename:  fname,
		Namespace: namespace,
		Key:       key,
		Plugin:    plugin,
		Meta:      map[string]interface{}{},
	}

	// Apply all options
	for _, opt := range opts {
		opt(meta)
	}

	return meta
}
