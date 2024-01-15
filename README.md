# 256-bit signed integer math lib

## Benchmark

Not include overflow cases.
```
goos: darwin
goarch: arm64
pkg: github.com/KyberNetwork/int256
BenchmarkAdd/big-10                     45337424                   26.37 ns/op               72 B/op              1 allocs/op
BenchmarkAdd/int256-10                  1000000000                0.6601 ns/op                0 B/op              0 allocs/op
BenchmarkSub/big-10                     44229032                   25.67 ns/op               72 B/op              1 allocs/op
BenchmarkSub/int256-10                  1000000000                0.6673 ns/op                0 B/op              0 allocs/op
BenchmarkMul/big-10                     94734348                   12.40 ns/op               28 B/op              0 allocs/op
BenchmarkMul/int256-10                  346241904                  3.438 ns/op                0 B/op              0 allocs/op
BenchmarkQuo/big-10                     15639882                   74.99 ns/op               77 B/op              1 allocs/op
BenchmarkQuo/int256-10                  38883817                   29.94 ns/op                0 B/op              0 allocs/op
BenchmarkRem/big-10                     16274137                   70.74 ns/op               76 B/op              1 allocs/op
BenchmarkRem/int256-10                  35832081                   33.01 ns/op                0 B/op              0 allocs/op
BenchmarkCmp/big-10                     283483120                  4.202 ns/op                0 B/op              0 allocs/op
BenchmarkCmp/int256-10                  365665669                  3.262 ns/op                0 B/op              0 allocs/op
BenchmarkFromDecimalString/big-10       2527873                    473.7 ns/op               88 B/op              3 allocs/op
BenchmarkFromDecimalString/int256-10    9883429                    119.6 ns/op                0 B/op              0 allocs/op
BenchmarkAnd/big-10                     22740283                   51.33 ns/op              148 B/op              2 allocs/op
BenchmarkAnd/int256-10                  1000000000                0.6833 ns/op                0 B/op              0 allocs/op
BenchmarkOr/big-10                      18030810                   65.77 ns/op              188 B/op              2 allocs/op
BenchmarkOr/int256-10                   1000000000                0.6828 ns/op                0 B/op              0 allocs/op
BenchmarkXor/big-10                     19921944                   60.22 ns/op              168 B/op              2 allocs/op
BenchmarkXor/int256-10                  1000000000                0.6836 ns/op                0 B/op              0 allocs/op
BenchmarkNot/big-10                     47593866                   25.20 ns/op               72 B/op              1 allocs/op
BenchmarkNot/int256-10                  1000000000                0.5794 ns/op                0 B/op              0 allocs/op
BenchmarkLsh/big-10                     37263223                   31.85 ns/op               88 B/op              1 allocs/op
BenchmarkLsh/int256-10                  449856327                  2.662 ns/op                0 B/op              0 allocs/op
BenchmarkRsh/big-10                     35134976                   34.03 ns/op               54 B/op              1 allocs/op
BenchmarkRsh/int256-10                  323528000                  3.717 ns/op                0 B/op              0 allocs/op
BenchmarkSqrt/big-10                    1000000                     1163 ns/op              722 B/op             10 allocs/op
BenchmarkSqrt/int256-10                 2973258                    405.5 ns/op                0 B/op              0 allocs/op
```
