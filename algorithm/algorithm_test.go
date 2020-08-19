package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	num      = 100000
	rangeNum = 100000
)

// 生成随机切片
func makeRandArr() []int {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	var buf []int
	for i := 0; i < num; i++ {
		buf = append(buf, randSeed.Intn(rangeNum))
	}
	return buf
}

// 冒泡
func TestMaopao(t *testing.T) {
	buf := makeRandArr()
	maopao(buf)
	t.Log(buf)
}

// 选择
func TestXuanze(t *testing.T) {
	buf := makeRandArr()
	ti := time.Now()
	xuanze(buf)
	fmt.Println(buf)
	fmt.Println(time.Since(ti))
}

// 插入
func TestCharu(t *testing.T) {
	buf := makeRandArr()
	ti := time.Now()
	charu(buf)
	fmt.Println(buf)
	fmt.Println(time.Since(ti))
}

// 希尔
func TestXier(t *testing.T) {
	buf := makeRandArr()
	ti := time.Now()
	xier(buf)
	fmt.Println(buf)
	fmt.Println(time.Since(ti))
}

// 快速
func TestKuaisu(t *testing.T) {
	buf := makeRandArr()
	ti := time.Now()
	kuaisu(buf)
	fmt.Println(buf)
	fmt.Println(time.Since(ti))
}

// 归并
func TestGuibing(t *testing.T) {
	buf := makeRandArr()
	ti := time.Now()
	guibing(buf)
	fmt.Println(buf)
	fmt.Println(time.Since(ti))
}

// 堆排
func TestDuipai(t *testing.T) {
	buf := makeRandArr()
	ti := time.Now()
	duipai(buf)
	fmt.Println(buf)
	fmt.Println(time.Since(ti))
}
