package codegen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtensionsConfig_Validate(t *testing.T) {
	cfg := DefaultExtensionsConfig()
	require.Error(t, cfg.Validate())
}
