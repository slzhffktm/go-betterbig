package betterbig_test

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/slzhffktm/go-betterbig"
	"github.com/stretchr/testify/assert"
)

func TestNewIntFromBigInt(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromBigInt(native)
	assert.Equal(t, native, i.BigInt())
}

func TestNewIntFromInt64(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, native, i.BigInt())
}

func TestNewIntFromUint64(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromUint64(10)
	assert.Equal(t, native, i.BigInt())
}

func TestNewIntFromString(t *testing.T) {
	native := big.NewInt(10)
	i, ok := betterbig.NewIntFromString("10", 10)
	assert.True(t, ok)
	assert.Equal(t, native, i.BigInt())
}

func TestNewIntFromBytes(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromBytes([]byte{10})
	assert.Equal(t, native, i.BigInt())
}

func TestNewIntFromInt(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt(10)
	assert.Equal(t, native, i.BigInt())
}

func TestNewIntFromBigWord(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromBigWord([]big.Word{10})
	assert.Equal(t, native, i.BigInt())
}

func TestNewIntFromMulRange(t *testing.T) {
	native := new(big.Int).MulRange(1, 10)
	i := betterbig.NewIntFromMulRange(1, 10)
	assert.Equal(t, native, i.BigInt())
}

func TestNewIntFromBinomial(t *testing.T) {
	native := new(big.Int).Binomial(10, 5)
	i := betterbig.NewIntFromBinomial(10, 5)
	assert.Equal(t, native, i.BigInt())
}

func TestRandInt(t *testing.T) {
	i := betterbig.RandInt(rand.New(rand.NewSource(123)), betterbig.NewIntFromInt64(10))
	assert.NotZero(t, i.BigInt())
}

func TestInt_Int64(t *testing.T) {
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, int64(10), i.Int64())
}

func TestInt_BigInt(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, native, i.BigInt())
}

func TestInt_Uint64(t *testing.T) {
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, uint64(10), i.Uint64())
}

func TestInt_String(t *testing.T) {
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, "10", i.String())
}

func TestInt_Bytes(t *testing.T) {
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, []byte{10}, i.Bytes())
}

func TestInt_Bits(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, native.Bits(), i.Bits())
}

func TestInt_Float64(t *testing.T) {
	native := big.NewInt(10)
	floatNative, accuracyNative := native.Float64()

	i := betterbig.NewIntFromInt64(10)
	float, accuracy := i.Float64()

	assert.Equal(t, floatNative, float)
	assert.Equal(t, accuracyNative, accuracy)
}

func TestInt_IsInt64(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)

	assert.Equal(t, native.IsInt64(), i.IsInt64())
}

func TestInt_IsUint64(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)

	assert.Equal(t, native.IsUint64(), i.IsUint64())
}

func TestInt_Sign(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, native.Sign(), i.Sign())
}

func TestInt_Abs(t *testing.T) {
	native := big.NewInt(-10)
	i := betterbig.NewIntFromInt64(-10)
	assert.Equal(t, new(big.Int).Abs(native), i.Abs().BigInt())
}

func TestInt_Neg(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Neg(native), i.Neg().BigInt())
}

func TestInt_Add(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Add(native, native), i.Add(i).BigInt())
}

func TestInt_Sub(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Sub(native, native), i.Sub(i).BigInt())
}

func TestInt_Mul(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Mul(native, native), i.Mul(i).BigInt())
}

func TestInt_Quo(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Quo(native, native), i.Quo(i).BigInt())
}

func TestInt_Rem(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Rem(native, native), i.Rem(i).BigInt())
}

func TestInt_QuoRem(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	q, r := new(big.Int).QuoRem(native, native, native)
	qi, ri := i.QuoRem(i, i)
	assert.Equal(t, q, qi.BigInt())
	assert.Equal(t, r, ri.BigInt())
}

func TestInt_Div(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Div(native, native), i.Div(i).BigInt())
}

