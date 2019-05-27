# restgrpcserverbenchmark
Project comparing REST HTTP/1.1 and HTTP/2 and gRPC with Go Language

**Running Benchmark test without workers on HTTP/1.1 and HTTP/2**

Results with HTTP/1.1
```
2019/05/26 18:10:08 error executing request.  Get https://localhost:9191: dial tcp [::1]:9191: socket: too many open files
```

Results with HTTP/2
```
BenchmarkHTTP2Get-8        20000             90329 ns/op
```

**Now Running the same Benchmark test with workers**
```
BenchmarkHTTP11GetWithWorkers-8            20000             94896 ns/op
BenchmarkHTTP2GetWithWokers-8              10000            136328 ns/op
```

**Generating gRPC Client service interface**
```
protoc -I pb/ pb/random.proto --go_out=plugins=grpc:pb
```