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
BenchmarkMain-4   	   10000	   1828116 ns/op
PASS
ok  	github.com/mekilis/purify/cmd/purify	19.059s

real	0m19.592s
user	0m3.547s
sys	0m16.935s
```

#### Setting the benchmark time to 10 seconds

```bash
time go test -bench=BenchmarkMain -benchtime=10s
```

```bash
goos: linux
goarch: amd64
pkg: github.com/mekilis/purify/cmd/purify
BenchmarkMain-4   	  100000	   2100485 ns/op
PASS
ok  	github.com/mekilis/purify/cmd/purify	212.389s

real	3m32.907s
user	0m22.942s
sys	3m14.752s
```
