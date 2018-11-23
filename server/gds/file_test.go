package gds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFile(t *testing.T) {
	_, err := NewFile("../../LICENSE", "LICENSE", 5)
	assert.NotNil(t, err)

	file, err := NewFile("../../LICENSE", "LICENSE", 11357)
	assert.Nil(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, file.Size, int64(11357))
	assert.Equal(t, file.Type, FILE)

	dir, _ := NewFile("..", "server", 11)
	assert.Equal(t, dir.Size, int64(0))
	assert.Equal(t, dir.Type, DIR)
}

func TestNewFile_jpeg(t *testing.T) {
	jpg, err := NewFile("bona2.jpg", "bona2", 66679)
	assert.Nil(t, err)
	assert.Equal(t, jpg.Type, IMAGE)
	assert.Equal(t, jpg.Sub, JPEG)
}
