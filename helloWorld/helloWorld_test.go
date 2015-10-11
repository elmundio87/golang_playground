package helloworld

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, ToString(), "Hello World", "Should say Hello World!")
}
