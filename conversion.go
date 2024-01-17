package int256

import (
	"errors"
	"io"
	"math/big"
	"math/bits"
	"strconv"
)

const (
	maxAbsI256Dec = "57896044618658097711785492504343953926634992332820282019728792003956564819968"
	maxWords      = 256 / bits.UintSize
)

var (
	ErrOverflow = errors.New("int256: overflow")

	multipliers = [5]*Int{
		nil,
		{0x8ac7230489e80000, 0, 0, 0},
		{0x98a224000000000, 0x4b3b4ca85a86c47a, 0, 0},
		{0x4a00000000000000, 0xebfdcb54864ada83, 0x28c87cb5c89a2571, 0},
		{0, 0x7775a5f171951000, 0x764b4abe8652979, 0x161bcca7119915b5},
	}
)

func FromBig(b *big.Int) (*Int, error) {
	var z Int
	if overflow := z.SetFromBig(b); overflow {
		return nil, ErrOverflow
	}
	return &z, nil
}

func MustFromBig(b *big.Int) *Int {
	z, err := FromBig(b)
	if err != nil {
		panic(err)
	}
	return z
}

func FromDec(decimal string) (*Int, error) {
	var z Int
	if err := z.SetFromDec(decimal); err != nil {
		return nil, err
	}
	return &z, nil
}

func MustFromDec(decimal string) *Int {
	z, err := FromDec(decimal)
	if err != nil {
		panic(err)
	}
	return z
}

func (z *Int) Dec() string {
	s := z.Sign()
	if s == 0 {
		return "0"
	}
	if z.IsInt64() {
		return strconv.FormatInt(z.Int64(), 10)
	}
	y := new(Int)
	if s > 0 {
		y.Set(z)
	} else {
		y.Neg(z)
	}
	var (
		out     = []byte("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
		divisor = new(Int).SetUint64(10000000000000000000)
		pos     = len(out)
		buf     = make([]byte, 0, 19)
	)

	for {
		var quot Int
		rem := udivrem(quot[:], y[:], divisor)
		y.Set(&quot)
		buf = strconv.AppendUint(buf[:0], rem.Uint64(), 10)
		copy(out[pos-len(buf):], buf)
		if y.IsZero() {
			break
		}
		pos -= 19
	}

	var res string
	if s < 0 {
		res = "-"
	}
	res += string(out[pos-len(buf):])
	return res
}

func (z *Int) SetFromDec(s string) error {
	var isNeg bool
	if len(s) > 0 && s[0] == '-' {
		s = s[1:]
		isNeg = true
	}
	if len(s) > 0 && s[0] == '0' {
		var (
			i int
			c rune
		)
		for i, c = range s {
			if c != '0' {
				break
			}
		}
		s = s[i:]
	}
	if len(s) > len(maxAbsI256Dec) ||
		(len(s) == len(maxAbsI256Dec) && s > maxAbsI256Dec) ||
		(s == maxAbsI256Dec && !isNeg) {
		return ErrOverflow
	}
	if err := z.fromDecimal(s); err != nil {
		return err
	}
	if isNeg {
		z.Neg(z)
	}
	return nil
}

func (z *Int) fromDecimal(bs string) error {
	z.Clear()
	var (
		num       uint64
		err       error
		remaining = len(bs)
	)
	if remaining == 0 {
		return io.EOF
	}
	for i, mult := range multipliers {
		if remaining <= 0 {
			return nil
		}
		if remaining > 19 {
			num, err = strconv.ParseUint(bs[remaining-19:remaining], 10, 64)
		} else {
			num, err = strconv.ParseUint(bs, 10, 64)
		}
		if err != nil {
			return err
		}
		if i == 0 {
			z.SetUint64(num)
		} else {
			base := new(Int).SetUint64(num)
			z.Add(z, base.Mul(base, mult))
		}
		if remaining > 19 {
			bs = bs[0 : remaining-19]
		}
		remaining -= 19
	}
	return nil
}

func (z *Int) SetFromBig(b *big.Int) bool {
	z.Clear()
	words := b.Bits()
	overflow := len(words) > maxWords

	switch maxWords {
	case 4:
		if len(words) > 0 {
			z[0] = uint64(words[0])
			if len(words) > 1 {
				z[1] = uint64(words[1])
				if len(words) > 2 {
					z[2] = uint64(words[2])
					if len(words) > 3 {
						z[3] = uint64(words[3])
					}
				}
			}
		}
	case 8:
		numWords := len(words)
		if overflow {
			numWords = maxWords
		}
		for i := 0; i < numWords; i++ {
			if i&1 == 0 {
				z[i>>1] = uint64(words[i])
			} else {
				z[i>>1] |= uint64(words[i]) << 32
			}
		}
	}

	bSign := b.Sign()
	if !overflow && z.IsNegative() && !(bSign < 0 && z.IsMinI256()) {
		overflow = true
	}
	if bSign == -1 {
		z.Neg(z)
	}
	return overflow
}

func (z *Int) ToBig() *big.Int {
	t := z.Clone()
	isNeg := t.IsNegative()
	if isNeg {
		t.Neg(t)
	}
	b := new(big.Int)
	switch maxWords {
	case 4:
		words := [4]big.Word{big.Word(t[0]), big.Word(t[1]), big.Word(t[2]), big.Word(t[3])}
		b.SetBits(words[:])
	case 8:
		words := [8]big.Word{
			big.Word(t[0]), big.Word(t[0] >> 32),
			big.Word(t[1]), big.Word(t[1] >> 32),
			big.Word(t[2]), big.Word(t[2] >> 32),
			big.Word(t[3]), big.Word(t[3] >> 32),
		}
		b.SetBits(words[:])
	}
	if isNeg {
		b.Neg(b)
	}
	return b
}

func (z *Int) MarshalJSON() ([]byte, error) {
	return []byte(`"` + z.Dec() + `"`), nil
}

func (z *Int) UnmarshalJSON(b []byte) error {
	if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
		return z.UnmarshalText(b)
	}
	return z.SetFromDec(string(b[1 : len(b)-1]))
}

func (z *Int) MarshalText() ([]byte, error) {
	return []byte(z.Dec()), nil
}

func (z *Int) UnmarshalText(input []byte) error {
	return z.SetFromDec(string(input))
}
