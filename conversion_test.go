package int256

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		z := new(Int)
		z[3], z[2], z[1], z[0] = 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff
		expected := "-1"
		actual := z.Dec()
		assert.Equal(t, expected, actual)
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		z := new(Int)
		z[3], z[2], z[1], z[0] = 0x8000000000000000, 0, 0, 0
		// min int256
		expected := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		actual := z.Dec()
		assert.Equal(t, expected, actual)
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		z := new(Int)
		z[3], z[2], z[1], z[0] = 0xffffffffffffffff, 0xfffffffffefc6299, 0xa769780616fe2a60, 0xd6f11e625171ce1f
		expected := "-5789604461865809771178549250434395392663499233"
		actual := z.Dec()
		assert.Equal(t, expected, actual)
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		z := new(Int)
		z[3], z[2], z[1], z[0] = 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff, 0x8000000000000000
		// min int64
		expected := "-9223372036854775808"
		actual := z.Dec()
		assert.Equal(t, expected, actual)
	})

	t.Run("5. should return correct result", func(t *testing.T) {
		z := new(Int)
		z[3], z[2], z[1], z[0] = 0xfffffefd7ac52f61, 0xb0fd0daa972f29e5, 0x0f79c8be960c7412, 0x7ac05fd7df5bb137
		expected := "-6969696969969696966996969669696969696969696969969696969696996969696969"
		actual := z.Dec()
		assert.Equal(t, expected, actual)
	})

	t.Run("6. should return correct result", func(t *testing.T) {
		z := new(Int)
		z[3], z[2], z[1], z[0] = 0, 0, 0, 0
		expected := "0"
		actual := z.Dec()
		assert.Equal(t, expected, actual)
	})

	t.Run("7. should return correct result", func(t *testing.T) {
		z := new(Int)
		z[3], z[2], z[1], z[0] = 0x7fffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff
		// max int256
		expected := "57896044618658097711785492504343953926634992332820282019728792003956564819967"
		actual := z.Dec()
		assert.Equal(t, expected, actual)
	})

	t.Run("8. should return correct result", func(t *testing.T) {
		z := new(Int)
		z[3], z[2], z[1], z[0] = 0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x7fffffffffffffff
		// max int256
		expected := "9223372036854775807"
		actual := z.Dec()
		assert.Equal(t, expected, actual)
	})

	t.Run("9. should return correct result", func(t *testing.T) {
		z := new(Int)
		z[3], z[2], z[1], z[0] = 0x0000010285417177, 0x065c2836aab5cb46, 0x5ddad3e004792537, 0xde600447c77583e0
		expected := "6969699696696969696696969696969696966969696969696969669696969696969696"
		actual := z.Dec()
		assert.Equal(t, expected, actual)
	})

	t.Run("10. should return correct result", func(t *testing.T) {
		z := new(Int)
		z[3], z[2], z[1], z[0] = 0x0000000000000000, 0x00000008727f5e06, 0x86e2a46fd1c06d19, 0xb46fb9323a6d889d
		expected := "12345678431937219573219471295439254379564372953245"
		actual := z.Dec()
		assert.Equal(t, expected, actual)
	})
}

