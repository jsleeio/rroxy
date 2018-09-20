# rroxy

## what is this?

A trivial round-robin HTTP reverse proxy. I made this because I was seeing some
strange behaviour in Traefik under high concurrency and wanted to see if I
could reproduce it with just the underlying `oxy` package.

Almost entirely a copy/paste from the [oxy](https://github.com/vulcand/oxy)
README.

## how?

```
$ go install github.com/jsleeio/rroxy
$ go install github.com/jsleeio/hi-m8
$ for b in $(seq 1 6) ; do hi-m8 -listen=:300$b & done
[1] 43394
[2] 43395
[3] 43396
[4] 43397
[5] 43398
[6] 43399
$ rroxy http://localhost:300{1..6} &
[7] 43594
$ curl localhost:8000
hi m8
```
