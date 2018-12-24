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

func TestName(t *testing.T) {
	a := assert.New(t)
	meta := &PluginMeta{}

	opt := Name("name")
	opt(meta)

	a.Equal(meta.Name, "name")
}

func TestVersion(t *testing.T) {
	a := assert.New(t)
	meta := &PluginMeta{}

	opt := Version("version")
	opt(meta)

	a.Equal(meta.Version, "version")
}

func TestLicense(t *testing.T) {
	a := assert.New(t)
	meta := &PluginMeta{}

	opt := License("license")
	opt(meta)

	a.Equal(meta.License, "license")
}

func TestDocs(t *testing.T) {
	a := assert.New(t)
	meta := &PluginMeta{}

	opt := Docs("docs")
	opt(meta)

	a.Equal(meta.Docs, "docs")
}

func TestAPIVersion(t *testing.T) {
	a := assert.New(t)
	meta := &PluginMeta{}

	opt := APIVersion(5)
	opt(meta)

	a.Equal(meta.APIVersion, 5)
}

func TestMeta(t *testing.T) {
	a := assert.New(t)
	meta := &PluginMeta{Meta: map[string]interface{}{}}

	opt := Meta("key", "value")
	opt(meta)

	a.Equal(meta.Meta, map[string]interface{}{
		"key": "value",
	})
}

func TestNewPluginMetaBase(t *testing.T) {
	a := assert.New(t)

	result := newPluginMeta("path", "fname", "name.space", "key", "plugin")

	a.Equal(result.Path, "path")
	a.Equal(result.Filename, "fname")
	a.Equal(result.Namespace, "name.space")
	a.Equal(result.Key, "key")
	a.Equal(result.Plugin, "plugin")
	a.Equal(result.Name, "")
	a.Equal(result.Version, "")
	a.Equal(result.License, "")
	a.Equal(result.Docs, "")
	a.Equal(result.APIVersion, 0)
	a.Equal(result.Meta, map[string]interface{}{})
}

func TestNewPluginMetaOptions(t *testing.T) {
	a := assert.New(t)

	result := newPluginMeta("path", "fname", "name.space", "key", "plugin",
		Name("name"),
		APIVersion(3),
	)

	a.Equal(result.Path, "path")
	a.Equal(result.Filename, "fname")
	a.Equal(result.Namespace, "name.space")
	a.Equal(result.Key, "key")
	a.Equal(result.Plugin, "plugin")
	a.Equal(result.Name, "name")
	a.Equal(result.Version, "")
	a.Equal(result.License, "")
	a.Equal(result.Docs, "")
	a.Equal(result.APIVersion, 3)
	a.Equal(result.Meta, map[string]interface{}{})
}
