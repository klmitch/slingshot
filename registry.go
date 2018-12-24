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

import (
	"errors"
	"path/filepath"
	"plugin"
	"sync"
)

// Registry describes the Slingshot registry.
type Registry interface {
	Get(namespace string, create bool) (Namespace, bool)
	GetPlugin(namespace, key string) (*PluginMeta, bool)
	GetAllPlugins(namespace, key string) ([]*PluginMeta, bool)
	Register(namespace, key string, plugin interface{}, opts ...PluginOption)
	Load(path string, params map[string]interface{}) error
}

// registry is an implementation of Registry which incorporates
// locking--allowing safe access from multiple threads.
type registry struct {
	sync.Mutex                      // Mutex protecting the map
	namespaces map[string]Namespace // Map of namespaces
}

// Get gets a specified namespace from the registry.  If the namespace
// doesn't have any entries and create is false, the second value will
// be false.
func (reg *registry) Get(namespace string, create bool) (Namespace, bool) {
	// Lock the mutex around the registry
	reg.Lock()
	defer reg.Unlock()

	// Get and return the namespace
	ns, ok := reg.namespaces[namespace]
	if !ok {
		if !create {
			return nil, false
		}

		// Create the namespace
		ns = newNamespace(namespace)
		reg.namespaces[namespace] = ns
	}

	return ns, true
}

// GetPlugin gets a specified plugin from the designated namespace of
// the registry.  If the namespace doesn't have any entries for the
// designated key, the second value will be false.
func (reg *registry) GetPlugin(namespace, key string) (*PluginMeta, bool) {
	// Lock the mutex around the registry
	reg.Lock()
	defer reg.Unlock()

	// Get the namespace
	ns, ok := reg.namespaces[namespace]
	if !ok {
		return nil, false
	}

	return ns.Get(key)
}

// GetAllPlugins gets all the plugin descriptors for the designated
// namespace of the registry.  If the namespace doesn't have any
// entries for the designated key, the second value will be false.
func (reg *registry) GetAllPlugins(namespace, key string) ([]*PluginMeta, bool) {
	// Lock the mutex around the registry
	reg.Lock()
	defer reg.Unlock()

	// Get the namespace
	ns, ok := reg.namespaces[namespace]
	if !ok {
		return []*PluginMeta{}, false
	}

	return ns.GetAll(key)
}

// Register is for registering a "core" plugin--that is, a plugin that
// is implemented within the code of the application, rather than one
// loaded from an external file using the plugin package.
func (reg *registry) Register(namespace, key string, plugin interface{}, opts ...PluginOption) {
	// First, get (or create) the namespace
	ns, _ := reg.Get(namespace, true)

	// Next, construct the plugin metadata
	meta := newPluginMeta("", "", namespace, key, plugin, opts...)

	// Finally, add the plugin metadata to the namespace
	ns.Add(key, meta)
}

// pluginInterface is an interface used for testing the Load method.
type pluginInterface interface {
	Lookup(symName string) (plugin.Symbol, error)
}

// Types for the hooks
type absHookType func(string) (string, error)
type baseHookType func(string) string
type openHookType func(string) (pluginInterface, error)

// These internal variables allow mocking out the system-dependent
// functions filepath.Abs, filepath.Base, and plugin.Open within the
// tests.
var (
	absHook  absHookType  = filepath.Abs
	baseHook baseHookType = filepath.Base
	openHook openHookType = func(path string) (pluginInterface, error) {
		return plugin.Open(path)
	}
)

// Load loads a plugin and instructs it to register its plugin points
// with the Slingshot registry.
func (reg *registry) Load(path string, params map[string]interface{}) (err error) {
	// Begin by resolving the path
	path, err = absHook(path)
	if err != nil {
		return
	}
	filename := baseHook(path)

	// Open the plugin
	plug, err := openHook(path)
	if err != nil {
		return
	}

	// Look up the initializer function
	initSym, err := plug.Lookup(SlingshotInit)
	if err != nil {
		return
	}
	initFn, ok := initSym.(func(Slingshot, map[string]interface{}) error)
	if !ok {
		return errors.New("Incompatible plugin initializer")
	}

	// OK, construct the slingshot and call the initializer
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Plugin initializer paniced")
		}
	}()
	return initFn(&slingshot{
		registry: reg,
		path:     path,
		filename: filename,
	}, params)
}
