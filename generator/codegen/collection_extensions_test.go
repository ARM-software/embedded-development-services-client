package codegen

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ARM-software/golang-utils/utils/filesystem"
)

func TestCollectionExtensions(t *testing.T) {
	t.Run("Test Successful Generation", func(t *testing.T) {
		tmpDir := t.TempDir()
		d, err := GenerateDataStruct(ExtensionsConfig{
			Input:    "testdata/all-def-flat.json",
			Template: filepath.Join("templates", "collections.go.tmpl"),
			Output:   filepath.Join(tmpDir, "test.go"),
		})
		require.NoError(t, err)

		swagger, err := LoadAPIDefinition(d.SpecPath)
		require.NoError(t, err)

		d.Params, err = GetCollections(swagger)
		require.NoError(t, err)

		err = GenerateTemplateFile(d)
		assert.NoError(t, err)
		assert.FileExists(t, d.DestinationPath)

		expected, err := filesystem.ReadFile(filepath.Join("testdata", "collection_extension.gen.go"))
		require.NoError(t, err)
		actual, err := filesystem.ReadFile(d.DestinationPath)
		require.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	})
}
