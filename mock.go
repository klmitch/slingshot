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
	"github.com/stretchr/testify/mock"
)

// MockRegistry is a mock object for Registry.
type MockRegistry struct {
	mock.Mock
}

// Get gets a specified namespace from the registry.  If the namespace
// doesn't have any entries and create is false, the second value will
// be false.
func (reg *MockRegistry) Get(namespace string, create bool) (Namespace, bool) {
	args := reg.MethodCalled("Get", namespace, create)

	ns := args.Get(0)
	if ns == nil {
		return nil, args.Bool(1)
	}
	return ns.(Namespace), args.Bool(1)
}

// GetPlugin gets a specified plugin from the designated namespace of
// the registry.  If the namespace doesn't have any entries for the
// designated key, the second value will be false.
func (reg *MockRegistry) GetPlugin(namespace, key string) (*PluginMeta, bool) {
	args := reg.MethodCalled("GetPlugin", namespace, key)

	meta := args.Get(0)
	if meta == nil {
		return nil, args.Bool(1)
	}
	return meta.(*PluginMeta), args.Bool(1)
}

// GetAllPlugins gets all the plugin descriptors for the designated
// namespace of the registry.  If the namespace doesn't have any
// entries for the designated key, the second value will be false.
func (reg *MockRegistry) GetAllPlugins(namespace, key string) ([]*PluginMeta, bool) {
	args := reg.MethodCalled("GetAllPlugins", namespace, key)

	meta := args.Get(0)
	if meta == nil {
		return nil, args.Bool(1)
	}
	return meta.([]*PluginMeta), args.Bool(1)
}

// Register is for registering a "core" plugin--that is, a plugin that
// is implemented within the code of the application, rather than one
// loaded from an external file using the plugin package.
func (reg *MockRegistry) Register(namespace, key string, plugin interface{}, opts ...PluginOption) {
	reg.MethodCalled("Register", namespace, key, plugin, newPluginMeta("", "", namespace, key, plugin, opts...))
}

// Load loads a plugin and instructs it to register its plugin points
// with the Slingshot registry.
func (reg *MockRegistry) Load(path string, params map[string]interface{}) error {
	args := reg.MethodCalled("Load", path, params)
	return args.Error(0)
}

// MockNamespace is a mock object for Namespace.
type MockNamespace struct {
	mock.Mock
	Name string
}

// Namespace returns the namespace string.  The Name element of the
// MockNamespace must be set for this method to work.
func (ns *MockNamespace) Namespace() string {
	return ns.Name
}

// Get returns the first plugin descriptor for the given key.  If
// there are no descriptors for that key, the second value will be
// false.
func (ns *MockNamespace) Get(key string) (*PluginMeta, bool) {
	args := ns.MethodCalled("Get", key)

	meta := args.Get(0)
	if meta == nil {
		return nil, args.Bool(1)
	}
	return meta.(*PluginMeta), args.Bool(1)
}

// GetAll returns all the plugin descriptors for the given key.  If
// there are no descriptors for that key, the second value will be
// false.
func (ns *MockNamespace) GetAll(key string) ([]*PluginMeta, bool) {
	args := ns.MethodCalled("GetAll", key)

	meta := args.Get(0)
	if meta == nil {
		return nil, args.Bool(1)
	}
	return meta.([]*PluginMeta), args.Bool(1)
}

// Add adds a new plugin descriptor under the given key.
func (ns *MockNamespace) Add(key string, plugin *PluginMeta) {
	ns.MethodCalled("Add", key, plugin)
}

// MockSlingshot is a mock object for Slingshot.
type MockSlingshot struct {
	mock.Mock
}

// Register is for registering a "core" plugin--that is, a plugin that
// is implemented within the code of the application, rather than one
// loaded from an external file using the plugin package.
func (sling *MockSlingshot) Register(namespace, key string, plugin interface{}, opts ...PluginOption) {
	sling.MethodCalled("Register", namespace, key, plugin, newPluginMeta("", "", namespace, key, plugin, opts...))
}

// SetRegistry sets the registry used by the top-level functions to
// the specified value, returning the original value.  This is
// intended for use by library callers to allow for mocking out the
// registry using the MockRegistry type.
func SetRegistry(newReg Registry) (oldReg Registry) {
	oldReg = reg
	reg = newReg
	return
}