func TestMustFromDec(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		dec := "-1"
		z := MustFromDec(dec)
		assert.Equal(t, z[3], uint64(0xffffffffffffffff))
		assert.Equal(t, z[2], uint64(0xffffffffffffffff))
		assert.Equal(t, z[1], uint64(0xffffffffffffffff))
		assert.Equal(t, z[0], uint64(0xffffffffffffffff))
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		dec := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		z := MustFromDec(dec)
		assert.Equal(t, z[3], uint64(0x8000000000000000))
		assert.Equal(t, z[2], uint64(0))
		assert.Equal(t, z[1], uint64(0))
		assert.Equal(t, z[0], uint64(0))
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		dec := "-5789604461865809771178549250434395392663499233"
		z := MustFromDec(dec)
		assert.Equal(t, z[3], uint64(0xffffffffffffffff))
		assert.Equal(t, z[2], uint64(0xfffffffffefc6299))
		assert.Equal(t, z[1], uint64(0xa769780616fe2a60))
		assert.Equal(t, z[0], uint64(0xd6f11e625171ce1f))
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		dec := "-9223372036854775808"
		z := MustFromDec(dec)
		assert.Equal(t, z[3], uint64(0xffffffffffffffff))
		assert.Equal(t, z[2], uint64(0xffffffffffffffff))
		assert.Equal(t, z[1], uint64(0xffffffffffffffff))
		assert.Equal(t, z[0], uint64(0x8000000000000000))
	})

	t.Run("5. should return correct result", func(t *testing.T) {
		dec := "-6969696969969696966996969669696969696969696969969696969696996969696969"
		z := MustFromDec(dec)
		assert.Equal(t, z[3], uint64(0xfffffefd7ac52f61))
		assert.Equal(t, z[2], uint64(0xb0fd0daa972f29e5))
		assert.Equal(t, z[1], uint64(0x0f79c8be960c7412))
		assert.Equal(t, z[0], uint64(0x7ac05fd7df5bb137))
	})

	t.Run("6. should return correct result", func(t *testing.T) {
		dec := "0"
		z := MustFromDec(dec)
		assert.Equal(t, z[3], uint64(0))
		assert.Equal(t, z[2], uint64(0))
		assert.Equal(t, z[1], uint64(0))
		assert.Equal(t, z[0], uint64(0))
	})

	t.Run("7. should return correct result", func(t *testing.T) {
		dec := "57896044618658097711785492504343953926634992332820282019728792003956564819967"
		z := MustFromDec(dec)
		assert.Equal(t, z[3], uint64(0x7fffffffffffffff))
		assert.Equal(t, z[2], uint64(0xffffffffffffffff))
		assert.Equal(t, z[1], uint64(0xffffffffffffffff))
		assert.Equal(t, z[0], uint64(0xffffffffffffffff))
	})

	t.Run("8. should return correct result", func(t *testing.T) {
		dec := "9223372036854775807"
		z := MustFromDec(dec)
		assert.Equal(t, z[3], uint64(0))
		assert.Equal(t, z[2], uint64(0))
		assert.Equal(t, z[1], uint64(0))
		assert.Equal(t, z[0], uint64(0x7fffffffffffffff))
	})

	t.Run("9. should return correct result", func(t *testing.T) {
		dec := "6969699696696969696696969696969696966969696969696969669696969696969696"
		z := MustFromDec(dec)
		assert.Equal(t, z[3], uint64(0x0000010285417177))
		assert.Equal(t, z[2], uint64(0x065c2836aab5cb46))
		assert.Equal(t, z[1], uint64(0x5ddad3e004792537))
		assert.Equal(t, z[0], uint64(0xde600447c77583e0))
	})

	t.Run("10. should return correct result", func(t *testing.T) {
		dec := "12345678431937219573219471295439254379564372953245"
		z := MustFromDec(dec)
		assert.Equal(t, z[3], uint64(0x0000000000000000))
		assert.Equal(t, z[2], uint64(0x00000008727f5e06))
		assert.Equal(t, z[1], uint64(0x86e2a46fd1c06d19))
		assert.Equal(t, z[0], uint64(0xb46fb9323a6d889d))
	})

	t.Run("11. should panic error", func(t *testing.T) {
		dec := "123456789XLXX"
		assert.Panics(t, func() { MustFromDec(dec) })
	})
}

func TestFromDec(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		dec := "12345678431937219573219471295439254379564372953245"
		z, err := FromDec(dec)
		assert.Nil(t, err)
		assert.Equal(t, z[3], uint64(0x0000000000000000))
		assert.Equal(t, z[2], uint64(0x00000008727f5e06))
		assert.Equal(t, z[1], uint64(0x86e2a46fd1c06d19))
		assert.Equal(t, z[0], uint64(0xb46fb9323a6d889d))
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		dec := "-0009223372036854775808"
		z := MustFromDec(dec)
		assert.Equal(t, z[3], uint64(0xffffffffffffffff))
		assert.Equal(t, z[2], uint64(0xffffffffffffffff))
		assert.Equal(t, z[1], uint64(0xffffffffffffffff))
		assert.Equal(t, z[0], uint64(0x8000000000000000))
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		dec := "00009223372036854775807"
		z, err := FromDec(dec)
		assert.Nil(t, err)
		assert.Equal(t, z[3], uint64(0))
		assert.Equal(t, z[2], uint64(0))
		assert.Equal(t, z[1], uint64(0))
		assert.Equal(t, z[0], uint64(0x7fffffffffffffff))
	})

	t.Run("4. should return error", func(t *testing.T) {
		dec := "#123"
		_, err := FromDec(dec)
		assert.Error(t, err)
	})

	t.Run("5. should return error", func(t *testing.T) {
		// (1 << 255)
		dec := "57896044618658097711785492504343953926634992332820282019728792003956564819968"
		_, err := FromDec(dec)
		assert.Error(t, err)
	})
}

