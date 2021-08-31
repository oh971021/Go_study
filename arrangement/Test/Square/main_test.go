package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	assert.Equal(t, 81, square(9), "square(9) should be 81")
}

func TestSquare2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("Square(3) should be 9 but returns %d", rst)
	}
}

func TestSquare3(t *testing.T) {
	rst := square(5)
	if rst != 25 {
		t.Errorf("Square(5) should be 25 but returns %d", rst)
	}
}

func TestAdd(t *testing.T) {
	rst := add(3, 5)
	if rst != 8 {
		t.Errorf("Add(3, 5) should be 8 but returns %d", rst)
	}
}
