package sharedsecret

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecover_sanity(t *testing.T) {
	// Create 10 shares that 4 or more of them can recover the secret.
	shares, secret := New(10, 4)

	// All shares should recover.
	assert.Equal(t, secret, Recover(shares))

	// The minimum number of shares should recover.
	assert.Equal(t, secret, Recover(shares[:4]))
	assert.Equal(t, secret, Recover(shares[4:]))

	// Less than the minimum number of shares should not recover.
	assert.NotEqual(t, secret, Recover(shares[:3]))
	assert.NotEqual(t, secret, Recover(shares[7:]))
}
