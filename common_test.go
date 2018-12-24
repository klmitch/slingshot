package slingshot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopGet(t *testing.T) {
	a := assert.New(t)
	ns := &namespace{namespace: "name.space"}
	reg := &MockRegistry{}
	defer SetRegistry(SetRegistry(reg))
	reg.On("Get", "name.space", false).Return(ns, true)

	result, ok := Get("name.space", false)

	a.Equal(result, ns)
	a.True(ok)
	reg.AssertExpectations(t)
}

func TestTopGetPlugin(t *testing.T) {
	a := assert.New(t)
	meta := newPluginMeta("", "", "name.space", "key", "plugin")
	reg := &MockRegistry{}
	defer SetRegistry(SetRegistry(reg))
	reg.On("GetPlugin", "name.space", "key").Return(meta, true)

	result, ok := GetPlugin("name.space", "key")

	a.Equal(result, meta)
	a.True(ok)
	reg.AssertExpectations(t)
}

func TestTopGetAllPlugins(t *testing.T) {
	a := assert.New(t)
	plugs := []*PluginMeta{
		newPluginMeta("", "", "name.space", "key", "plugin", Name("Plug1")),
		newPluginMeta("", "", "name.space", "key", "plugin", Name("Plug2")),
	}
	reg := &MockRegistry{}
	defer SetRegistry(SetRegistry(reg))
	reg.On("GetAllPlugins", "name.space", "key").Return(plugs, true)

	result, ok := GetAllPlugins("name.space", "key")

	a.Equal(result, plugs)
	a.True(ok)
	reg.AssertExpectations(t)
}

func TestTopRegister(t *testing.T) {
	reg := &MockRegistry{}
	defer SetRegistry(SetRegistry(reg))
	reg.On("Register", "name.space", "key", "plugin", newPluginMeta("", "", "name.space", "key", "plugin"))

	Register("name.space", "key", "plugin")

	reg.AssertExpectations(t)
}

func TestTopLoad(t *testing.T) {
	a := assert.New(t)
	reg := &MockRegistry{}
	defer SetRegistry(SetRegistry(reg))
	reg.On("Load", "some/path", map[string]interface{}{
		"a": "value",
		"b": 3,
	}).Return(nil)

	err := Load("some/path", map[string]interface{}{
		"a": "value",
		"b": 3,
	})

	a.NoError(err)
	reg.AssertExpectations(t)
}
