// SiGG-GoLang-On-the-Fly //
//go:build reference
// +build reference

package config

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly/internal/apiserver"
	"github.com/hyperledger/firefly/internal/namespace"
	"github.com/stretchr/testify/assert"
)

func TestGenerateConfigDocs(t *testing.T) {
	// Initialize config of all plugins
	namespace.NewNamespaceManager(false)
	apiserver.InitConfig()
	f, err := os.Create(filepath.Join("reference", "config.md"))
	assert.NoError(t, err)
	generatedConfig, err := config.GenerateConfigMarkdown(context.Background(), configDocHeader, config.GetKnownKeys())
	assert.NoError(t, err)
	_, err = f.Write(generatedConfig)
	assert.NoError(t, err)
	err = f.Close()
	assert.NoError(t, err)
}

const configDocHeader = `---
layout: default
title: Configuration Reference
parent: pages.reference
nav_order: 2
---

# Configuration Reference
{: .no_toc }

<!-- ## Table of contents
{: .no_toc .text-delta }

1. TOC
{:toc} -->

---
`
