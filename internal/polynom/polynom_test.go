package polynom

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpolate(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		xs := []*big.Int{big.NewInt(0), big.NewInt(2)}
		got := Interpolate(big.NewInt(1), xs, xs)
		assert.Equal(t, big.NewInt(1), got)
	})

	t.Run("at given points", func(t *testing.T) {
		xs := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
		ys := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}

		assert.Equal(t, big.NewInt(1), Interpolate(big.NewInt(1), xs, ys))
		assert.Equal(t, big.NewInt(2), Interpolate(big.NewInt(2), xs, ys))
		assert.Equal(t, big.NewInt(3), Interpolate(big.NewInt(3), xs, ys))
	})
}
