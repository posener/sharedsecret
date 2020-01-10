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
to reconstruct the original secret. See [https://en.wikipedia.org/wiki/Shamir&#39;s_Secret_Sharing](https://en.wikipedia.org/wiki/Shamir's_Secret_Sharing).


---

Created by [goreadme](https://github.com/apps/goreadme)
