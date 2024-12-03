package webhook

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func loadData(t *testing.T, filepath string) []byte {
	content, err := os.ReadFile(filepath)
	assert.NoError(t, err)
	return content
}
