package float

import (
	"encoding/binary"
	"math"
	"strconv"

	"github.com/imroc/biu"
)

func ConvertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}

// 二进制相加
func AddBinary(a string, b string) string {
	result := ""
	flag := 0 // 存储进位
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		t1, t2 := 0, 0
		if i >= 0 {
			t1 = int(a[i] - '0')
		}
		if j >= 0 {
			t2 = int(b[j] - '0')
		}
		sum := t1 + t2 + flag // 计算当前位置
		switch sum {
		case 3:
			flag = 1
			result = "1" + result
		case 2:
			flag = 1
			result = "0" + result
		case 1:
			flag = 0
			result = "1" + result
		case 0:
			flag = 0
			result = "0" + result
		}
		i--
		j--
	}
	if flag == 1 { // 最终进位
		result = "1" + result
	}
	return result
}

func Float32ToBitString(float float32) string {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	// 得到的切片顺序是反的
	var str string
	for i := len(bytes) - 1; i >= 0; i-- {
		s := biu.ByteToBinaryString(bytes[i])
		// 不足8位 补齐8位  前面加0
		if l := len(s); l < 8 {
			var newS string
			for j := 1; j < 8-l; j++ {
				newS = newS + "0"
			}
			s = newS + s
		}
		str = str + s
	}

	return str
}

func BitStringToFloat32(str string) float32 {
	bytes := biu.BinaryStringToBytes(str)
	// 翻转顺序
	l := len(bytes) - 1
	for i := 0; i < len(bytes)/2-1; i++ {
		bytes[i], bytes[l-i] = bytes[l-i], bytes[i]
	}
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func Float64ToBitString(float float64) string {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	// 得到的切片顺序是反的
	var str string
	for i := len(bytes) - 1; i >= 0; i-- {
		s := biu.ByteToBinaryString(bytes[i])
		// 不足8位 补齐8位  前面加0
		if l := len(s); l < 8 {
			var newS string
			for j := 1; j < 8-l; j++ {
				newS = newS + "0"
			}
			s = newS + s
		}
		str = str + s
	}

	return str
}

func BitStringToFloat64(str string) float64 {
	bytes := biu.BinaryStringToBytes(str)
	// 翻转顺序
	l := len(bytes) - 1
	for i := 0; i < len(bytes)/2-1; i++ {
		bytes[i], bytes[l-i] = bytes[l-i], bytes[i]
	}
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}
