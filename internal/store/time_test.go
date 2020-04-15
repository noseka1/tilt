package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)
func TestBefore(t *testing.T) {
  a := time.Now()
  b := time.Now()
  assert.False(t, a.Before(b))
  assert.True(t, a.Equal(b))
}
