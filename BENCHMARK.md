# Benchmarks
### System Specifications
* elementary OS 0.4.1 Loki,
* Dual-Core Intel® Core™ i7-6500U CPU @ 2.50GHz,
* 16.1 GB memory


### Results

#### Default
```bash
time go test -bench=BenchmarkMain
```

```bash
goos: linux
goarch: amd64
pkg: github.com/mekilis/purify/cmd/purify
BenchmarkMain-4   	    5000	    344304 ns/op
PASS
ok  	github.com/mekilis/purify/cmd/purify	2.495s

real	0m2.924s
user	0m2.180s
sys	0m0.843s
```

#### Setting the benchmark time to 10 seconds

```bash
time go test -bench=BenchmarkMain -benchtime=10s
```

```bash
goos: linux
goarch: amd64
pkg: github.com/mekilis/purify/cmd/purify
BenchmarkMain-4   	   50000	   2401518 ns/op
PASS
ok  	github.com/mekilis/purify/cmd/purify	123.325s

real	2m3.851s
user	0m14.111s
sys	1m50.683s
```
