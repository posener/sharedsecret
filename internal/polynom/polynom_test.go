package polynom

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInterpolate(t *testing.T) {
	t.Parallel()
	t.Run("simple", func(t *testing.T) {
		x := []*big.Int{big.NewInt(0), big.NewInt(2)}
		got := Interpolate(big.NewInt(1), x, x, big.NewInt(3))
		assert.Equal(t, big.NewInt(1), got)
	})

	t.Run("at given points", func(t *testing.T) {
		mod := big.NewInt(13)
		x := []*big.Int{big.NewInt(1), big.NewInt(12), big.NewInt(6)}
		y := []*big.Int{big.NewInt(4), big.NewInt(2), big.NewInt(7)}

		assert.Equal(t, big.NewInt(4), Interpolate(big.NewInt(1), x, y, mod))
		assert.Equal(t, big.NewInt(2), Interpolate(big.NewInt(12), x, y, mod))
		assert.Equal(t, big.NewInt(7), Interpolate(big.NewInt(6), x, y, mod))
	})
}

func TestInterpolate_panic(t *testing.T) {
	t.Parallel()
	t.Run("Different array sizes", func(t *testing.T) {
		result := Interpolate(big.NewInt(0), []*big.Int{big.NewInt(1), big.NewInt(2)}, []*big.Int{big.NewInt(1)}, big.NewInt(11))
		assert.Nil(t, result)
	})
	t.Run("x points are not unique", func(t *testing.T) {
		result := Interpolate(big.NewInt(0), []*big.Int{big.NewInt(1), big.NewInt(1)}, []*big.Int{big.NewInt(1), big.NewInt(2)}, big.NewInt(11))
		assert.Nil(t, result)
	})
}

func TestNewRandom(t *testing.T) {
	t.Parallel()
	p := NewRandom(1, big.NewInt(11))
	assert.Equal(t, 1, p.Deg())
	assert.Equal(t, p.ValueAt(big.NewInt(0)), p.Coeff(0))
}

func TestNewRandom_def2Fuzz(t *testing.T) {
	t.Parallel()
	for i := 0; i < 10000; i++ {
		x, err := rand.Int(rand.Reader, big.NewInt(10000))
		require.NoError(t, err)

		mod, err := rand.Int(rand.Reader, big.NewInt(10000))
		mod.Add(mod, big.NewInt(1)) // mod should be greater than 0
		require.NoError(t, err)

		p := NewRandom(2, mod)

		want := cp(p.Coeff(1))
		want.Mul(want, x)
		want.Add(want, p.Coeff(0))
		want.Mod(want, mod)

		assert.Equal(t, p.ValueAt(x), want)
	}
}

func TestNewRandom_panic(t *testing.T) {
	t.Parallel()
	assert.Panics(t, func() { NewRandom(0, big.NewInt(11)) })
	assert.Panics(t, func() { NewRandom(-1, big.NewInt(11)) })
}
