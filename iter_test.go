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

func TestPluginIterImplementsPluginIter(t *testing.T) {
	assert.Implements(t, (*PluginIter)(nil), &pluginIter{})
}

func TestIterPlugins(t *testing.T) {
	a := assert.New(t)
	plugins := []*PluginMeta{
		{Name: "plug1"},
		{Name: "plug2"},
	}

	result := IterPlugins(plugins)

	iter := result.(*pluginIter)
	a.Equal(iter.plugins, plugins)
	a.Equal(iter.idx, 0)
}

func TestNextBegin(t *testing.T) {
	a := assert.New(t)
	iter := &pluginIter{
		plugins: []*PluginMeta{
			{Name: "plug1"},
			{Name: "plug2"},
		},
		idx: 0,
	}

	result := iter.Next()

	a.Equal(result, iter.plugins[0])
	a.Equal(iter.idx, 1)
}

func TestNextLast(t *testing.T) {
	a := assert.New(t)
	iter := &pluginIter{
		plugins: []*PluginMeta{
			{Name: "plug1"},
			{Name: "plug2"},
		},
		idx: 1,
	}

	result := iter.Next()

	a.Equal(result, iter.plugins[1])
	a.Equal(iter.idx, 2)
}

func TestNextEnd(t *testing.T) {
	a := assert.New(t)
	iter := &pluginIter{
		plugins: []*PluginMeta{
			{Name: "plug1"},
			{Name: "plug2"},
		},
		idx: 2,
	}

	result := iter.Next()

	a.Nil(result)
	a.Equal(iter.idx, 2)
}

func TestNextEmpty(t *testing.T) {
	a := assert.New(t)
	iter := &pluginIter{
		plugins: []*PluginMeta{},
		idx:     0,
	}

	result := iter.Next()

	a.Nil(result)
	a.Equal(iter.idx, 0)
}
