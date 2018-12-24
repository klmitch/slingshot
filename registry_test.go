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
	"plugin"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegistryImplementsRegistry(t *testing.T) {
	assert.Implements(t, (*Registry)(nil), &registry{})
}

func TestGetExists(t *testing.T) {
	a := assert.New(t)
	ns := &namespace{namespace: "name.space"}
	reg := &registry{
		namespaces: map[string]Namespace{
			"name.space": ns,
		},
	}

	result, ok := reg.Get("name.space", false)

	resNs := result.(*namespace)
	a.Equal(ns, resNs)
	a.True(ok)
	a.Equal(reg.namespaces, map[string]Namespace{
		"name.space": ns,
	})
}

func TestGetMissing(t *testing.T) {
	a := assert.New(t)
	reg := &registry{
		namespaces: map[string]Namespace{},
	}

	result, ok := reg.Get("name.space", false)

	a.Nil(result)
	a.False(ok)
	a.Equal(reg.namespaces, map[string]Namespace{})
}

func TestGetCreate(t *testing.T) {
	a := assert.New(t)
	reg := &registry{
		namespaces: map[string]Namespace{},
	}

	result, ok := reg.Get("name.space", true)

	resNs := result.(*namespace)
	a.Equal(resNs, newNamespace("name.space"))
	a.True(ok)
	a.Equal(reg.namespaces, map[string]Namespace{
		"name.space": resNs,
	})
}

func TestGetPluginNoNamespace(t *testing.T) {
	a := assert.New(t)
	reg := &registry{
		namespaces: map[string]Namespace{},
	}

	result, ok := reg.GetPlugin("name.space", "key")

	a.Nil(result)
	a.False(ok)
}

func TestGetPluginWithNamespace(t *testing.T) {
	a := assert.New(t)
	ns := &MockNamespace{}
	meta := &PluginMeta{Name: "plug"}
	reg := &registry{
		namespaces: map[string]Namespace{
			"name.space": ns,
		},
	}
	ns.On("Get", "key").Return(meta, true)

	result, ok := reg.GetPlugin("name.space", "key")

	a.Equal(result, meta)
	a.True(ok)
	ns.AssertExpectations(t)
}

func TestGetAllPluginsNoNamespace(t *testing.T) {
	a := assert.New(t)
	reg := &registry{
		namespaces: map[string]Namespace{},
	}

	result, ok := reg.GetAllPlugins("name.space", "key")

	a.Equal(result, []*PluginMeta{})
	a.False(ok)
}

func TestGetAllPluginsWithNamespace(t *testing.T) {
	a := assert.New(t)
	ns := &MockNamespace{}
	expected := []*PluginMeta{
		&PluginMeta{Name: "plug1"},
		&PluginMeta{Name: "plug2"},
	}
	reg := &registry{
		namespaces: map[string]Namespace{
			"name.space": ns,
		},
	}
	ns.On("GetAll", "key").Return(expected, true)

	result, ok := reg.GetAllPlugins("name.space", "key")

	a.Equal(result, expected)
	a.True(ok)
	ns.AssertExpectations(t)
}

func TestRegister(t *testing.T) {
	ns := &MockNamespace{}
	reg := &registry{
		namespaces: map[string]Namespace{
			"name.space": ns,
		},
	}
	ns.On("Add", "key", newPluginMeta("", "", "name.space", "key", "plugin"))

	reg.Register("name.space", "key", "plugin")

	ns.AssertExpectations(t)
}

func setLoadHooks(newAbsHook absHookType, newBaseHook baseHookType, newOpenHook openHookType) (absHookType, baseHookType, openHookType) {
	// Get the originals
	origAbsHook := absHook
	origBaseHook := baseHook
	origOpenHook := openHook

	// Set the new ones
	absHook = newAbsHook
	baseHook = newBaseHook
	openHook = newOpenHook

	// Return the originals
	return origAbsHook, origBaseHook, origOpenHook
}

