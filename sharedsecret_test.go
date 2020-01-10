package sharedsecret

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	shares, secret := New(5, 10)

	got := Recover(shares)

	assert.Equal(t, secret, got)
}
