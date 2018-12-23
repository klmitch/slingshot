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

	a.Nil(result)
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
