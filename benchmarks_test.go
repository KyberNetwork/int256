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
			_ = new(Int).SetFromDec(testcases[testID])
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

func BenchmarkAnd(bench *testing.B) {
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

	andint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).And(testcasesI256[testID][0], testcasesI256[testID][1])
		}
	}

	andbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).And(testcasesBI[testID][0], testcasesBI[testID][1])
		}
	}

	bench.Run("big", andbig)
	bench.Run("int256", andint256)
}

func BenchmarkOr(bench *testing.B) {
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

	orint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).Or(testcasesI256[testID][0], testcasesI256[testID][1])
		}
	}

	orbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).Or(testcasesBI[testID][0], testcasesBI[testID][1])
		}
	}

	bench.Run("big", orbig)
	bench.Run("int256", orint256)
}

func BenchmarkXor(bench *testing.B) {
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

	xorint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).Xor(testcasesI256[testID][0], testcasesI256[testID][1])
		}
	}

	xorbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).Xor(testcasesBI[testID][0], testcasesBI[testID][1])
		}
	}

	bench.Run("big", xorbig)
	bench.Run("int256", xorint256)
}

func BenchmarkNot(bench *testing.B) {
	var (
		// 2^255 - 1
		lim, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		rnd    = rand.New(rand.NewSource(rand.Int63()))

		testcasesI256 = []*Int{}
		testcasesBI   = []*big.Int{}
	)

	for i := 0; i < 200; i++ {
		x := new(big.Int).Rand(rnd, lim)
		negx := new(big.Int).Neg(x)

		testcasesBI = append(testcasesBI, x)
		testcasesBI = append(testcasesBI, negx)

		testcasesI256 = append(testcasesI256, MustFromBig(x))
		testcasesI256 = append(testcasesI256, MustFromBig(negx))
	}

	sz := len(testcasesBI)

	notint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).Not(testcasesI256[testID])
		}
	}

	notbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).Not(testcasesBI[testID])
		}
	}

	bench.Run("big", notbig)
	bench.Run("int256", notint256)
}

func BenchmarkLsh(bench *testing.B) {
	type pairI256 struct {
		X *Int
		N uint
	}

	type pairBI struct {
		X *big.Int
		N uint
	}

	var (
		// 2^255 - 1
		lim, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		rnd    = rand.New(rand.NewSource(rand.Int63()))

		testcasesI256 = []pairI256{}
		testcasesBI   = []pairBI{}
	)

	for i := 0; i < 200; i++ {
		x := new(big.Int).Rand(rnd, lim)
		nx := uint(rnd.Int63() % 256)

		negx := new(big.Int).Neg(x)
		nnegx := uint(rnd.Int63() % 256)

		testcasesBI = append(testcasesBI, pairBI{
			X: x,
			N: nx,
		})
		testcasesBI = append(testcasesBI, pairBI{
			X: negx,
			N: nnegx,
		})

		testcasesI256 = append(testcasesI256, pairI256{
			X: MustFromBig(x),
			N: nx,
		})
		testcasesI256 = append(testcasesI256, pairI256{
			X: MustFromBig(negx),
			N: nnegx,
		})
	}

	sz := len(testcasesBI)

	lshint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).Lsh(testcasesI256[testID].X, testcasesI256[testID].N)
		}
	}

	lshbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).Lsh(testcasesBI[testID].X, testcasesBI[testID].N)
		}
	}

	bench.Run("big", lshbig)
	bench.Run("int256", lshint256)
}

func BenchmarkRsh(bench *testing.B) {
	type pairI256 struct {
		X *Int
		N uint
	}

	type pairBI struct {
		X *big.Int
		N uint
	}

	var (
		// 2^255 - 1
		lim, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		rnd    = rand.New(rand.NewSource(rand.Int63()))

		testcasesI256 = []pairI256{}
		testcasesBI   = []pairBI{}
	)

	for i := 0; i < 200; i++ {
		x := new(big.Int).Rand(rnd, lim)
		nx := uint(rnd.Int63() % 256)

		negx := new(big.Int).Neg(x)
		nnegx := uint(rnd.Int63() % 256)

		testcasesBI = append(testcasesBI, pairBI{
			X: x,
			N: nx,
		})
		testcasesBI = append(testcasesBI, pairBI{
			X: negx,
			N: nnegx,
		})

		testcasesI256 = append(testcasesI256, pairI256{
			X: MustFromBig(x),
			N: nx,
		})
		testcasesI256 = append(testcasesI256, pairI256{
			X: MustFromBig(negx),
			N: nnegx,
		})
	}

	sz := len(testcasesBI)

	rshint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).Rsh(testcasesI256[testID].X, testcasesI256[testID].N)
		}
	}

	rshbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).Rsh(testcasesBI[testID].X, testcasesBI[testID].N)
		}
	}

	bench.Run("big", rshbig)
	bench.Run("int256", rshint256)
}

func BenchmarkSqrt(bench *testing.B) {
	var (
		// 2^255 - 1
		lim, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		rnd    = rand.New(rand.NewSource(rand.Int63()))

		testcasesI256 = []*Int{}
		testcasesBI   = []*big.Int{}
	)

	for i := 0; i < 500; i++ {
		x := new(big.Int).Rand(rnd, lim)
		testcasesBI = append(testcasesBI, x)
		testcasesI256 = append(testcasesI256, MustFromBig(x))
	}

	sz := len(testcasesBI)

	sqrtint256 := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(Int).Sqrt(testcasesI256[testID])
		}
	}

	sqrtbig := func(bench *testing.B) {
		testID := 0
		bench.ResetTimer()
		for i := 0; i < bench.N; i++ {
			testID = i % sz
			new(big.Int).Sqrt(testcasesBI[testID])
		}
	}

	bench.Run("big", sqrtbig)
	bench.Run("int256", sqrtint256)
}
