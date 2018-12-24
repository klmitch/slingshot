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
