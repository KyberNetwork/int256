package int256

import (
	"math/big"
	"math/rand"
	"testing"
)

func BenchmarkAdd(bench *testing.B) {
	var (
		// 2^255 - 1
		lim, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		rnd    = rand.New(rand.NewSource(rand.Int63()))

		testcasesBI   = [][2]*big.Int{}
		testcasesI256 = [][2]*Int{}
	)

	for i := 0; i < 200; i++ {
		xBI := new(big.Int).Rand(rnd, lim)
		yBI := new(big.Int).Rand(rnd, new(big.Int).Sub(lim, xBI))
		negxBI := new(big.Int).Neg(xBI)
		negyBI := new(big.Int).Neg(yBI)

		xI256 := MustFromBig(xBI)
		yI256 := MustFromBig(yBI)
		negxI256 := new(Int).Neg(xI256)
		negyI256 := new(Int).Neg(yI256)

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, yI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, negyI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, negyI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, yI256})
	}

	sz := len(testcasesBI)

	addint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).Add(testcasesI256[testID][0], testcasesI256[testID][1])
		}
	}

	addbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).Add(testcasesBI[testID][0], testcasesBI[testID][1])
		}
	}

	bench.Run("big", addbig)
	bench.Run("int256", addint256)
}

func BenchmarkSub(bench *testing.B) {
	var (
		// 2^255 - 1
		lim, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		rnd    = rand.New(rand.NewSource(rand.Int63()))

		testcasesBI   = [][2]*big.Int{}
		testcasesI256 = [][2]*Int{}
	)

	for i := 0; i < 100; i++ {
		xBI := new(big.Int).Rand(rnd, lim)
		yBI := new(big.Int).Rand(rnd, xBI)
		negxBI := new(big.Int).Neg(xBI)
		negyBI := new(big.Int).Neg(yBI)

		xI256 := MustFromBig(xBI)
		yI256 := MustFromBig(yBI)
		negxI256 := new(Int).Neg(xI256)
		negyI256 := new(Int).Neg(yI256)

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, yI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, negyI256})
	}

	for i := 0; i < 100; i++ {
		xBI := new(big.Int).Rand(rnd, lim)
		yBI := new(big.Int).Rand(rnd, new(big.Int).Sub(lim, xBI))
		negxBI := new(big.Int).Neg(xBI)
		negyBI := new(big.Int).Neg(yBI)

		xI256 := MustFromBig(xBI)
		yI256 := MustFromBig(yBI)
		negxI256 := new(Int).Neg(xI256)
		negyI256 := new(Int).Neg(yI256)

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, negyI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, yI256})
	}

	sz := len(testcasesBI)

	subint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).Sub(testcasesI256[testID][0], testcasesI256[testID][1])
		}
	}

	subbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).Sub(testcasesBI[testID][0], testcasesBI[testID][1])
		}
	}

	bench.Run("big", subbig)
	bench.Run("int256", subint256)
}

func BenchmarkMul(bench *testing.B) {
	var (
		// 2^255 - 1
		lim, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		rnd    = rand.New(rand.NewSource(rand.Int63()))

		testcasesBI   = [][2]*big.Int{}
		testcasesI256 = [][2]*Int{}
	)

	for i := 0; i < 200; i++ {
		xBI := new(big.Int).Rand(rnd, lim)
		yBI := new(big.Int).Rand(rnd, new(big.Int).Quo(lim, xBI))
		negxBI := new(big.Int).Neg(xBI)
		negyBI := new(big.Int).Neg(yBI)

		xI256 := MustFromBig(xBI)
		yI256 := MustFromBig(yBI)
		negxI256 := new(Int).Neg(xI256)
		negyI256 := new(Int).Neg(yI256)

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, yI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, negyI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, negyI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, yI256})
	}

	sz := len(testcasesBI)

	mulint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).Mul(testcasesI256[testID][0], testcasesI256[testID][1])
		}
	}

	mulbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).Mul(testcasesBI[testID][0], testcasesBI[testID][1])
		}
	}

	bench.Run("big", mulbig)
	bench.Run("int256", mulint256)
}

