package sharedsecret

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// With the `New` function, a random secret is generated and distributed into shares. Both the
// secret and the shares are returned.
func Example_new() {
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

// With the `Distribute` function, a given secret can be distributed to shares.
func Example_distribute() {
	secret := big.NewInt(120398491412912873)

	// Create 5 shares that 3 or more of them can recover the secret.
	shares := Distribute(secret, 5, 3)

	// We can recover from only 3 (or more) shares:
	recovered := Recover(shares[1], shares[3], shares[0])

	fmt.Println(recovered)
	// Output: 120398491412912873
}

const (
	testN = 10
	testK = 4
)

func TestNewRecover_sanity(t *testing.T) {
	t.Parallel()

	// Create testN shares that testK or more of them can recover the secret.
	shares, secret := New(testN, testK)

	testSharesAndSecret(t, shares, secret)
}

func TestDistributeRecover_sanity(t *testing.T) {
	t.Parallel()

	// Create a secret and distribute it to testN shares that testK or more of them can recover the
	// secret.
	secret := big.NewInt(123456)
	shares := Distribute(secret, testN, testK)

	testSharesAndSecret(t, shares, secret)
}

func testSharesAndSecret(t *testing.T, shares []Share, secret *big.Int) {
	t.Helper()

	// All shares should recover.
	assert.Equal(t, secret, Recover(shares...))

	// The minimum number of shares should recover.
	assert.Equal(t, secret, Recover(shares[:testK]...))
	assert.Equal(t, secret, Recover(shares[testK:]...))

	// Less than the minimum number of shares should not recover.
	assert.NotEqual(t, secret, Recover(shares[:testK-1]...))
	assert.NotEqual(t, secret, Recover(shares[testN-testK+1:]...))
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