type mockPlugin struct {
	mock.Mock
}

func (plug *mockPlugin) Lookup(symName string) (plugin.Symbol, error) {
	args := plug.MethodCalled("Lookup", symName)

	sym := args.Get(0)
	if sym == nil {
		return nil, args.Error(1)
	}
	return sym.(plugin.Symbol), args.Error(1)
}

func TestLoadAbsFails(t *testing.T) {
	a := assert.New(t)
	origAbsHook, origBaseHook, origOpenHook := setLoadHooks(
		func(path string) (string, error) {
			a.Equal(path, "orig/path")

			return "", errors.New("Abs failed")
		},
		func(path string) string {
			a.Equal(path, "/full/path.so")

			return "path.so"
		},
		func(path string) (pluginInterface, error) {
			a.Equal(path, "/full/path.so")

			return &mockPlugin{}, nil
		},
	)
	defer setLoadHooks(origAbsHook, origBaseHook, origOpenHook)
	reg := &registry{
		namespaces: map[string]Namespace{
			"name.space": &namespace{namespace: "name.space"},
		},
	}

	err := reg.Load("orig/path", nil)

	a.EqualError(err, "Abs failed")
}

func TestLoadOpenFails(t *testing.T) {
	a := assert.New(t)
	origAbsHook, origBaseHook, origOpenHook := setLoadHooks(
		func(path string) (string, error) {
			a.Equal(path, "orig/path")

			return "/full/path.so", nil
		},
		func(path string) string {
			a.Equal(path, "/full/path.so")

			return "path.so"
		},
		func(path string) (pluginInterface, error) {
			a.Equal(path, "/full/path.so")

			return nil, errors.New("Open failed")
		},
	)
	defer setLoadHooks(origAbsHook, origBaseHook, origOpenHook)
	reg := &registry{
		namespaces: map[string]Namespace{
			"name.space": &namespace{namespace: "name.space"},
		},
	}

	err := reg.Load("orig/path", nil)

	a.EqualError(err, "Open failed")
}

func TestLoadLookupFails(t *testing.T) {
	a := assert.New(t)
	plug := &mockPlugin{}
	origAbsHook, origBaseHook, origOpenHook := setLoadHooks(
		func(path string) (string, error) {
			a.Equal(path, "orig/path")

			return "/full/path.so", nil
		},
		func(path string) string {
			a.Equal(path, "/full/path.so")

			return "path.so"
		},
		func(path string) (pluginInterface, error) {
			a.Equal(path, "/full/path.so")

			return plug, nil
		},
	)
	defer setLoadHooks(origAbsHook, origBaseHook, origOpenHook)
	plug.On("Lookup", SlingshotInit).Return(nil, errors.New("Lookup failed"))
	reg := &registry{
		namespaces: map[string]Namespace{
			"name.space": &namespace{namespace: "name.space"},
		},
	}

	err := reg.Load("orig/path", nil)

	a.EqualError(err, "Lookup failed")
	plug.AssertExpectations(t)
}

func TestLoadCastFails(t *testing.T) {
	a := assert.New(t)
	plug := &mockPlugin{}
	origAbsHook, origBaseHook, origOpenHook := setLoadHooks(
		func(path string) (string, error) {
			a.Equal(path, "orig/path")

			return "/full/path.so", nil
		},
		func(path string) string {
			a.Equal(path, "/full/path.so")

			return "path.so"
		},
		func(path string) (pluginInterface, error) {
			a.Equal(path, "/full/path.so")

			return plug, nil
		},
	)
	defer setLoadHooks(origAbsHook, origBaseHook, origOpenHook)
	plug.On("Lookup", SlingshotInit).Return("value", nil)
	reg := &registry{
		namespaces: map[string]Namespace{
			"name.space": &namespace{namespace: "name.space"},
		},
	}

	err := reg.Load("orig/path", nil)

	a.EqualError(err, "Incompatible plugin initializer")
	plug.AssertExpectations(t)
}

