package polynom

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	securityLevelBits = 128
)

var prime *big.Int

func init() {
	// Set prime to 2<<127 - 1.
	prime = big.NewInt(2)
	prime.Exp(prime, big.NewInt(127), nil)
	prime.Sub(prime, big.NewInt(1))
	println("prime:", prime.String())
}

type Polynom struct {
	coeff []*big.Int
}

// NewRandom returns a new random polynom of the given degree.
func NewRandom(deg int64) Polynom {
	if deg <= 0 {
		panic("deg must be positive number")
	}

	var (
		err   error
		coeff = make([]*big.Int, deg)
	)
	for i := range coeff {
		coeff[i], err = rand.Int(rand.Reader, prime)
		if err != nil {
			panic(fmt.Sprintf("creating random int: %s", err))
		}
	}
	return Polynom{coeff: coeff}
}

// Deg returns the degree of the polynom.
func (p Polynom) Deg() int {
	return len(p.coeff)
}

// Coeff returns the i'th coefficient.
//
// Can panic with index out of range when i >= p.Deg().
func (p Polynom) Coeff(i int) *big.Int {
	return p.coeff[i]
}

func (p Polynom) ValueAt(x *big.Int) *big.Int {
	val := big.NewInt(0)
	for i := len(p.coeff) - 1; i >= 0; i-- {
		val.Mul(val, x)
		val.Add(val, p.coeff[i])
		val.Mod(val, prime)
	}
	return val
}

func Interpolate(x0 *big.Int, x []*big.Int, y []*big.Int) (y0 *big.Int) {
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
	for i := range x {
		nums[i].Mul(nums[i], den)
		nums[i].Mul(nums[i], y[i])
		nums[i].Mod(nums[i], prime)
		num.Add(num, divmod(nums[i], dens[i]))
	}

	y0 = divmod(num, den)
	y0.Add(y0, prime)
	y0.Mod(y0, prime)
	return y0
}

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

func divmod(a, b *big.Int) *big.Int {
	p := cp(prime)
	a = cp(a)
	b = cp(b)
	return a.Mul(a, b.ModInverse(b, p))
}

func gcd(a, b *big.Int) (*big.Int, *big.Int) {

	x := big.NewInt(0)
	lastX := big.NewInt(1)
	y := big.NewInt(1)
	lastY := big.NewInt(0)
	for b.Cmp(big.NewInt(0)) != 0 {
		quot := cp(a)
		quot.Div(quot, b)

		oldB := cp(b)
		b.Mod(a, b)
		a = oldB

		x, lastX = xchange(x, lastX, quot)
		y, lastY = xchange(y, lastY, quot)
	}
	return lastX, lastY
}

func xchange(v, lastV, quot *big.Int) (*big.Int, *big.Int) {
	newV := cp(quot)
	newV.Mul(newV, v)
	newV.Sub(lastV, newV)
	return newV, lastV
}

func cp(x *big.Int) *big.Int {
	var xx big.Int
	xx.Set(x)
	return &xx
}

func assertDistinct(vals []*big.Int) {
	s := make(map[*big.Int]bool, len(vals))
	for _, v := range vals {
		if s[v] {
			panic("points must be distinct")
		}
		s[v] = true
	}
}
