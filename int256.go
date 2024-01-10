package int256

import (
	"errors"
	"math"
	"math/bits"
)

type Int [4]uint64

var (
	MinI256 = &Int{0, 0, 0, 0x8000000000000000}

	ErrZeroDivision = errors.New("zero division")
)

func NewInt(val int64) *Int {
	z := &Int{}
	z.SetInt64(val)
	return z
}

func (z *Int) Set(x *Int) *Int {
	z[0], z[1], z[2], z[3] = x[0], x[1], x[2], x[3]
	return z
}

func (z *Int) SetInt64(x int64) *Int {
	if x >= 0 {
		z[3], z[2], z[1], z[0] = 0, 0, 0, uint64(x)
		return z
	}

	z[3], z[2], z[1], z[0] = 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff, uint64(x)
	return z
}

func (z *Int) IsInt64() bool {
	return ((z[1]|z[2]|z[3]) == 0 && z[0] <= 0x7fffffffffffffff) || // zero or positive int64
		((z[1]&z[2]&z[3]) == 0xffffffffffffffff && z[0] >= 0x8000000000000000) // negative int64
}

func (z *Int) Int64() int64 {
	s := z.Sign()
	if s == 0 {
		return 0
	}
	if s > 0 {
		// overflow when z[0] > math.MaxInt64
		return int64(z[0])
	}
	// -(2^64 - z[0])
	return -int64(math.MaxUint64 - z[0] + 1)
}

func (z *Int) SetUint64(x uint64) *Int {
	z[3], z[2], z[1], z[0] = 0, 0, 0, x
	return z
}

func (z *Int) IsUint64() bool {
	return (z[1] | z[2] | z[3]) == 0
}

func (z *Int) Uint64() uint64 {
	return z[0]
}

func (z *Int) Sign() int {
	if z.IsZero() {
		return 0
	}
	if z[3]&0x8000000000000000 == 0 {
		return 1
	}
	return -1
}

func (z *Int) IsZero() bool {
	return (z[0] | z[1] | z[2] | z[3]) == 0
}

func (z *Int) IsNegative() bool {
	return z[3]&0x8000000000000000 != 0
}

func (z *Int) IsPositive() bool {
	return z[3]&0x8000000000000000 == 0
}

func (z *Int) IsMinI256() bool {
	return (z[3] == 0x8000000000000000) && ((z[2] | z[1] | z[0]) == 0)
}

func (z *Int) Neg(x *Int) *Int {
	var carry uint64
	z[0], z[1], z[2], z[3] = ^x[0], ^x[1], ^x[2], ^x[3]
	z[0], carry = bits.Add64(z[0], 1, 0)
	z[1], carry = bits.Add64(z[1], 0, carry)
	z[2], carry = bits.Add64(z[2], 0, carry)
	z[3] += carry
	return z
}

func (z *Int) Eq(x *Int) bool {
	return (z[0] == x[0]) && (z[1] == x[1]) && (z[2] == x[2]) && (z[3] == x[3])
}

func (z *Int) Add(x, y *Int) *Int {
	var carry uint64
	z[0], carry = bits.Add64(x[0], y[0], 0)
	z[1], carry = bits.Add64(x[1], y[1], carry)
	z[2], carry = bits.Add64(x[2], y[2], carry)
	z[3] = x[3] + y[3] + carry
	return z
}

func (z *Int) Sub(x, y *Int) *Int {
	var carry uint64
	z[0], carry = bits.Sub64(x[0], y[0], 0)
	z[1], carry = bits.Sub64(x[1], y[1], carry)
	z[2], carry = bits.Sub64(x[2], y[2], carry)
	z[3] = x[3] - y[3] - carry
	return z
}

func (z *Int) Mul(x, y *Int) *Int {
	var (
		res              Int
		carry            uint64
		res1, res2, res3 uint64
	)

	carry, res[0] = bits.Mul64(x[0], y[0])
	carry, res1 = umulHop(carry, x[1], y[0])
	carry, res2 = umulHop(carry, x[2], y[0])
	res3 = x[3]*y[0] + carry

	carry, res[1] = umulHop(res1, x[0], y[1])
	carry, res2 = umulStep(res2, x[1], y[1], carry)
	res3 = res3 + x[2]*y[1] + carry

	carry, res[2] = umulHop(res2, x[0], y[2])
	res3 = res3 + x[1]*y[2] + carry

	res[3] = res3 + x[0]*y[3]

	return z.Set(&res)
}

func umulStep(z, x, y, carry uint64) (hi, lo uint64) {
	hi, lo = bits.Mul64(x, y)
	lo, carry = bits.Add64(lo, carry, 0)
	hi += carry
	lo, carry = bits.Add64(lo, z, 0)
	hi += carry
	return hi, lo
}

func umulHop(z, x, y uint64) (hi, lo uint64) {
	hi, lo = bits.Mul64(x, y)
	lo, carry := bits.Add64(lo, z, 0)
	hi += carry
	return hi, lo
}

func (z *Int) Clear() *Int {
	z[0], z[1], z[2], z[3] = 0, 0, 0, 0
	return z
}

func (z *Int) SetOne() *Int {
	z[3], z[2], z[1], z[0] = 0, 0, 0, 1
	return z
}

func (z *Int) Div(x, y *Int) *Int {
	if x.Sign() > 0 {
		if y.Sign() > 0 {
			z.div(x, y)
			return z
		} else {
			z.div(x, new(Int).Neg(y))
			return z.Neg(z)
		}
	}
	if y.Sign() < 0 {
		z.div(new(Int).Neg(x), new(Int).Neg(y))
		return z
	}
	z.div(new(Int).Neg(x), y)
	return z.Neg(z)
}

func (z *Int) div(x, y *Int) *Int {
	if y.IsZero() {
		panic(ErrZeroDivision)
	}
	if x.IsZero() {
		return z.Clear()
	}
	if x.Eq(y) {
		return z.SetOne()
	}
	if x.IsInt64() && y.IsInt64() {
		return z.SetInt64(x.Int64() / y.Int64())
	}
	var quot Int
	udivrem(quot[:], x[:], y)
	return z.Set(&quot)
}

func (z *Int) Pow(x *Int, n uint64) *Int {
	z.SetOne()
	if n == 0 {
		return z
	}
	for n > 0 {
		if n&1 == 1 {
			z.Mul(z, x)
		}
		n >>= 1
		x.Mul(x, x)
	}
	return z
}

func (z *Int) Clone() *Int {
	return &Int{z[0], z[1], z[2], z[3]}
}
