# GoStore

A simple demonstration that serves a key/value store on a specified port  

## Features

* thread safe store using [actor model](https://en.wikipedia.org/wiki/Actor_model) 


## Usage

### Start the gostore
`go run ./... -port 8080`

### commands

`curl localhost:8080/test`


## Benchmarks

```
goos: darwin
goarch: arm64
pkg: gostore/store
=== RUN   BenchmarkStore_Put
BenchmarkStore_Put
=== RUN   BenchmarkStore_Put/Benchmarking_PUT_on_the_store
BenchmarkStore_Put/Benchmarking_PUT_on_the_store
BenchmarkStore_Put/Benchmarking_PUT_on_the_store-10               327740              3522 ns/op             544 B/op          8 allocs/op
PASS
ok      gostore/store   1.442s


goos: darwin
goarch: arm64
pkg: gostore/store
=== RUN   BenchmarkStore_Get
BenchmarkStore_Get
=== RUN   BenchmarkStore_Get/Benchmarking_GET_on_the_store
BenchmarkStore_Get/Benchmarking_GET_on_the_store
BenchmarkStore_Get/Benchmarking_GET_on_the_store-10               358952              3389 ns/op             440 B/op          9 allocs/op
PASS
ok      gostore/store   2.328s
```


