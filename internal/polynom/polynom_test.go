package polynom

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpolate(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		x := []*big.Int{big.NewInt(0), big.NewInt(2)}
		got := Interpolate(big.NewInt(1), x, x)
		assert.Equal(t, big.NewInt(1), got)
	})

	t.Run("at given points", func(t *testing.T) {
		x := []*big.Int{big.NewInt(1), big.NewInt(12), big.NewInt(6)}
		y := []*big.Int{big.NewInt(111), big.NewInt(15), big.NewInt(34)}

		assert.Equal(t, big.NewInt(111), Interpolate(big.NewInt(1), x, y))
		assert.Equal(t, big.NewInt(15), Interpolate(big.NewInt(12), x, y))
		assert.Equal(t, big.NewInt(34), Interpolate(big.NewInt(6), x, y))
	})
}