func TestFromBig(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		v := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		b, _ := new(big.Int).SetString(v, 10)
		z, err := FromBig(b)
		assert.Nil(t, err)
		assert.Equal(t, v, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		v := "57896044618658097711785492504343953926634992332820282019728792003956564819967"
		b, _ := new(big.Int).SetString(v, 10)
		z, err := FromBig(b)
		assert.Nil(t, err)
		assert.Equal(t, v, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		v := "0"
		b, _ := new(big.Int).SetString(v, 10)
		z, err := FromBig(b)
		assert.Nil(t, err)
		assert.Equal(t, v, z.Dec())
	})

	t.Run("4. should return error overflow", func(t *testing.T) {
		v := "-57896044618658097711785492504343953926634992332820282019728792003956564819969"
		b, _ := new(big.Int).SetString(v, 10)
		_, err := FromBig(b)
		assert.ErrorIs(t, err, ErrOverflow)
	})

	t.Run("5. should return error overflow", func(t *testing.T) {
		v := "57896044618658097711785492504343953926634992332820282019728792003956564819968"
		b, _ := new(big.Int).SetString(v, 10)
		_, err := FromBig(b)
		assert.ErrorIs(t, err, ErrOverflow)
	})

	t.Run("6. should return error overflow", func(t *testing.T) {
		v := "57896044618658097711785492504343953926634992332820282019728792003956564819968431242"
		b, _ := new(big.Int).SetString(v, 10)
		_, err := FromBig(b)
		assert.ErrorIs(t, err, ErrOverflow)
	})

	t.Run("7. should return error overflow", func(t *testing.T) {
		v := "-57896044618658097711785492504343953926634992332820282019728792003956564819968431242"
		b, _ := new(big.Int).SetString(v, 10)
		_, err := FromBig(b)
		assert.ErrorIs(t, err, ErrOverflow)
	})
}

func TestMustFromBig(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		v := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		b, _ := new(big.Int).SetString(v, 10)
		z := MustFromBig(b)
		assert.Equal(t, v, z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		v := "57896044618658097711785492504343953926634992332820282019728792003956564819967"
		b, _ := new(big.Int).SetString(v, 10)
		z := MustFromBig(b)
		assert.Equal(t, v, z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		v := "0"
		b, _ := new(big.Int).SetString(v, 10)
		z := MustFromBig(b)
		assert.Equal(t, v, z.Dec())
	})

	t.Run("4. should panic error overflow", func(t *testing.T) {
		v := "-57896044618658097711785492504343953926634992332820282019728792003956564819969"
		b, _ := new(big.Int).SetString(v, 10)
		assert.Panics(t, func() { MustFromBig(b) })
	})

	t.Run("5. should panic error overflow", func(t *testing.T) {
		v := "57896044618658097711785492504343953926634992332820282019728792003956564819968"
		b, _ := new(big.Int).SetString(v, 10)
		assert.Panics(t, func() { MustFromBig(b) })
	})

	t.Run("6. should panic error overflow", func(t *testing.T) {
		v := "57896044618658097711785492504343953926634992332820282019728792003956564819968431242"
		b, _ := new(big.Int).SetString(v, 10)
		assert.Panics(t, func() { MustFromBig(b) })
	})

	t.Run("7. should panic error overflow", func(t *testing.T) {
		v := "-57896044618658097711785492504343953926634992332820282019728792003956564819968431242"
		b, _ := new(big.Int).SetString(v, 10)
		assert.Panics(t, func() { MustFromBig(b) })
	})
}

func TestToBig(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		v := "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
		z := MustFromDec(v)
		b := z.ToBig()
		assert.Equal(t, v, b.String())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		v := "57896044618658097711785492504343953926634992332820282019728792003956564819967"
		z := MustFromDec(v)
		b := z.ToBig()
		assert.Equal(t, v, b.String())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		v := "0"
		z := MustFromDec(v)
		b := z.ToBig()
		assert.Equal(t, v, b.String())
	})

	t.Run("4. should return correct result", func(t *testing.T) {
		v := "-9223372036854775808"
		z := MustFromDec(v)
		b := z.ToBig()
		assert.Equal(t, v, b.String())
	})

	t.Run("5. should return correct result", func(t *testing.T) {
		v := "-43217597390350847214095743109472109521"
		z := MustFromDec(v)
		b := z.ToBig()
		assert.Equal(t, v, b.String())
	})
}

func TestMarshalJson(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := MustFromDec("141243")
		b, err := json.Marshal(x)
		assert.Nil(t, err)
		assert.Equal(t, `"141243"`, string(b))
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := MustFromDec("-141243")
		b, err := json.Marshal(x)
		assert.Nil(t, err)
		assert.Equal(t, `"-141243"`, string(b))
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := MustFromDec("0")
		b, err := json.Marshal(x)
		assert.Nil(t, err)
		assert.Equal(t, `"0"`, string(b))
	})
}

func TestUnmarshalJson(t *testing.T) {
	t.Run("1. should return correct result", func(t *testing.T) {
		x := `"14214214"`
		var z Int
		err := json.Unmarshal([]byte(x), &z)
		assert.Nil(t, err)
		assert.Equal(t, "14214214", z.Dec())
	})

	t.Run("2. should return correct result", func(t *testing.T) {
		x := `14214214`
		var z Int
		err := json.Unmarshal([]byte(x), &z)
		assert.Nil(t, err)
		assert.Equal(t, "14214214", z.Dec())
	})

	t.Run("3. should return correct result", func(t *testing.T) {
		x := `-14214214`
		var z Int
		err := json.Unmarshal([]byte(x), &z)
		assert.Nil(t, err)
		assert.Equal(t, "-14214214", z.Dec())
	})
}
