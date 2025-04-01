package codegen

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/ARM-software/golang-utils/utils/filesystem"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinkFollowerExtensions(t *testing.T) {
	t.Run("Test Successful Generation", func(t *testing.T) {
		tmpDir := t.TempDir()
		d, err := GenerateDataStruct(ExtensionsConfig{
			Input:    filepath.Join("testdata", "linkfollowers", "test.json"),
			Template: filepath.Join("templates", "linkfollowers.go.tmpl"),
			Output:   filepath.Join(tmpDir, "test.go"),
		})
		require.NoError(t, err)

		swagger, err := loadAPIDefinition(d.SpecPath)
		require.NoError(t, err)

		d.Params, err = getLinkFollowersParams(swagger, d)
		require.NoError(t, err)

		err = GenerateTemplateFile(context.Background(), d)
		assert.NoError(t, err)
		assert.FileExists(t, d.DestinationPath)

		expected, err := filesystem.ReadFile(filepath.Join("testdata", "linkfollowers", "test.gen.go"))
		require.NoError(t, err)
		actual, err := filesystem.ReadFile(d.DestinationPath)
		require.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	})
}
