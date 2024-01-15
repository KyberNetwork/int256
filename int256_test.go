package int256

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInt(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		v := int64(math.MaxInt64)
		z := NewInt(v)
		assert.True(t, z.IsInt64())
		assert.Equal(t, v, z.Int64())
	})

	t.Run("negative", func(t *testing.T) {
		v := int64(math.MinInt64)
		z := NewInt(v)
		assert.True(t, z.IsInt64())
		assert.Equal(t, v, z.Int64())
	})

	t.Run("zero", func(t *testing.T) {
		v := int64(0)
		z := NewInt(v)
		assert.True(t, z.IsInt64())
		assert.Equal(t, v, z.Int64())
	})
}

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

	t.Run("6. should return correct result", func(t *testing.T) {
		x := MustFromDec("510482085320157124")
		y := MustFromDec("510482085320157124")
		expected := "1"
		z := new(Int).Quo(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("7. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		y := MustFromDec("510482085320157124")
		expected := "0"
		z := new(Int).Quo(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("8. should panic error", func(t *testing.T) {
		x := MustFromDec("1")
		y := MustFromDec("0")
		assert.Panics(t, func() { new(Int).Quo(x, y) })
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

func TestIsUint64(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		// max uint64
		x := MustFromDec("18446744073709551615")
		assert.True(t, x.IsUint64())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		// max uint64 + 1
		x := MustFromDec("18446744073709551616")
		assert.False(t, x.IsUint64())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		assert.True(t, x.IsUint64())
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

func TestRem(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("3")
		y := MustFromDec("-2")
		expected := "1"
		z := new(Int).Rem(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("35719473219571942749314729421")
		expected := "-34167184328512991083512640217"
		z := new(Int).Rem(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-35719473219571942749314729421")
		expected := "34167184328512991083512640216"
		z := new(Int).Rem(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("-43179374921751324719573491471294")
		y := MustFromDec("-43127519734921524")
		expected := "-22155214380278890"
		z := new(Int).Rem(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("5. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		y := MustFromDec("-2")
		expected := "0"
		z := new(Int).Rem(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("6. should return correct result", func(t *testing.T) {
		x := MustFromDec("4372195701247205721942146816424")
		y := MustFromDec("1974924792517421647328142")
		expected := "549633341738318150337156"
		z := new(Int).Rem(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("7. should return correct error", func(t *testing.T) {
		x := MustFromDec("1974924792517421647328142")
		y := MustFromDec("1974924792517421647328142")
		expected := "0"
		z := new(Int).Rem(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("8. should panic error", func(t *testing.T) {
		x := MustFromDec("1")
		y := MustFromDec("0")
		assert.Panics(t, func() { new(Int).Rem(x, y) })
	})
}

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

	t.Run("11. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("1")
		expected := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.False(t, overflow)
	})

	t.Run("12. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("-1")
		expected := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		z, overflow := new(Int).MulOverflow(x, y)
		assert.Equal(t, expected, z.Dec())
		assert.True(t, overflow)
	})
}

func TestIsPostive(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("100")
		assert.True(t, x.IsPositive())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("-100")
		assert.False(t, x.IsPositive())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		assert.False(t, x.IsPositive())
	})
}

func TestCmp(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.Equal(t, -1, x.Cmp(y))
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		assert.Equal(t, 0, x.Cmp(y))
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		assert.Equal(t, 1, x.Cmp(y))
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.Equal(t, 0, x.Cmp(y))
	})

	t.Run("5. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		y := MustFromDec("0")
		assert.Equal(t, 0, x.Cmp(y))
	})

	t.Run("6. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		y := MustFromDec("-10")
		assert.Equal(t, 1, x.Cmp(y))
	})

	t.Run("7. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.Equal(t, -1, x.Cmp(y))
	})

	t.Run("8. should return correct result", func(t *testing.T) {
		x := MustFromDec("-578960446186580977117854925043439539266349923328202820197287920039565648")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.Equal(t, 1, x.Cmp(y))
	})
}

func TestLt(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.True(t, x.Lt(y))
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.False(t, x.Lt(y))
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		assert.False(t, x.Lt(y))
	})
}

func TestLte(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.True(t, x.Lte(y))
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.True(t, x.Lte(y))
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		assert.False(t, x.Lte(y))
	})
}

func TestGt(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.False(t, x.Gt(y))
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.False(t, x.Gt(y))
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		assert.True(t, x.Gt(y))
	})
}

func TestGte(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.False(t, x.Gte(y))
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		assert.True(t, x.Gte(y))
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		assert.True(t, x.Gte(y))
	})
}

func TestClone(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("100")
		y := x.Clone()
		assert.Equal(t, x.Dec(), y.Dec())
		assert.NotEqual(t, fmt.Sprintf("%p", x), fmt.Sprintf("%p", y))
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("-100")
		y := x.Clone()
		assert.Equal(t, x.Dec(), y.Dec())
		assert.NotEqual(t, fmt.Sprintf("%p", x), fmt.Sprintf("%p", y))
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		y := x.Clone()
		assert.Equal(t, x.Dec(), y.Dec())
		assert.NotEqual(t, fmt.Sprintf("%p", x), fmt.Sprintf("%p", y))
	})
}

