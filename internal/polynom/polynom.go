package polynom

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Polynom represents a big integer polynom.
type Polynom struct {
	// coeff are the coefficients of the polynom. coeff[i] is the coefficient of x^i.
	coeff []*big.Int
	// mod is the modulus for polynom arithmetics calculations.
	mod *big.Int
}

// NewRandom returns a new random polynom of the given degree, which is subjected to arithmetics of
// the given modulus.
func NewRandom(degree int64, modulus *big.Int) Polynom {
	if degree <= 0 {
		panic("deg must be positive number")
	}

	var (
		err   error
		coeff = make([]*big.Int, degree)
	)
	for i := range coeff {
		coeff[i], err = rand.Int(rand.Reader, modulus)
		if err != nil {
			panic(fmt.Sprintf("creating random int: %s", err))
		}
	}
	return Polynom{coeff: coeff, mod: modulus}
}

// Deg returns the degree of the polynom.
func (p Polynom) Deg() int {
	return len(p.coeff)
}

// Coeff returns the i'th coefficient.
//
// Can panic with index out of range when i >= p.Deg().
func (p Polynom) Coeff(i int) *big.Int {
	return cp(p.coeff[i])
}

// ValueAt returns the y value of the polynom on a given x0 value.
func (p Polynom) ValueAt(x0 *big.Int) *big.Int {
	val := big.NewInt(0)
	for i := len(p.coeff) - 1; i >= 0; i-- {
		val.Mul(val, x0)
		val.Add(val, p.coeff[i])
		val.Mod(val, p.mod)
	}
	return val
}

// Interpolate returns the y value at x0 of a polynom that lies on points (x[i], y[i]), with modulus
// arithmetics for the given modulus.
func Interpolate(x0 *big.Int, x []*big.Int, y []*big.Int, modulus *big.Int) (y0 *big.Int) {
	if len(x) != len(y) {
		panic("x and y lists must have the same length.")
	}
	assertDistinct(x)

	nums := make([]*big.Int, len(x))
	dens := make([]*big.Int, len(x))

	for i := range x {
		nums[i] = product(x, x0, i)
		dens[i] = product(x, x[i], i)
	}

	den := product(dens, nil, -1)

	num := big.NewInt(0)
	for i := range nums {
		nums[i].Mul(nums[i], den)
		nums[i].Mul(nums[i], y[i])
		nums[i].Mod(nums[i], modulus)
		num.Add(num, divmod(nums[i], dens[i], modulus))
	}

	y0 = divmod(num, den, modulus)
	y0.Add(y0, modulus)
	y0.Mod(y0, modulus)
	return y0
}

// product returns the product of vals. If sub is given, the returned product is of (sub-vals[i]).
// If skip is given, the i'th value will be ignored.
func product(vals []*big.Int, sub *big.Int, skip int) *big.Int {
	p := big.NewInt(1)
	for i := range vals {
		if i == skip {
			continue
		}
		v := cp(vals[i])
		if sub != nil {
			v.Sub(sub, v)
		}
		p.Mul(p, v)
	}
	return p
}

// divmod computes num / den modulo mod.
func divmod(a, b, mod *big.Int) *big.Int {
	return a.Mul(a, b.ModInverse(b, mod))
}

// cp copies a big.Int.
func cp(v *big.Int) *big.Int {
	var u big.Int
	u.Set(v)
	return &u
}

// assertDistinct panics if there are two identical values in vals.
func assertDistinct(vals []*big.Int) {
	s := make(map[*big.Int]bool, len(vals))
	for _, v := range vals {
		if s[v] {
			panic("points must be distinct")
		}
		s[v] = true
	}
}
