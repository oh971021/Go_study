package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	assert.Equal(t, 81, square(9), "should be 81")

	assert.Equal(t, 9, square(3), "should be 9")

	assert.Equal(t, 16, square(4), "should be 16")

	assert.Equal(t, 25, square(5), "should be 25")

	assert.Equal(t, 36, square(-6), "should be 36")
}
