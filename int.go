package betterbig

import (
	"math/big"
	"math/rand"
)

type Int struct {
	native *big.Int
}

func NewIntFromBigInt(native *big.Int) Int {
	return Int{native: native}
}

func NewIntFromInt64(n int64) Int {
	return Int{native: big.NewInt(n)}
}

func NewIntFromUint64(n uint64) Int {
	return Int{native: new(big.Int).SetUint64(n)}
}

func NewIntFromString(s string, base int) (Int, bool) {
	native, ok := new(big.Int).SetString(s, base)
	return Int{native: native}, ok
}

func NewIntFromBytes(buf []byte) Int {
	return Int{native: new(big.Int).SetBytes(buf)}
}

func NewIntFromInt(n int) Int {
	return Int{native: big.NewInt(int64(n))}
}

func NewIntFromBigWord(n []big.Word) Int {
	return Int{native: new(big.Int).SetBits(n)}
}

func NewIntFromMulRange(a, b int64) Int {
	return Int{native: new(big.Int).MulRange(a, b)}
}

func NewIntFromBinomial(n, k int64) Int {
	return Int{native: new(big.Int).Binomial(n, k)}
}

func RandInt(rnd *rand.Rand, n Int) Int {
	return Int{native: new(big.Int).Rand(rnd, n.native)}
}

func (i Int) BigInt() *big.Int {
	return i.native
}

func (i Int) Int64() int64 {
	return i.native.Int64()
}

func (i Int) Uint64() uint64 {
	return i.native.Uint64()
}

func (i Int) String() string {
	return i.native.String()
}

func (i Int) Bytes() []byte {
	return i.native.Bytes()
}

func (i Int) Bits() []big.Word {
	return i.native.Bits()
}

func (i Int) Float64() (float64, big.Accuracy) {
	return i.native.Float64()
}

func (i Int) IsInt64() bool {
	return i.native.IsInt64()
}

func (i Int) IsUint64() bool {
	return i.native.IsUint64()
}

func (i Int) Sign() int {
	return i.native.Sign()
}

func (i Int) Abs() Int {
	return Int{native: new(big.Int).Abs(i.native)}
}

func (i Int) Neg() Int {
	return Int{native: new(big.Int).Neg(i.native)}
}

func (i Int) Add(j Int) Int {
	return Int{native: new(big.Int).Add(i.native, j.native)}
}

func (i Int) Sub(j Int) Int {
	return Int{native: new(big.Int).Sub(i.native, j.native)}
}

func (i Int) Mul(j Int) Int {
	return Int{native: new(big.Int).Mul(i.native, j.native)}
}

func (i Int) Quo(j Int) Int {
	return Int{native: new(big.Int).Quo(i.native, j.native)}
}

func (i Int) Rem(j Int) Int {
	return Int{native: new(big.Int).Rem(i.native, j.native)}
}

func (i Int) QuoRem(j, k Int) (Int, Int) {
	q, r := new(big.Int).QuoRem(i.native, j.native, k.native)
	return Int{native: q}, Int{native: r}
}

func (i Int) Div(j Int) Int {
	return Int{native: new(big.Int).Div(i.native, j.native)}
}

func (i Int) Mod(j Int) Int {
	return Int{native: new(big.Int).Mod(i.native, j.native)}
}

func (i Int) DivMod(j, k Int) (Int, Int) {
	d, m := new(big.Int).DivMod(i.native, j.native, k.native)
	return Int{native: d}, Int{native: m}
}

func (i Int) BitLen() int {
	return i.native.BitLen()
}

func (i Int) TrailingZeroBits() uint {
	return i.native.TrailingZeroBits()
}

// Cmp compares i and j and returns:
//
//	-1 if i <  j
//	 0 if i == j
//	+1 if i >  j
func (i Int) Cmp(j Int) int {
	return i.native.Cmp(j.native)
}

// CmpAbs compares the absolute values of i and j and returns:
//
//	-1 if |i| <  |j|
//	 0 if |i| == |j|
//	+1 if |i| >  |j|
func (i Int) CmpAbs(j Int) int {
	return i.native.CmpAbs(j.native)
}

func (i Int) ModInverse(j Int) Int {
	return Int{native: new(big.Int).ModInverse(i.native, j.native)}
}

func (i Int) ModSqrt(j Int) Int {
	return Int{native: new(big.Int).ModSqrt(i.native, j.native)}
}

func (i Int) Lsh(n uint) Int {
	return Int{native: new(big.Int).Lsh(i.native, n)}
}

func (i Int) Rsh(n uint) Int {
	return Int{native: new(big.Int).Rsh(i.native, n)}
}

func (i Int) Bit(n int) uint {
	return i.native.Bit(n)
}

func (i Int) SetBit(n int, b uint) Int {
	return Int{native: new(big.Int).SetBit(i.native, n, b)}
}

func (i Int) And(j Int) Int {
	return Int{native: new(big.Int).And(i.native, j.native)}
}

func (i Int) AndNot(j Int) Int {
	return Int{native: new(big.Int).AndNot(i.native, j.native)}
}

func (i Int) Or(j Int) Int {
	return Int{native: new(big.Int).Or(i.native, j.native)}
}

func (i Int) Xor(j Int) Int {
	return Int{native: new(big.Int).Xor(i.native, j.native)}
}

func (i Int) Not() Int {
	return Int{native: new(big.Int).Not(i.native)}
}

func (i Int) Sqrt() Int {
	return Int{native: new(big.Int).Sqrt(i.native)}
}

func (i Int) Exp(exp Int, m *Int) Int {
	if m == nil {
		return Int{native: new(big.Int).Exp(i.native, exp.native, nil)}
	}
	return Int{native: new(big.Int).Exp(i.native, exp.native, m.native)}
}

// TODO: FillBytes, GCD
