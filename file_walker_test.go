package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileWalker(t *testing.T) {
	files, err := WalkMatch("tests", WalkMatchOptions{patterns: []string{"*.css"}})
	assert.NoError(t, err)
	assert.Equal(t, len(files), 2)
}
