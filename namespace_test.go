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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamespaceImplementsNamespace(t *testing.T) {
	assert.Implements(t, (*Namespace)(nil), &namespace{})
}

func TestNamespace(t *testing.T) {
	a := assert.New(t)
	ns := &namespace{namespace: "namespace"}

	result := ns.Namespace()

	a.Equal(result, "namespace")
}

func TestGetEmpty(t *testing.T) {
	a := assert.New(t)
	ns := &namespace{contents: map[string][]*PluginMeta{}}

	result, ok := ns.Get("key")

	a.Nil(result)
	a.False(ok)
}

func TestGetNonempty(t *testing.T) {
	a := assert.New(t)
	plug1 := &PluginMeta{}
	ns := &namespace{contents: map[string][]*PluginMeta{
		"key": []*PluginMeta{plug1},
	}}

	result, ok := ns.Get("key")

	a.Equal(result, plug1)
	a.True(ok)
}

func TestGetAllEmpty(t *testing.T) {
	a := assert.New(t)
	ns := &namespace{contents: map[string][]*PluginMeta{}}

	result, ok := ns.GetAll("key")

	a.Equal(result, []*PluginMeta{})
	a.False(ok)
}

func TestGetAllNonempty(t *testing.T) {
	a := assert.New(t)
	plugs := []*PluginMeta{
		&PluginMeta{Name: "plug1"},
		&PluginMeta{Name: "plug2"},
	}
	ns := &namespace{contents: map[string][]*PluginMeta{
		"key": plugs,
	}}

	result, ok := ns.GetAll("key")

	a.Equal(result, plugs)
	a.True(ok)
}

func TestAddEmpty(t *testing.T) {
	a := assert.New(t)
	plug1 := &PluginMeta{}
	ns := &namespace{contents: map[string][]*PluginMeta{}}

	ns.Add("key", plug1)

	a.Equal(ns.contents, map[string][]*PluginMeta{
		"key": []*PluginMeta{plug1},
	})
}

func TestAddNonempty(t *testing.T) {
	a := assert.New(t)
	plugs := []*PluginMeta{
		&PluginMeta{Name: "plug1"},
		&PluginMeta{Name: "plug2"},
	}
	ns := &namespace{contents: map[string][]*PluginMeta{
		"key": []*PluginMeta{plugs[0]},
	}}

	ns.Add("key", plugs[1])

	a.Equal(ns.contents, map[string][]*PluginMeta{"key": plugs})
}

func TestNewNamespace(t *testing.T) {
	a := assert.New(t)

	result := newNamespace("name.space")

	ns := result.(*namespace)
	a.Equal(ns.namespace, "name.space")
	a.Equal(ns.contents, map[string][]*PluginMeta{})
}
