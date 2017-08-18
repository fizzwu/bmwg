package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnOneKittenWhenSearchGarfield(t *testing.T) {
	store := MemoryStore{}
	kittens := store.Search("Garfield")

	assert.Equal(t, 1, len(kittens))
}

func TestReturnZeroKittenWhenSearchTom(t *testing.T) {
	store := MemoryStore{}
	kittens := store.Search("Tom")

	assert.Equal(t, 0, len(kittens))
}
