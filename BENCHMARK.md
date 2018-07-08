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
BenchmarkMain-4   	   20000	     64510 ns/op
PASS
ok  	github.com/mekilis/purify/cmd/purify	1.995s

real	0m2.444s
user	0m2.414s
sys	0m0.938s
```

#### Setting the benchmark time to 10 seconds

```bash
time go test -bench=BenchmarkMain -benchtime=10s
```

```bash
Main -benchtime 10s
goos: linux
goarch: amd64
pkg: github.com/mekilis/purify/cmd/purify
BenchmarkMain-4   	  200000	     63687 ns/op
PASS
ok  	github.com/mekilis/purify/cmd/purify	13.490s

real	0m14.010s
user	0m13.097s
sys	0m5.737s
```
