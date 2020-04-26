package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	value := true
	assert.True(t, value)
}

func Test2(t *testing.T) {
	value := false
	assert.False(t, value)
}
