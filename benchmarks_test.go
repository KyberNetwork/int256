package int256

import (
	"math/big"
	"testing"
)

func BenchmarkAdd(bench *testing.B) {
	addint256 := func(bench *testing.B) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		z := new(Int)
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			z.Add(x, y)
		}
	}

	addbig := func(bench *testing.B) {
		x, _ := new(big.Int).SetString("-57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		y, _ := new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		z := new(big.Int)
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			z.Add(x, y)
		}
	}

	bench.Run("big", addbig)
	bench.Run("int256", addint256)

	/*
		goos: darwin
		goarch: arm64
		pkg: github.com/KyberNetwork/int256
		BenchmarkAdd/big-10         	102299776	        11.62 ns/op	       0 B/op	       0 allocs/op
		BenchmarkAdd/int256-10      	1000000000	         0.3110 ns/op	       0 B/op	       0 allocs/op
	*/
}

func BenchmarkSub(bench *testing.B) {
	subint256 := func(bench *testing.B) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		z := new(Int)
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			z.Sub(x, y)
		}
	}

	subbig := func(bench *testing.B) {
		x, _ := new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		y, _ := new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		z := new(big.Int)
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			z.Sub(x, y)
		}
	}

	bench.Run("big", subbig)
	bench.Run("int256", subint256)

	/*
	   goos: darwin
	   goarch: arm64
	   pkg: github.com/KyberNetwork/int256
	   BenchmarkSub/big-10         	99478418	        11.68 ns/op	       0 B/op	       0 allocs/op
	   BenchmarkSub/int256-10      	1000000000	         0.3080 ns/op	       0 B/op	       0 allocs/op
	*/
}

func BenchmarkMul(bench *testing.B) {
	mulint256 := func(bench *testing.B) {
		x := MustFromDec("24061596916800451154503377247762505692")
		y := MustFromDec("-24061596916800451154503377247762505692")
		z := new(Int)
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			z.Mul(x, y)
		}
	}

	mulbig := func(bench *testing.B) {
		x, _ := new(big.Int).SetString("24061596916800451154503377247762505692", 10)
		y, _ := new(big.Int).SetString("-24061596916800451154503377247762505692", 10)
		z := new(big.Int)
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			z.Mul(x, y)
		}
	}

	bench.Run("big", mulbig)
	bench.Run("int256", mulint256)

	/*
	   goos: darwin
	   goarch: arm64
	   pkg: github.com/KyberNetwork/int256
	   BenchmarkMul/big-10         	95598220	        12.28 ns/op	       0 B/op	       0 allocs/op
	   BenchmarkMul/int256-10      	416813202	         2.894 ns/op	       0 B/op	       0 allocs/op
	*/
}

func BenchmarkDiv(bench *testing.B) {
	divint256 := func(bench *testing.B) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("23")
		z := new(Int)
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			z.Mul(x, y)
		}
	}

	divbig := func(bench *testing.B) {
		x, _ := new(big.Int).SetString("-57896044618658097711785492504343953926634992332820282019728792003956564819968", 10)
		y, _ := new(big.Int).SetString("23", 10)
		z := new(big.Int)
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			z.Mul(x, y)
		}
	}

	bench.Run("big", divbig)
	bench.Run("int256", divint256)

	/*
	   goos: darwin
	   goarch: arm64
	   pkg: github.com/KyberNetwork/int256
	   BenchmarkDiv/big-10         	166427436	         7.089 ns/op	       0 B/op	       0 allocs/op
	   BenchmarkDiv/int256-10      	414451401	         2.899 ns/op	       0 B/op	       0 allocs/op
	*/
}