func TestInt_Mod(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Mod(native, native), i.Mod(i).BigInt())
}

func TestInt_DivMod(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	d, m := new(big.Int).DivMod(native, native, native)
	di, mi := i.DivMod(i, i)
	assert.Equal(t, d, di.BigInt())
	assert.Equal(t, m, mi.BigInt())
}

func TestInt_BitLen(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, native.BitLen(), i.BitLen())
}

func TestInt_TrailingZeroBits(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, native.TrailingZeroBits(), i.TrailingZeroBits())
}

func TestInt_Cmp(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, native.Cmp(native), i.Cmp(i))

	native2 := big.NewInt(20)
	i2 := betterbig.NewIntFromInt64(20)
	assert.Equal(t, native.Cmp(native2), i.Cmp(i2))

	native3 := big.NewInt(0)
	i3 := betterbig.NewIntFromInt64(0)
	assert.Equal(t, native.Cmp(native3), i.Cmp(i3))
}

func TestInt_CmpAbs(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, native.CmpAbs(native), i.CmpAbs(i))

	native2 := big.NewInt(-10)
	i2 := betterbig.NewIntFromInt64(-10)
	assert.Equal(t, native.CmpAbs(native2), i.CmpAbs(i2))

	native3 := big.NewInt(0)
	i3 := betterbig.NewIntFromInt64(0)
	assert.Equal(t, native.CmpAbs(native3), i.CmpAbs(i3))
}

func TestInt_ModInverse(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).ModInverse(native, native), i.ModInverse(i).BigInt())
}

func TestInt_ModSqrt(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).ModSqrt(native, big.NewInt(3)), i.ModSqrt(betterbig.NewIntFromInt64(3)).BigInt())
}

func TestInt_Lsh(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Lsh(native, 10), i.Lsh(10).BigInt())
}

func TestInt_Rsh(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Rsh(native, 10), i.Rsh(10).BigInt())
}

func TestInt_Bit(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, native.Bit(10), i.Bit(10))
}

func TestInt_SetBit(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).SetBit(native, 10, 1), i.SetBit(10, 1).BigInt())
}

func TestInt_And(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).And(native, native), i.And(i).BigInt())
}

func TestInt_AndNot(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).AndNot(native, native), i.AndNot(i).BigInt())
}

func TestInt_Or(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Or(native, native), i.Or(i).BigInt())
}

func TestInt_Xor(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Xor(native, native), i.Xor(i).BigInt())
}

func TestInt_Sqrt(t *testing.T) {
	native := big.NewInt(10)
	i := betterbig.NewIntFromInt64(10)
	assert.Equal(t, new(big.Int).Sqrt(native), i.Sqrt().BigInt())
}

func TestInt_MultipleOperations(t *testing.T) {
	i := betterbig.NewIntFromInt64(10)
	i = i.Add(i).Mul(i).Quo(i).Rem(betterbig.NewIntFromInt64(12)).Div(i).Mod(i).Lsh(4).
		Rsh(2).And(i).AndNot(i).Or(i).Xor(i).Sqrt().Neg().Abs().ModInverse(i)

	native := big.NewInt(10)
	native = new(big.Int).Add(native, native)
	native = new(big.Int).Mul(native, native)
	native = new(big.Int).Quo(native, native)
	native = new(big.Int).Rem(native, big.NewInt(12))
	native = new(big.Int).Div(native, native)
	native = new(big.Int).Mod(native, native)
	native = new(big.Int).Lsh(native, 4)
	native = new(big.Int).Rsh(native, 2)
	native = new(big.Int).And(native, native)
	native = new(big.Int).AndNot(native, native)
	native = new(big.Int).Or(native, native)
	native = new(big.Int).Xor(native, native)
	native = new(big.Int).Sqrt(native)
	native = new(big.Int).Neg(native)
	native = new(big.Int).Abs(native)
	native = new(big.Int).ModInverse(native, native)

	assert.Equal(t, native, i.BigInt())
}
