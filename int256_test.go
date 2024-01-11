package int256

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSetInt64(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		v := int64(math.MaxInt64)
		z := new(Int).SetInt64(v)
		assert.Equal(t, v, z.Int64())
	})

	t.Run("negative", func(t *testing.T) {
		v := int64(math.MinInt64)
		z := new(Int).SetInt64(v)
		assert.Equal(t, v, z.Int64())
	})

	t.Run("zero", func(t *testing.T) {
		v := int64(0)
		z := new(Int).SetInt64(v)
		assert.Equal(t, v, z.Int64())
	})
}

func TestAdd(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564606844")
		y := MustFromDec("-431294739547329532759")
		expected := "57896044618658097711785492504343953926634992332820282019297497264409235500333"
		z := new(Int).Add(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "0"
		z := new(Int).Add(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "-1"
		z := new(Int).Add(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("10000")
		y := MustFromDec("-2000000")
		expected := "-1990000"
		z := new(Int).Add(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("5. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "-2"
		z := new(Int).Add(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("6. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819961")
		y := MustFromDec("431405283104328105143242031570231414")
		expected := "-57896044618658097711785492504343953926634560927537177691623648761924994588561"
		z := new(Int).Add(x, y)
		assert.Equal(t, expected, z.Dec())
	})
}

func TestNeg(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("100000")
		expected := "-100000"
		z := new(Int).Neg(x)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		z := new(Int).Neg(x)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "-57896044618658097711785492504343953926634992332820282019728792003956564819967"
		z := new(Int).Neg(x)
		assert.Equal(t, expected, z.Dec())
	})
}

func TestSub(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("100")
		expected := "57896044618658097711785492504343953926634992332820282019728792003956564819868"
		z := new(Int).Sub(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		y := MustFromDec("9999")
		expected := "-9999"
		z := new(Int).Sub(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "-1"
		z := new(Int).Sub(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "1"
		z := new(Int).Sub(x, y)
		assert.Equal(t, expected, z.Dec())
	})
}

func TestQuo(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("100")
		y := MustFromDec("-2")
		expected := "-50"

		z := new(Int).Quo(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("3194721952341")
		expected := "-18122404854742851078740194629148605029554483750497671811411863963"

		z := new(Int).Quo(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819961")
		y := MustFromDec("-4134320184329015710493829104")
		expected := "-14003764110508630918697735423914335136046266814161"

		z := new(Int).Quo(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819961")
		y := MustFromDec("-4134320184329015710493829104")
		expected := "14003764110508630918697735423914335136046266814161"

		z := new(Int).Quo(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("-578960446186580977117854925043439539266349923328202564819961")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "0"

		z := new(Int).Quo(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("5. should return correct result", func(t *testing.T) {
		x := MustFromDec("3214732915732914729142135321421")
		y := MustFromDec("510482085320157124")
		expected := "6297445117426"

		z := new(Int).Quo(x, y)
		assert.Equal(t, expected, z.Dec())
	})
}

func TestMul(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("100")
		y := MustFromDec("-2")
		expected := "-200"
		z := new(Int).Mul(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		z := new(Int).Mul(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "1"
		z := new(Int).Mul(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "0"
		z := new(Int).Mul(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("5. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "0"
		z := new(Int).Mul(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("6. should return correct result", func(t *testing.T) {
		x := MustFromDec("412421424314214830214")
		y := MustFromDec("491735014023482390148157914")
		expected := "202802054868735010725494286098171881252370413596"
		z := new(Int).Mul(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("7. should return correct result", func(t *testing.T) {
		x := MustFromDec("12")
		y := MustFromDec("12")
		expected := "144"
		z := new(Int).Mul(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("8. should return correct result", func(t *testing.T) {
		x := MustFromDec("-500000000000000")
		y := MustFromDec("5000000000000")
		expected := "-2500000000000000000000000000"
		z := new(Int).Mul(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("9. should return correct result", func(t *testing.T) {
		x := MustFromDec("-412421424314214830214")
		y := MustFromDec("491735014023482390148157914")
		expected := "-202802054868735010725494286098171881252370413596"
		z := new(Int).Mul(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("10. should return correct result", func(t *testing.T) {
		x := MustFromDec("202802054868735010725494286098171881252370413596")
		y := MustFromDec("-202802054868735010725494286098171881252370413596")
		expected := "17633904406147578101277522337283808353988667996273383726665993364273483857136"
		z := new(Int).Mul(x, y)
		assert.Equal(t, expected, z.Dec())
	})
}

func TestIsInt64(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		// max int64
		x := MustFromDec("9223372036854775807")
		assert.True(t, x.IsInt64())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		// min int64
		x := MustFromDec("-9223372036854775808")
		assert.True(t, x.IsInt64())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		// max int64 + 1
		x := MustFromDec("9223372036854775808")
		assert.False(t, x.IsInt64())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		// min int64 - 1
		x := MustFromDec("-9223372036854775809")
		assert.False(t, x.IsInt64())
	})

	t.Run("5. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		assert.True(t, x.IsInt64())
	})
}

func TestPow(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("2")
		n := uint64(8)
		expected := "256"
		z := new(Int).Pow(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		n := uint64(10)
		expected := "0"
		z := new(Int).Pow(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819")
		n := uint64(10)
		expected := "13268422908299897685045269624725657137437981925847170761094850041536535620361"
		z := new(Int).Pow(x, n)
		assert.Equal(t, expected, z.Dec())
	})
}

// func TestRem(t *testing.T) {
// 	t.Run("1. should return correct result", func(t *testing.T) {
// 		x := MustFromDec("3")
// 		y := MustFromDec("-2")
// 		expected := "1"
// 		z := new(Int).Rem(x, y)
// 		assert.Equal(t, expected, z.Dec())
// 	})

// 	t.Run("2. should return correct result", func(t *testing.T) {
// 		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
// 		y := MustFromDec("35719473219571942749314729421")
// 		expected := "-34167184328512991083512640217"
// 		z := new(Int).Rem(x, y)
// 		assert.Equal(t, expected, z.Dec())
// 	})

// 	t.Run("3. should return correct result", func(t *testing.T) {
// 		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
// 		y := MustFromDec("-35719473219571942749314729421")
// 		expected := "34167184328512991083512640216"
// 		z := new(Int).Rem(x, y)
// 		assert.Equal(t, expected, z.Dec())
// 	})

// 	t.Run("4. should return correct result", func(t *testing.T) {
// 		x := MustFromDec("-43179374921751324719573491471294")
// 		y := MustFromDec("-43127519734921524")
// 		expected := "-22155214380278890"
// 		z := new(Int).Rem(x, y)
// 		assert.Equal(t, expected, z.Dec())
// 	})

// 	t.Run("5. should return correct result", func(t *testing.T) {
// 		x := MustFromDec("0")
// 		y := MustFromDec("-2")
// 		expected := "0"
// 		z := new(Int).Rem(x, y)
// 		assert.Equal(t, expected, z.Dec())
// 	})

// 	t.Run("6. should return correct result", func(t *testing.T) {
// 		x := MustFromDec("4372195701247205721942146816424")
// 		y := MustFromDec("1974924792517421647328142")
// 		expected := "549633341738318150337156"
// 		z := new(Int).Rem(x, y)
// 		assert.Equal(t, expected, z.Dec())
// 	})
// }

func TestAddOverflow(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564606844")
		y := MustFromDec("-431294739547329532759")
		expected := "57896044618658097711785492504343953926634992332820282019297497264409235500333"
		z, overflow := new(Int).AddOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.True(t, overflow)
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "0"
		z, overflow := new(Int).AddOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.True(t, overflow)
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "-1"
		z, overflow := new(Int).AddOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.False(t, overflow)
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("10000")
		y := MustFromDec("-2000000")
		expected := "-1990000"
		z, overflow := new(Int).AddOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.False(t, overflow)
	})

	t.Run("5. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "-2"
		z, overflow := new(Int).AddOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.True(t, overflow)
	})

	t.Run("6. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819961")
		y := MustFromDec("431405283104328105143242031570231414")
		expected := "-57896044618658097711785492504343953926634560927537177691623648761924994588561"
		z := new(Int).Add(x, y)
		assert.Equal(t, expected, z.Dec())

	})
}

func TestSubOverflow(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("100")
		expected := "57896044618658097711785492504343953926634992332820282019728792003956564819868"
		z, overflow := new(Int).SubOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.True(t, overflow)
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		y := MustFromDec("9999")
		expected := "-9999"
		z, overflow := new(Int).SubOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.False(t, overflow)
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "-1"
		z, overflow := new(Int).SubOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.True(t, overflow)
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "1"
		z, overflow := new(Int).SubOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.True(t, overflow)
	})
}

func TestMulOverFlow(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("100")
		y := MustFromDec("-2")
		expected := "-200"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.False(t, overflow)
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.True(t, overflow)
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "1"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.True(t, overflow)
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "0"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.True(t, overflow)
	})

	t.Run("5. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		expected := "0"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.False(t, overflow)
	})

	t.Run("6. should return correct result", func(t *testing.T) {
		x := MustFromDec("412421424314214830214")
		y := MustFromDec("491735014023482390148157914")
		expected := "202802054868735010725494286098171881252370413596"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.False(t, overflow)
	})

	t.Run("7. should return correct result", func(t *testing.T) {
		x := MustFromDec("12")
		y := MustFromDec("12")
		expected := "144"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.False(t, overflow)
	})

	t.Run("8. should return correct result", func(t *testing.T) {
		x := MustFromDec("-500000000000000")
		y := MustFromDec("5000000000000")
		expected := "-2500000000000000000000000000"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.False(t, overflow)
	})

	t.Run("9. should return correct result", func(t *testing.T) {
		x := MustFromDec("-412421424314214830214")
		y := MustFromDec("491735014023482390148157914")
		expected := "-202802054868735010725494286098171881252370413596"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.False(t, overflow)
	})

	t.Run("10. should return correct result", func(t *testing.T) {
		x := MustFromDec("202802054868735010725494286098171881252370413596")
		y := MustFromDec("-202802054868735010725494286098171881252370413596")
		expected := "17633904406147578101277522337283808353988667996273383726665993364273483857136"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.True(t, overflow)
	})
}
