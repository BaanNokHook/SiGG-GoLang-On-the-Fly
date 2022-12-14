// SiGG-GoLang-On-the-Fly //
//go:build reference
// +build reference

package reference

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestGenerateMarkdownPages(t *testing.T) {
	// TODO: Generate multiple languages when supported in the future here
	ctx := i18n.WithLang(context.Background(), language.AmericanEnglish)
	markdownMap, err := GenerateObjectsReferenceMarkdown(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, markdownMap)

	for pageName, markdown := range markdownMap {
		f, err := os.Create(filepath.Join("..", "..", "docs", "reference", "types", fmt.Sprintf("%s.md", pageName)))
		assert.NoError(t, err)
		_, err = f.Write(markdown)
		assert.NoError(t, err)
		err = f.Close()
		assert.NoError(t, err)
	}
}
