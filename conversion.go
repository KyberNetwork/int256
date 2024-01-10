package int256

import (
	"errors"
	"io"
	"math/big"
	"math/bits"
	"strconv"
)

const (
	minI256Dec = "-57896044618658097711785492504343953926634992332820282019728792003956564819968"
	maxI256Dec = "57896044618658097711785492504343953926634992332820282019728792003956564819967"
	maxWords   = 256 / bits.UintSize
)

var (
	ErrI256Range = errors.New("int256: out of range")
	ErrOverflow  = errors.New("int256: overflow")

	multipliers = [5]*Int{
		nil,
		{10000000000000000000, 0, 0, 0},
		{687399551400673280, 5421010862427522170, 0, 0},
		{5332261958806667264, 17004971331911604867, 2938735877055718769, 0},
		{0, 8607968719199866880, 532749306367912313, 1593091911132452277},
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
	if z.IsMinI256() {
		return minI256Dec
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
	if s == minI256Dec {
		z.Set(MinI256)
		return nil
	}

	var isNeg bool
	if len(s) > 0 && s[0] == '-' {
		s = s[1:]
		isNeg = true
	}
	if len(s) > 0 && s[0] == '0' {
		var i int
		var c rune
		for i, c = range s {
			if c != '0' {
				break
			}
		}
		s = s[i:]
	}

	if len(s) > len(maxI256Dec) || (len(s) == len(maxI256Dec) && s > maxI256Dec) {
		return ErrI256Range
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
		} else if remaining > 19 {
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
			if i%2 == 0 {
				z[i/2] = uint64(words[i])
			} else {
				z[i/2] |= uint64(words[i]) << 32
			}
		}
	}
	if b.Sign() == -1 {
		z.Neg(z)
	}
	return overflow
}

func (z *Int) ToBig() *big.Int {
	isNeg := z.IsNegative()
	if isNeg {
		z.Neg(z)
	}
	b := new(big.Int)
	switch maxWords {
	case 4:
		words := [4]big.Word{big.Word(z[0]), big.Word(z[1]), big.Word(z[2]), big.Word(z[3])}
		b.SetBits(words[:])
	case 8:
		words := [8]big.Word{
			big.Word(z[0]), big.Word(z[0] >> 32),
			big.Word(z[1]), big.Word(z[1] >> 32),
			big.Word(z[2]), big.Word(z[2] >> 32),
			big.Word(z[3]), big.Word(z[3] >> 32),
		}
		b.SetBits(words[:])
	}
	if isNeg {
		b.Neg(b)
	}
	return b
}
