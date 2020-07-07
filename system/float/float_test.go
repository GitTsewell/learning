package float

import (
	"fmt"
	"testing"
)

func TestConvertToBin(t *testing.T) {
	t.Log(ConvertToBin(129))
}

func TestAddBinary(t *testing.T) {
	s := AddBinary("0.11001100110011001100111", "1.10011001100110011001101")
	t.Log(s)
}

func TestFloat32ToBitString(t *testing.T) {
	t.Log(Float32ToBitString(6.125))
}

func TestBitStringToFloat32(t *testing.T) {
	t.Log(BitStringToFloat32("01000000110100000000000000000000"))
}

func TestFloat64ToBitString(t *testing.T) {
	t.Log(Float64ToBitString(0.1))
	t.Log(Float64ToBitString(0.2))
}

func TestBitStringToFloat64(t *testing.T) {
	t.Log(BitStringToFloat64("0011111111010011001100110011001100110011001100110011001100110100"))
}

func TestFloat32(t *testing.T) {
	var f1 float32 = 0.1
	var f2 float32 = 0.2
	fmt.Println(f1 + f2)
}

func TestFloat64(t *testing.T) {
	var f1 float64 = 0.1
	var f2 float64 = 0.2
	fmt.Println(f1 + f2)
}

//0 01111011 10011001100110011001101
//0 01111100 10011001100110011001101

//0.11001100110011001100111
//1.10011001100110011001101
//10.01100110011001100110100
//1.00110011001100110011010
//00111110100110011001100110011010