func TestAnd(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("-100")
		expected := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		z := new(Int).And(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("41457129491534261876432718654783265437285638275")
		expected := "41457129491534261876432718654783265437285638275"
		z := new(Int).And(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("7")
		expected := "1"
		z := new(Int).And(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("0")
		expected := "0"
		z := new(Int).And(x, y)
		assert.Equal(t, expected, z.Dec())
	})
}

func TestOr(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("-100")
		expected := "-100"
		z := new(Int).Or(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("41457129491534261876432718654783265437285638275")
		expected := "57896044618658097711785492504343953926634992332820282019728792003956564819967"
		z := new(Int).Or(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "-1"
		z := new(Int).Or(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("0")
		expected := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		z := new(Int).Or(x, y)
		assert.Equal(t, expected, z.Dec())
	})
}

func TestXor(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("-100")
		expected := "57896044618658097711785492504343953926634992332820282019728792003956564819868"
		z := new(Int).Xor(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		y := MustFromDec("41457129491534261876432718654783265437285638275")
		expected := "57896044618658097711785492504302496797143458070943849301074008738519279181692"
		z := new(Int).Xor(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "-1"
		z := new(Int).Xor(x, y)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
		y := MustFromDec("0")
		expected := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		z := new(Int).Xor(x, y)
		assert.Equal(t, expected, z.Dec())
	})
}

func TestNot(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("100")
		expected := "-101"
		z := new(Int).Not(x)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("-100")
		expected := "99"
		z := new(Int).Not(x)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		expected := "-1"
		z := new(Int).Not(x)
		assert.Equal(t, expected, z.Dec())
	})
}

