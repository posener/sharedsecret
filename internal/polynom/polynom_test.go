package polynom

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpolate(t *testing.T) {
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
