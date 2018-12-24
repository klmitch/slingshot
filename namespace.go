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
	"sync"
)

// Namespace describes a Slingshot namespace.
type Namespace interface {
	Namespace() string
	Get(key string) (*PluginMeta, bool)
	GetAll(key string) ([]*PluginMeta, bool)
	Add(key string, plugin *PluginMeta)
}

// namespace is an implementation of Namespace which incorporates
// locking--allowing safe access from multiple threads.
type namespace struct {
	sync.Mutex                          // Mutex protecting the map
	namespace  string                   // The name of the namespace
	contents   map[string][]*PluginMeta // Contents of the namespace
}

// Namespace returns the namespace string of the namespace.
func (ns *namespace) Namespace() string {
	return ns.namespace
}

// Get returns the first plugin descriptor for the given key.  If
// there are no descriptors for that key, the second value will be
// false.
func (ns *namespace) Get(key string) (*PluginMeta, bool) {
	// Lock the mutex around the namespace
	ns.Lock()
	defer ns.Unlock()

	// Get the plugins for the key
	plugs, ok := ns.contents[key]
	if !ok {
		return nil, false
	}

	return plugs[0], true
}

// GetAll returns all the plugin descriptors for the given key.  If
// there are no descriptors for that key, the second value will be
// false.
func (ns *namespace) GetAll(key string) ([]*PluginMeta, bool) {
	// Lock the mutex around the namespace
	ns.Lock()
	defer ns.Unlock()

	// Get the plugins for the key
	plugs, ok := ns.contents[key]
	if !ok {
		return []*PluginMeta{}, false
	}

	// Make a point-in-time result
	result := make([]*PluginMeta, len(plugs))
	copy(result, plugs)

	return result, true
}

// Add adds a new plugin descriptor under the given key.
func (ns *namespace) Add(key string, plugin *PluginMeta) {
	// Lock the mutex around the namespace
	ns.Lock()
	defer ns.Unlock()

	// Check if we need to initialize the contents
	_, ok := ns.contents[key]
	if !ok {
		ns.contents[key] = []*PluginMeta{}
	}

	// Add the plugin
	ns.contents[key] = append(ns.contents[key], plugin)
}

// newNamespace constructs a new namespace with the designated name.
func newNamespace(name string) Namespace {
	return &namespace{
		namespace: name,
		contents:  map[string][]*PluginMeta{},
	}
}