func BenchmarkQuo(bench *testing.B) {
	var (
		// 2^255 - 1
		lim, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		rnd    = rand.New(rand.NewSource(rand.Int63()))

		testcasesBI   = [][2]*big.Int{}
		testcasesI256 = [][2]*Int{}
	)

	for i := 0; i < 200; i++ {
		xBI := new(big.Int).Rand(rnd, lim)
		yBI := new(big.Int).Rand(rnd, lim)
		negxBI := new(big.Int).Neg(xBI)
		negyBI := new(big.Int).Neg(yBI)

		xI256 := MustFromBig(xBI)
		yI256 := MustFromBig(yBI)
		negxI256 := new(Int).Neg(xI256)
		negyI256 := new(Int).Neg(yI256)

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, yI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, negyI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, negyI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, yI256})
	}

	sz := len(testcasesBI)

	divint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).Quo(testcasesI256[testID][0], testcasesI256[testID][1])
		}
	}

	divbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).Quo(testcasesBI[testID][0], testcasesBI[testID][1])
		}
	}

	bench.Run("big", divbig)
	bench.Run("int256", divint256)
}

func BenchmarkRem(bench *testing.B) {
	var (
		// 2^255 - 1
		lim, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		rnd    = rand.New(rand.NewSource(rand.Int63()))

		testcasesBI   = [][2]*big.Int{}
		testcasesI256 = [][2]*Int{}
	)

	for i := 0; i < 200; i++ {
		xBI := new(big.Int).Rand(rnd, lim)
		yBI := new(big.Int).Rand(rnd, lim)
		negxBI := new(big.Int).Neg(xBI)
		negyBI := new(big.Int).Neg(yBI)

		xI256 := MustFromBig(xBI)
		yI256 := MustFromBig(yBI)
		negxI256 := new(Int).Neg(xI256)
		negyI256 := new(Int).Neg(yI256)

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, yI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, negyI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, negyI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, yI256})
	}

	sz := len(testcasesBI)

	remint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).Rem(testcasesI256[testID][0], testcasesI256[testID][1])
		}
	}

	rembig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).Rem(testcasesBI[testID][0], testcasesBI[testID][1])
		}
	}

	bench.Run("big", rembig)
	bench.Run("int256", remint256)
}

func BenchmarkCmp(bench *testing.B) {
	var (
		// 2^255 - 1
		lim, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		rnd    = rand.New(rand.NewSource(rand.Int63()))

		testcasesBI   = [][2]*big.Int{}
		testcasesI256 = [][2]*Int{}
	)

	for i := 0; i < 200; i++ {
		xBI := new(big.Int).Rand(rnd, lim)
		yBI := new(big.Int).Rand(rnd, lim)
		negxBI := new(big.Int).Neg(xBI)
		negyBI := new(big.Int).Neg(yBI)

		xI256 := MustFromBig(xBI)
		yI256 := MustFromBig(yBI)
		negxI256 := new(Int).Neg(xI256)
		negyI256 := new(Int).Neg(yI256)

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, yI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, negyI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{xBI, negyBI})
		testcasesI256 = append(testcasesI256, [2]*Int{xI256, negyI256})

		testcasesBI = append(testcasesBI, [2]*big.Int{negxBI, yBI})
		testcasesI256 = append(testcasesI256, [2]*Int{negxI256, yI256})
	}

	sz := len(testcasesBI)

	cmpint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			testcasesI256[testID][0].Cmp(testcasesI256[testID][1])
		}
	}

	cmpbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			testcasesBI[testID][0].Cmp(testcasesBI[testID][1])
		}
	}

	bench.Run("big", cmpbig)
	bench.Run("int256", cmpint256)
}

func BenchmarkFromDecimalString(bench *testing.B) {
	var (
		// 2^255 - 1
		lim, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		rnd    = rand.New(rand.NewSource(rand.Int63()))

		testcases = []string{}
	)

	for i := 0; i < 200; i++ {
		x := new(big.Int).Rand(rnd, lim)
		negx := new(big.Int).Neg(x)
		testcases = append(testcases, x.String())
		testcases = append(testcases, negx.String())
	}

	sz := len(testcases)

	fromdecint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).SetFromDec(testcases[testID])
		}
	}

	fromdecbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).SetString(testcases[testID], 10)
		}
	}

	bench.Run("big", fromdecbig)
	bench.Run("int256", fromdecint256)
}
