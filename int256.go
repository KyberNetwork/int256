package int256

import (
	"errors"
	"math"
	"math/bits"
)

type Int [4]uint64

var (
	MinI256 = &Int{0, 0, 0, 0x8000000000000000}
	MaxI256 = &Int{0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff, 0x7fffffffffffffff}

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

func (z *Int) IsOne() bool {
	return (z[0] == 1) && (z[1]|z[2]|z[3]) == 0
}

func (z *Int) IsNegative() bool {
	return z[3]&0x8000000000000000 != 0
}

func (z *Int) IsPositive() bool {
	return (z[3]&0x8000000000000000) == 0 && (z[3]|z[2]|z[1]|z[0]) != 0
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

func (z *Int) AddOverflow(x, y *Int) (*Int, bool) {
	var carry uint64
	z[0], carry = bits.Add64(x[0], y[0], 0)
	z[1], carry = bits.Add64(x[1], y[1], carry)
	z[2], carry = bits.Add64(x[2], y[2], carry)
	z[3] = x[3] + y[3] + carry
	var overflow bool
	signX, signY, signZ := x.Sign(), y.Sign(), z.Sign()
	if (signX == signY) && (signX != signZ) {
		overflow = true
	}
	return z, overflow
}

func (z *Int) Sub(x, y *Int) *Int {
	var carry uint64
	z[0], carry = bits.Sub64(x[0], y[0], 0)
	z[1], carry = bits.Sub64(x[1], y[1], carry)
	z[2], carry = bits.Sub64(x[2], y[2], carry)
	z[3] = x[3] - y[3] - carry
	return z
}

func (z *Int) SubOverflow(x, y *Int) (*Int, bool) {
	var carry uint64
	z[0], carry = bits.Sub64(x[0], y[0], 0)
	z[1], carry = bits.Sub64(x[1], y[1], carry)
	z[2], carry = bits.Sub64(x[2], y[2], carry)
	z[3] = x[3] - y[3] - carry
	var overflow bool
	signX, signY, signZ := x.Sign(), y.Sign(), z.Sign()
	if (signX == 0 && y.IsMinI256()) || ((signX != 0) && (signX != signY) && (signX != signZ)) {
		overflow = true
	}
	return z, overflow
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

func (z *Int) MulOverflow(x, y *Int) (*Int, bool) {
	if (x.IsMinI256() && y.IsOne()) || (x.IsOne() && y.IsMinI256()) {
		return z.Set(MinI256), false
	}

	var flipSign bool
	xSign, ySign := x.Sign(), y.Sign()
	if xSign*ySign == -1 {
		flipSign = true
	}
	if xSign < 0 {
		x.Neg(x)
	}
	if ySign < 0 {
		y.Neg(y)
	}

	p := umul(x, y)
	z[0], z[1], z[2], z[3] = p[0], p[1], p[2], p[3]

	var overflow bool
	if ((p[4] | p[5] | p[6] | p[7]) != 0) || z.IsNegative() {
		overflow = true
	}

	if flipSign {
		z.Neg(z)
	}

	return z, overflow
}

func umul(x, y *Int) [8]uint64 {
	var (
		res                           [8]uint64
		carry, carry4, carry5, carry6 uint64
		res1, res2, res3, res4, res5  uint64
	)

	carry, res[0] = bits.Mul64(x[0], y[0])
	carry, res1 = umulHop(carry, x[1], y[0])
	carry, res2 = umulHop(carry, x[2], y[0])
	carry4, res3 = umulHop(carry, x[3], y[0])

	carry, res[1] = umulHop(res1, x[0], y[1])
	carry, res2 = umulStep(res2, x[1], y[1], carry)
	carry, res3 = umulStep(res3, x[2], y[1], carry)
	carry5, res4 = umulStep(carry4, x[3], y[1], carry)

	carry, res[2] = umulHop(res2, x[0], y[2])
	carry, res3 = umulStep(res3, x[1], y[2], carry)
	carry, res4 = umulStep(res4, x[2], y[2], carry)
	carry6, res5 = umulStep(carry5, x[3], y[2], carry)

	carry, res[3] = umulHop(res3, x[0], y[3])
	carry, res[4] = umulStep(res4, x[1], y[3], carry)
	carry, res[5] = umulStep(res5, x[2], y[3], carry)
	res[7], res[6] = umulStep(carry6, x[3], y[3], carry)

	return res
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

func (z *Int) Quo(x, y *Int) *Int {
	if x.Sign() > 0 {
		if y.Sign() > 0 {
			return z.uquo(x, y)
		}
		z.uquo(x, new(Int).Neg(y))
		return z.Neg(z)
	}
	if y.Sign() < 0 {
		return z.uquo(new(Int).Neg(x), new(Int).Neg(y))
	}
	z.uquo(new(Int).Neg(x), y)
	return z.Neg(z)
}

func (z *Int) uquo(x, y *Int) *Int {
	if y.IsZero() {
		panic(ErrZeroDivision)
	}
	if x.IsZero() {
		z.Clear()
	}
	if x.Eq(y) {
		return z.SetOne()
	}
	if x.IsInt64() && y.IsInt64() {
		return z.SetInt64(x.Int64() / y.Int64())
	}
	quot := Int{}
	udivrem(quot[:], x[:], y)
	return z.Set(&quot)
}

func (z *Int) Rem(x, y *Int) *Int {
	if x.Sign() > 0 {
		if y.Sign() > 0 {
			return z.urem(x, y)
		}
		return z.urem(x, new(Int).Neg(y))
	}
	if y.Sign() < 0 {
		z.urem(new(Int).Neg(x), new(Int).Neg(y))
		return z.Neg(z)
	}
	z.urem(new(Int).Neg(x), y)
	return z.Neg(z)
}

func (z *Int) urem(x, y *Int) *Int {
	if y.IsZero() {
		panic(ErrZeroDivision)
	}
	if x.IsZero() {
		z.Clear()
	}
	if x.Eq(y) {
		return z.Clear()
	}
	if x.IsInt64() && y.IsInt64() {
		xInt64 := x.Int64()
		yInt64 := y.Int64()
		return z.SetInt64(xInt64 % yInt64)
	}
	quot := Int{}
	rem := udivrem(quot[:], x[:], y)
	return z.Set(&rem)
}

func (z *Int) Pow(x *Int, n uint64) *Int {
	z.SetOne()
	for n > 0 {
		if n&1 == 1 {
			z.Mul(z, x)
		}
		n >>= 1
		x.Mul(x, x)
	}
	return z
}

func (z *Int) Lt(x *Int) bool {
	return z.Cmp(x) < 0
}

func (z *Int) Lte(x *Int) bool {
	return z.Cmp(x) <= 0
}

func (z *Int) Gt(x *Int) bool {
	return z.Cmp(x) > 0
}

func (z *Int) Gte(x *Int) bool {
	return z.Cmp(x) >= 0
}

func (z *Int) Cmp(x *Int) int {
	zSign, xSign := z.Sign(), x.Sign()
	if zSign != xSign {
		if zSign > xSign {
			return 1
		}
		return -1
	}
	if zSign == 0 {
		return 0
	}
	for i := 3; i >= 0; i-- {
		if z[i] > x[i] {
			return 1
		}
		if z[i] < x[i] {
			return -1
		}
	}
	return 0
}

func (z *Int) Clone() *Int {
	return &Int{z[0], z[1], z[2], z[3]}
}
