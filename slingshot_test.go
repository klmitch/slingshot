package slingshot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlingshotImplementsSlingshot(t *testing.T) {
	assert.Implements(t, (*Slingshot)(nil), &slingshot{})
}

func TestSlingshotRegister(t *testing.T) {
	ns := &MockNamespace{}
	reg := &MockRegistry{}
	sling := &slingshot{
		registry: reg,
		path:     "/full/path.so",
		filename: "path.so",
	}
	reg.On("Get", "name.space", true).Return(ns, true)
	ns.On("Add", "key", newPluginMeta("/full/path.so", "path.so", "name.space", "key", "plugin"))

	sling.Register("name.space", "key", "plugin")

	ns.AssertExpectations(t)
	reg.AssertExpectations(t)
}
