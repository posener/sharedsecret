# sharedsecret

[![Build Status](https://travis-ci.org/posener/sharedsecret.svg?branch=master)](https://travis-ci.org/posener/sharedsecret)
[![codecov](https://codecov.io/gh/posener/sharedsecret/branch/master/graph/badge.svg)](https://codecov.io/gh/posener/sharedsecret)
[![GoDoc](https://godoc.org/github.com/posener/sharedsecret?status.svg)](http://godoc.org/github.com/posener/sharedsecret)
[![goreadme](https://goreadme.herokuapp.com/badge/posener/sharedsecret.svg)](https://goreadme.herokuapp.com)

Package sharedsecret is implementation of Shamir's Secret Sharing algorithm.

Shamir's Secret Sharing is an algorithm in cryptography created by Adi Shamir. It is a form of
secret sharing, where a secret is divided into parts, giving each participant its own unique
part. To reconstruct the original secret, a minimum number of parts is required. In the threshold
scheme this number is less than the total number of parts. Otherwise all participants are needed
to reconstruct the original secret.
See [wiki page](https://en.wikipedia.org/wiki/Shamir's_Secret_Sharing).

#### Examples

##### Distribute

With the `Distribute` function, a given secret can be distributed to shares.

```golang
secret := big.NewInt(120398491412912873)

// Create 5 shares that 3 or more of them can recover the secret.
shares := Distribute(secret, 5, 3)

// We can recover from only 3 (or more) shares:
recovered := Recover(shares[1], shares[3], shares[0])

fmt.Println(recovered)
```

 Output:

```
120398491412912873

```

##### New

With the `New` function, a random secret is generated and distributed into shares. Both the
secret and the shares are returned.

```golang
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
```

 Output:

```
true true

```


---

Created by [goreadme](https://github.com/apps/goreadme)
