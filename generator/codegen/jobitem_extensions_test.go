package codegen

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ARM-software/golang-utils/utils/filesystem"
)

func TestJobItemsExtensions(t *testing.T) {
	t.Run("Test Successful Generation", func(t *testing.T) {
		tmpDir := t.TempDir()
		d, err := GenerateDataStruct(ExtensionsConfig{
			Input:             filepath.Join("testdata", "jobs", "test.json"),
			Template:          filepath.Join("templates", "jobs.go.tmpl"),
			Output:            filepath.Join(tmpDir, "test.go"),
			ClientPackagePath: filepath.Join("..", "..", "client"),
		})
		require.NoError(t, err)

		swagger, err := loadAPIDefinition(d.SpecPath)
		require.NoError(t, err)

		d.Params, err = GetJobItems(swagger)
		require.NoError(t, err)

		err = GenerateTemplateFile(context.Background(), d)
		assert.NoError(t, err)
		assert.FileExists(t, d.DestinationPath)

		expected, err := filesystem.ReadFile(filepath.Join("testdata", "jobs", "test.gen.go"))
		require.NoError(t, err)
		actual, err := filesystem.ReadFile(d.DestinationPath)
		require.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	})

	t.Run("Test Successful Generation (with all redacts)", func(t *testing.T) {
		tmpDir := t.TempDir()
		d, err := GenerateDataStruct(ExtensionsConfig{
			Input:             filepath.Join("testdata", "jobs", "test-all-redacted.json"),
			Template:          filepath.Join("templates", "jobs.go.tmpl"),
			Output:            filepath.Join(tmpDir, "test.go"),
			ClientPackagePath: filepath.Join("..", "..", "client"),
		})
		require.NoError(t, err)

		swagger, err := loadAPIDefinition(d.SpecPath)
		require.NoError(t, err)

		d.Params, err = GetJobItems(swagger)
		require.NoError(t, err)

		err = GenerateTemplateFile(context.Background(), d)
		assert.NoError(t, err)
		assert.FileExists(t, d.DestinationPath)

		expected, err := filesystem.ReadFile(filepath.Join("testdata", "jobs", "test-all-redacted.gen.go"))
		require.NoError(t, err)
		actual, err := filesystem.ReadFile(d.DestinationPath)
		require.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	})

	t.Run("Test Successful Generation (with some redacts)", func(t *testing.T) {
		tmpDir := t.TempDir()
		d, err := GenerateDataStruct(ExtensionsConfig{
			Input:             filepath.Join("testdata", "jobs", "test-some-redacted.json"),
			Template:          filepath.Join("templates", "jobs.go.tmpl"),
			Output:            filepath.Join(tmpDir, "test.go"),
			ClientPackagePath: filepath.Join("..", "..", "client"),
		})
		require.NoError(t, err)

		swagger, err := loadAPIDefinition(d.SpecPath)
		require.NoError(t, err)

		d.Params, err = GetJobItems(swagger)
		require.NoError(t, err)

		err = GenerateTemplateFile(context.Background(), d)
		assert.NoError(t, err)
		assert.FileExists(t, d.DestinationPath)

		expected, err := filesystem.ReadFile(filepath.Join("testdata", "jobs", "test-some-redacted.gen.go"))
		require.NoError(t, err)
		actual, err := filesystem.ReadFile(d.DestinationPath)
		require.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	})
}