func TestLoadInitFnFails(t *testing.T) {
	a := assert.New(t)
	plug := &mockPlugin{}
	origAbsHook, origBaseHook, origOpenHook := setLoadHooks(
		func(path string) (string, error) {
			a.Equal(path, "orig/path")

			return "/full/path.so", nil
		},
		func(path string) string {
			a.Equal(path, "/full/path.so")

			return "path.so"
		},
		func(path string) (pluginInterface, error) {
			a.Equal(path, "/full/path.so")

			return plug, nil
		},
	)
	defer setLoadHooks(origAbsHook, origBaseHook, origOpenHook)
	reg := &registry{
		namespaces: map[string]Namespace{
			"name.space": &namespace{namespace: "name.space"},
		},
	}
	initFn := func(sling Slingshot, params map[string]interface{}) error {
		ss := sling.(*slingshot)
		a.Equal(ss, &slingshot{
			registry: reg,
			path:     "/full/path.so",
			filename: "path.so",
		})
		a.Nil(params)

		return errors.New("InitFn fails")
	}
	plug.On("Lookup", SlingshotInit).Return(initFn, nil)

	err := reg.Load("orig/path", nil)

	a.EqualError(err, "InitFn fails")
	plug.AssertExpectations(t)
}

func TestLoadInitFnPanics(t *testing.T) {
	a := assert.New(t)
	plug := &mockPlugin{}
	origAbsHook, origBaseHook, origOpenHook := setLoadHooks(
		func(path string) (string, error) {
			a.Equal(path, "orig/path")

			return "/full/path.so", nil
		},
		func(path string) string {
			a.Equal(path, "/full/path.so")

			return "path.so"
		},
		func(path string) (pluginInterface, error) {
			a.Equal(path, "/full/path.so")

			return plug, nil
		},
	)
	defer setLoadHooks(origAbsHook, origBaseHook, origOpenHook)
	reg := &registry{
		namespaces: map[string]Namespace{
			"name.space": &namespace{namespace: "name.space"},
		},
	}
	initFn := func(sling Slingshot, params map[string]interface{}) error {
		ss := sling.(*slingshot)
		a.Equal(ss, &slingshot{
			registry: reg,
			path:     "/full/path.so",
			filename: "path.so",
		})
		a.Nil(params)

		panic("panic my initializer")

		return nil
	}
	plug.On("Lookup", SlingshotInit).Return(initFn, nil)

	err := reg.Load("orig/path", nil)

	a.EqualError(err, "Plugin initializer paniced")
	plug.AssertExpectations(t)
}

func TestLoadInitFnGetsParams(t *testing.T) {
	a := assert.New(t)
	plug := &mockPlugin{}
	origAbsHook, origBaseHook, origOpenHook := setLoadHooks(
		func(path string) (string, error) {
			a.Equal(path, "orig/path")

			return "/full/path.so", nil
		},
		func(path string) string {
			a.Equal(path, "/full/path.so")

			return "path.so"
		},
		func(path string) (pluginInterface, error) {
			a.Equal(path, "/full/path.so")

			return plug, nil
		},
	)
	defer setLoadHooks(origAbsHook, origBaseHook, origOpenHook)
	reg := &registry{
		namespaces: map[string]Namespace{
			"name.space": &namespace{namespace: "name.space"},
		},
	}
	initFn := func(sling Slingshot, params map[string]interface{}) error {
		ss := sling.(*slingshot)
		a.Equal(ss, &slingshot{
			registry: reg,
			path:     "/full/path.so",
			filename: "path.so",
		})
		a.Equal(params, map[string]interface{}{
			"a": "value",
			"b": 3,
		})

		return nil
	}
	plug.On("Lookup", SlingshotInit).Return(initFn, nil)

	err := reg.Load("orig/path", map[string]interface{}{
		"a": "value",
		"b": 3,
	})

	a.NoError(err)
	plug.AssertExpectations(t)
}
