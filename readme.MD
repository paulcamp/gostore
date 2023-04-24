# GoStore

A simple demonstration that serves a key/value store on a specified port  

## Notes

* thread safe store using [actor model](https://en.wikipedia.org/wiki/Actor_model) 
* store is volatile (stopping the service empties the store)


## Usage

Start the gostore (specifying port):

`go run ./... -port 8080`

## Endpoints


Test Service is running (HTTP GET):   
`curl localhost:8080/test`
  

POST Value:  
`curl -d "key=k&value=v" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/post`
  

GET Value:  
`curl localhost:8080/key`
  

DELETE Key:  
`curl -X "DELETE" 'http://localhost:8080/key'`
  

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


