# 256-bit signed integer math lib

## Benchmark

Not include overflow cases.

```
goos: darwin
goarch: arm64
pkg: github.com/KyberNetwork/int256
BenchmarkAdd/big-10             118787947                9.862 ns/op           0 B/op          0 allocs/op
BenchmarkAdd/int256-10          1000000000               0.6621 ns/op          0 B/op          0 allocs/op
BenchmarkSub/big-10             131103660                9.221 ns/op           0 B/op          0 allocs/op
BenchmarkSub/int256-10          1000000000               0.6607 ns/op          0 B/op          0 allocs/op
BenchmarkMul/big-10             203141756                5.987 ns/op           0 B/op          0 allocs/op
BenchmarkMul/int256-10          361456080                3.315 ns/op           0 B/op          0 allocs/op
BenchmarkDiv/big-10             17711032                65.35 ns/op           71 B/op          1 allocs/op
BenchmarkDiv/int256-10          39363513                30.22 ns/op            0 B/op          0 allocs/op
BenchmarkRem/big-10             21490580                54.21 ns/op            4 B/op          0 allocs/op
BenchmarkRem/int256-10          36572589                32.34 ns/op            0 B/op          0 allocs/op
BenchmarkCmp/big-10             276971919                4.307 ns/op           0 B/op          0 allocs/op
BenchmarkCmp/int256-10          364436605                3.268 ns/op           0 B/op          0 allocs/op
```
