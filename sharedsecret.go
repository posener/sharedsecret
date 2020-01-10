// Package sharedsecret is implementation of Shamir's Secret Sharing algorithm.
//
// Shamir's Secret Sharing is an algorithm in cryptography created by Adi Shamir. It is a form of
// secret sharing, where a secret is divided into parts, giving each participant its own unique
// part. To reconstruct the original secret, a minimum number of parts is required. In the threshold
// scheme this number is less than the total number of parts. Otherwise all participants are needed
// to reconstruct the original secret. See https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing.
package sharedsecret

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/posener/sharedsecret/internal/polynom"
)

// Share is a part of a secret.
type Share struct {
	x, y *big.Int
}

// Secret is secret.
type Secret *big.Int

// New creates n Shares and a secret. k defines the minimum number of shares that should be
// collected in order to recover the secret. Recovering the secret can be done by calling Recover
// with more than k Share objects.
func New(n, k int64) ([]Share, Secret) {
	if n < k {
		panic("Irrecoverable: not enough shares to reconstruct the secret.")
	}
	p := polynom.NewRandom(k)

	// Create the shares which are the value of p at any point but x != 0. Choose x in [1..n].
	shares := make([]Share, 0, n)
	for i := int64(1); i <= n; i++ {
		x := big.NewInt(i)
		y := p.ValueAt(x)
		shares = append(shares, Share{x: x, y: y})
	}

	// Secret is the value for x=0 which is the first coefficient (of x^0).
	secret := Secret(p.Coeff(0))

	return shares, secret
}

// Recover the secret from shares. Notice that the number of shares that is used should be at least
// the recover amount (k) that was used in order to create them in the New function.
func Recover(shares []Share) Secret {
	// Convert the shares to a list of points x[i], y[i].
	xs := make([]*big.Int, len(shares))
	ys := make([]*big.Int, len(shares))
	for i := range shares {
		xs[i] = shares[i].x
		ys[i] = shares[i].y
	}
	// Evaluate the polynom that goes through all (x[i], y[i]) points at x=0.
	return Secret(polynom.Interpolate(big.NewInt(0), xs, ys))
}

// String dumps the share object to a string.
func (s Share) String() string {
	return s.x.String() + "," + s.y.String()
}

// MarshalText implements the encoding.TextMarshaler interface.
func (s Share) MarshalText() ([]byte, error) {
	x, err := s.x.MarshalText()
	if err != nil {
		return nil, err
	}
	y, err := s.y.MarshalText()
	if err != nil {
		return nil, err
	}
	return append(append(x, ','), y...), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (s *Share) UnmarshalText(txt []byte) error {
	parts := bytes.Split(txt, []byte{','})
	if len(parts) != 2 {
		return errors.New("expected two parts")
	}
	s.x = &big.Int{}
	s.y = &big.Int{}
	err := s.x.UnmarshalText(parts[0])
	if err != nil {
		return err
	}
	return s.y.UnmarshalText(parts[1])
}
