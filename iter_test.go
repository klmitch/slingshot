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
		&PluginMeta{Name: "plug1"},
		&PluginMeta{Name: "plug2"},
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
			&PluginMeta{Name: "plug1"},
			&PluginMeta{Name: "plug2"},
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
			&PluginMeta{Name: "plug1"},
			&PluginMeta{Name: "plug2"},
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
			&PluginMeta{Name: "plug1"},
			&PluginMeta{Name: "plug2"},
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
