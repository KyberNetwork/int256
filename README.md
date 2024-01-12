# 256-bit signed integer math lib

## Benchmark

Not include overflow cases.

```
goos: darwin
goarch: arm64
pkg: github.com/KyberNetwork/int256
BenchmarkAdd/big-10                         45337424                26.37 ns/op           72 B/op          1 allocs/op
BenchmarkAdd/int256-10                      1000000000               0.6601 ns/op          0 B/op          0 allocs/op
BenchmarkSub/big-10                         44229032                25.67 ns/op           72 B/op          1 allocs/op
BenchmarkSub/int256-10                      1000000000               0.6673 ns/op          0 B/op          0 allocs/op
BenchmarkMul/big-10                         94734348                12.40 ns/op           28 B/op          0 allocs/op
BenchmarkMul/int256-10                      346241904                3.438 ns/op           0 B/op          0 allocs/op
BenchmarkQuo/big-10                         15639882                74.99 ns/op           77 B/op          1 allocs/op
BenchmarkQuo/int256-10                      38883817                29.94 ns/op            0 B/op          0 allocs/op
BenchmarkRem/big-10                         16274137                70.74 ns/op           76 B/op          1 allocs/op
BenchmarkRem/int256-10                      35832081                33.01 ns/op            0 B/op          0 allocs/op
BenchmarkCmp/big-10                         283483120                4.202 ns/op           0 B/op          0 allocs/op
BenchmarkCmp/int256-10                      365665669                3.262 ns/op           0 B/op          0 allocs/op
BenchmarkFromDecimalString/big-10                2527873               473.7 ns/op            88 B/op          3 allocs/op
BenchmarkFromDecimalString/int256-10             9883429               119.6 ns/op             0 B/op          0 allocs/op
```
