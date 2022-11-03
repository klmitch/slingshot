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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockRegistryImplementsRegistry(t *testing.T) {
	assert.Implements(t, (*Registry)(nil), &MockRegistry{})
}

func TestMockNamespaceImplementsNamespace(t *testing.T) {
	assert.Implements(t, (*Namespace)(nil), &MockNamespace{})
}

func TestMockSlingshotImplementsSlingshot(t *testing.T) {
	assert.Implements(t, (*Slingshot)(nil), &MockSlingshot{})
}

func TestMockRegistryGetNil(t *testing.T) {
	a := assert.New(t)
	reg := &MockRegistry{}
	reg.On("Get", "name.space", true).Return(nil, false)

	result, ok := reg.Get("name.space", true)

	a.Nil(result)
	a.False(ok)
	reg.AssertExpectations(t)
}

func TestMockRegistryGetNonNil(t *testing.T) {
	a := assert.New(t)
	reg := &MockRegistry{}
	ns := &MockNamespace{}
	reg.On("Get", "name.space", true).Return(ns, true)

	result, ok := reg.Get("name.space", true)

	a.Equal(result, ns)
	a.True(ok)
	reg.AssertExpectations(t)
}

func TestMockRegistryGetPluginNil(t *testing.T) {
	a := assert.New(t)
	reg := &MockRegistry{}
	reg.On("GetPlugin", "name.space", "key").Return(nil, false)

	result, ok := reg.GetPlugin("name.space", "key")

	a.Nil(result)
	a.False(ok)
	reg.AssertExpectations(t)
}

func TestMockRegistryGetPluginNonNil(t *testing.T) {
	a := assert.New(t)
	reg := &MockRegistry{}
	meta := &PluginMeta{}
	reg.On("GetPlugin", "name.space", "key").Return(meta, true)

	result, ok := reg.GetPlugin("name.space", "key")

	a.Equal(result, meta)
	a.True(ok)
	reg.AssertExpectations(t)
}

func TestMockRegistryGetAllPluginsNil(t *testing.T) {
	a := assert.New(t)
	reg := &MockRegistry{}
	reg.On("GetAllPlugins", "name.space", "key").Return(nil, false)

	result, ok := reg.GetAllPlugins("name.space", "key")

	a.Nil(result)
	a.False(ok)
	reg.AssertExpectations(t)
}

func TestMockRegistryGetAllPluginsNonNil(t *testing.T) {
	a := assert.New(t)
	reg := &MockRegistry{}
	plugs := []*PluginMeta{
		{},
		{},
	}
	reg.On("GetAllPlugins", "name.space", "key").Return(plugs, true)

	result, ok := reg.GetAllPlugins("name.space", "key")

	a.Equal(result, plugs)
	a.True(ok)
	reg.AssertExpectations(t)
}

func TestMockRegistryRegisterBase(t *testing.T) {
	reg := &MockRegistry{}
	reg.On("Register", "name.space", "key", "plugin", &PluginMeta{
		Namespace: "name.space",
		Key:       "key",
		Plugin:    "plugin",
		Meta:      map[string]interface{}{},
	})

	reg.Register("name.space", "key", "plugin")

	reg.AssertExpectations(t)
}

func TestMockRegistryRegisterOptions(t *testing.T) {
	reg := &MockRegistry{}
	reg.On("Register", "name.space", "key", "plugin", &PluginMeta{
		Namespace: "name.space",
		Key:       "key",
		Plugin:    "plugin",
		Name:      "plug",
		Meta:      map[string]interface{}{},
	})

	reg.Register("name.space", "key", "plugin", Name("plug"))

	reg.AssertExpectations(t)
}

func TestMockRegistryLoad(t *testing.T) {
	a := assert.New(t)
	reg := &MockRegistry{}
	reg.On("Load", "some/path.so", map[string]interface{}{
		"a": "value",
		"b": 3,
	}).Return(errors.New("an error")) //nolint:goerr113

	err := reg.Load("some/path.so", map[string]interface{}{
		"a": "value",
		"b": 3,
	})

	a.EqualError(err, "an error")
	reg.AssertExpectations(t)
}

func TestMockNamespaceNamespace(t *testing.T) {
	a := assert.New(t)
	ns := &MockNamespace{Name: "ns"}

	result := ns.Namespace()

	a.Equal(result, "ns")
}

func TestMockNamespaceGetNil(t *testing.T) {
	a := assert.New(t)
	ns := &MockNamespace{}
	ns.On("Get", "key").Return(nil, false)

	result, ok := ns.Get("key")

	a.Nil(result)
	a.False(ok)
	ns.AssertExpectations(t)
}

func TestMockNamespaceGetNonNil(t *testing.T) {
	a := assert.New(t)
	ns := &MockNamespace{}
	meta := &PluginMeta{}
	ns.On("Get", "key").Return(meta, true)

	result, ok := ns.Get("key")

	a.Equal(result, meta)
	a.True(ok)
	ns.AssertExpectations(t)
}

func TestMockNamespaceGetAllNil(t *testing.T) {
	a := assert.New(t)
	ns := &MockNamespace{}
	ns.On("GetAll", "key").Return(nil, false)

	result, ok := ns.GetAll("key")

	a.Nil(result)
	a.False(ok)
	ns.AssertExpectations(t)
}

func TestMockNamespaceGetAllNonNil(t *testing.T) {
	a := assert.New(t)
	ns := &MockNamespace{}
	plugs := []*PluginMeta{
		{},
		{},
	}
	ns.On("GetAll", "key").Return(plugs, true)

	result, ok := ns.GetAll("key")

	a.Equal(result, plugs)
	a.True(ok)
	ns.AssertExpectations(t)
}

func TestMockNamespaceAdd(t *testing.T) {
	ns := &MockNamespace{}
	plug := &PluginMeta{}
	ns.On("Add", "key", plug)

	ns.Add("key", plug)

	ns.AssertExpectations(t)
}

func TestMockSlingshotRegisterBase(t *testing.T) {
	sling := &MockSlingshot{}
	sling.On("Register", "name.space", "key", "plugin", &PluginMeta{
		Namespace: "name.space",
		Key:       "key",
		Plugin:    "plugin",
		Meta:      map[string]interface{}{},
	})

	sling.Register("name.space", "key", "plugin")

	sling.AssertExpectations(t)
}

func TestMockSlingshotRegisterOptions(t *testing.T) {
	sling := &MockSlingshot{}
	sling.On("Register", "name.space", "key", "plugin", &PluginMeta{
		Namespace: "name.space",
		Key:       "key",
		Plugin:    "plugin",
		Name:      "plug",
		Meta:      map[string]interface{}{},
	})

	sling.Register("name.space", "key", "plugin", Name("plug"))

	sling.AssertExpectations(t)
}

func TestSetRegistry(t *testing.T) {
	a := assert.New(t)
	expected := reg
	defer func() {
		reg = expected
	}()
	newReg := &MockRegistry{}

	result := SetRegistry(newReg)

	a.Equal(result, expected)
	a.Equal(reg, newReg)
}