func TestLsh(t *testing.T) {
	t.Run("1. n is equal 0", func(t *testing.T) {
		n := uint(0)
		x := MustFromDec("100")
		expected := "100"
		z := new(Int).Lsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. n is greater than or equal 256", func(t *testing.T) {
		n := uint(256)
		x := MustFromDec("-4125871947195612497219427349")
		expected := "0"
		z := new(Int).Lsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. x is non-neg & n is smaller than 256 and greater than or equal 192,", func(t *testing.T) {
		n := uint(200)
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "-1606938044258990275541962092341162602522202993782792835301376"
		z := new(Int).Lsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. x is non-neg & n is smaller than 192 and greater than or equal 128", func(t *testing.T) {
		n := uint(150)
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "-1427247692705959881058285969449495136382746624"
		z := new(Int).Lsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("5. x is non-neg & n is smaller than 128 and greater than or equal 64", func(t *testing.T) {
		n := uint(100)
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "-1267650600228229401496703205376"
		z := new(Int).Lsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("6. x is non-neg & n is smaller than 64", func(t *testing.T) {
		n := uint(32)
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "-4294967296"
		z := new(Int).Lsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("7. x is neg & n is smaller than 256 and greater than or equal 192,", func(t *testing.T) {
		n := uint(200)
		x := MustFromDec("-11111")
		expected := "-17854688609761640951546740808002657676624197463920611193033588736"
		z := new(Int).Lsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("8. x is neg & n is smaller than 192 and greater than or equal 128", func(t *testing.T) {
		n := uint(150)
		x := MustFromDec("-11111")
		expected := "-15858149113655920238438615406553340460348697739264"
		z := new(Int).Lsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("9. x is neg & n is smaller than 128 and greater than or equal 64", func(t *testing.T) {
		n := uint(100)
		x := MustFromDec("-11111")
		expected := "-14084865819135856880029869314932736"
		z := new(Int).Lsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("10. x is neg & n is smaller than 64", func(t *testing.T) {
		n := uint(32)
		x := MustFromDec("-11111")
		expected := "-47721381625856"
		z := new(Int).Lsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})
}

func TestRsh(t *testing.T) {
	t.Run("1. n is equal 0", func(t *testing.T) {
		n := uint(0)
		x := MustFromDec("100")
		expected := "100"
		z := new(Int).Rsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. x is non-neg is greater than or equal 255", func(t *testing.T) {
		n := uint(255)
		x := MustFromDec("431247391574329147932")
		expected := "0"
		z := new(Int).Rsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. x is non-neg & n is smaller than 256 and greater than or equal 192,", func(t *testing.T) {
		n := uint(200)
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "36028797018963967"
		z := new(Int).Rsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. x is non-neg & n is smaller than 192 and greater than or equal 128", func(t *testing.T) {
		n := uint(150)
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "40564819207303340847894502572031"
		z := new(Int).Rsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("5. x is non-neg & n is smaller than 128 and greater than or equal 64", func(t *testing.T) {
		n := uint(100)
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "45671926166590716193865151022383844364247891967"
		z := new(Int).Rsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("6. x is non-neg & n is smaller than 64", func(t *testing.T) {
		n := uint(32)
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "13479973333575319897333507543509815336818572211270286240551805124607"
		z := new(Int).Rsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("7. x is neg & n is smaller than 256 and greater than or equal 192,", func(t *testing.T) {
		n := uint(200)
		x := MustFromDec("-11111")
		expected := "-1"
		z := new(Int).Rsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("8. x is neg & n is smaller than 192 and greater than or equal 128", func(t *testing.T) {
		n := uint(150)
		x := MustFromDec("-11111")
		expected := "-1"
		z := new(Int).Rsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("9. x is neg & n is smaller than 128 and greater than or equal 64", func(t *testing.T) {
		n := uint(100)
		x := MustFromDec("-11111")
		expected := "-1"
		z := new(Int).Rsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("10. x is neg & n is smaller than 64", func(t *testing.T) {
		n := uint(2)
		x := MustFromDec("-11111")
		expected := "-2778"
		z := new(Int).Rsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("11. x is neg & n is greater than or equal 255", func(t *testing.T) {
		n := uint(255)
		x := MustFromDec("-11111")
		expected := "-1"
		z := new(Int).Rsh(x, n)
		assert.Equal(t, expected, z.Dec())
	})
}

func TestSqrt(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("100")
		expected := "10"
		z := new(Int).Sqrt(x)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("57896044618658097711785492504343953926634992332820282019728792003956564819967")
		expected := "240615969168004511545033772477625056927"
		z := new(Int).Sqrt(x)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("13795704712047281502140217429180472156210462196574839653795")
		expected := "117455117862302115618237725775"
		z := new(Int).Sqrt(x)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		expected := "0"
		z := new(Int).Sqrt(x)
		assert.Equal(t, expected, z.Dec())
	})

	t.Run("5. should panic negative number", func(t *testing.T) {
		x := MustFromDec("-1000")
		assert.Panics(t, func() { new(Int).Sqrt(x) })
	})
}
