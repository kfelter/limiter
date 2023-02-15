# limiter

Benchmark results
```
limiter % go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/kfelter/limiter
BenchmarkRunSmall-10                4330            268090 ns/op
BenchmarkRunWaitMedium-10           4120            309699 ns/op
BenchmarkWaitRunLarge-10            3478            303175 ns/op
PASS
ok      github.com/kfelter/limiter      4.599s
```