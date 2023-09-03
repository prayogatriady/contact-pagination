package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateSort(t *testing.T) {
	assert.True(t, ValidateSort("Id asc"))
}

func TestValidateSort_FalseSorting(t *testing.T) {
	assert.False(t, ValidateSort("Id asc."))
}