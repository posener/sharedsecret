package sharedsecret

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Example() {
	// Create 5 shares that 3 or more of them can recover the secret.
	shares, secret := New(5, 3)

	// Now we should distribute the shares to different parties and forget about the shares and
	// secret. Once the original secret is needed, at least 3 shares should be used in order to
	// recover it:

	// We can't recover from only 2 shares:
	wrong := Recover(shares[1], shares[3])

	// We can recover from only 3 (or more) shares:
	correct := Recover(shares[1], shares[3], shares[0])

	fmt.Println(secret.Cmp(wrong) != 0, secret.Cmp(correct) == 0)
	// Output: true true
}

func TestRecover_sanity(t *testing.T) {
	t.Parallel()
	// Create 10 shares that 4 or more of them can recover the secret.
	shares, secret := New(10, 4)

	// All shares should recover.
	assert.Equal(t, secret, Recover(shares...))

	// The minimum number of shares should recover.
	assert.Equal(t, secret, Recover(shares[:4]...))
	assert.Equal(t, secret, Recover(shares[4:]...))

	// Less than the minimum number of shares should not recover.
	assert.NotEqual(t, secret, Recover(shares[:3]...))
	assert.NotEqual(t, secret, Recover(shares[7:]...))
}

func TestRecover_panic(t *testing.T) {
	t.Parallel()
	assert.Panics(t, func() { New(1, 2) })
	assert.Panics(t, func() { New(1, 0) })
}

func TestShareString(t *testing.T) {
	t.Parallel()
	s := Share{big.NewInt(0), big.NewInt(1)}
	assert.Equal(t, "0,1", s.String())
}

func TestShareMarshalText_fuzz(t *testing.T) {
	t.Parallel()
	for i := 0; i < 10000; i++ {
		x, err := rand.Int(rand.Reader, big.NewInt(10000))
		require.NoError(t, err)
		y, err := rand.Int(rand.Reader, big.NewInt(10000))
		require.NoError(t, err)
		want := Share{x: x, y: y}

		text, err := want.MarshalText()
		require.NoError(t, err)
		var got Share
		err = got.UnmarshalText(text)
		require.NoError(t, err)
		assert.True(t, want.x.Cmp(got.x) == 0)
		assert.True(t, want.y.Cmp(got.y) == 0)
	}
}

func TestShareUnMarshalText_errors(t *testing.T) {
	t.Parallel()
	var s Share
	assert.Error(t, s.UnmarshalText([]byte("")))
	assert.Error(t, s.UnmarshalText([]byte("1,2,3")))
	assert.Error(t, s.UnmarshalText([]byte("a,1")))
	assert.Error(t, s.UnmarshalText([]byte("1,a")))
}
