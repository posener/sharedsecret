package sharedsecret

import (
	"math/big"

	"github.com/posener/sharedsecret/internal/polynom"
)

type Share struct {
	x, y *big.Int
}

type Secret *big.Int

func New(recoverCount int64, sharesCount int64) ([]Share, Secret) {
	if sharesCount < recoverCount {
		panic("Irrecoverable: total number of shares must be greater than the min number of shares.")
	}
	p := polynom.NewRandom(recoverCount)

	// Create the shares which are the value of p at any point but x != 0. Choose x in [1..n].
	shares := make([]Share, 0, sharesCount)
	for x := int64(1); x <= sharesCount; x++ {
		bigX := big.NewInt(x)
		shares = append(shares, Share{x: bigX, y: p.ValueAt(bigX)})
	}

	// Secret is the value for x=0 which is the first coefficient (of x^0).
	secret := Secret(p.Coeff(0))

	return shares, secret
}

func Recover(shares []Share) Secret {
	xs := make([]*big.Int, len(shares))
	ys := make([]*big.Int, len(shares))
	for i := range shares {
		xs[i] = shares[i].x
		ys[i] = shares[i].y
	}
	return Secret(polynom.Interpolate(big.NewInt(0), xs, ys))
}
