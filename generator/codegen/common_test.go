/*
 * Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */
package codegen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtensionsConfig_Validate(t *testing.T) {
	cfg := DefaultExtensionsConfig()
	require.Error(t, cfg.Validate())
}
