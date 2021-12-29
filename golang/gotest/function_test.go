package gotest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20) // 运行 Fib 函数 N 次
	}
}

// go test -bench BenchmarkRequest(函数名) -run=none
// -benchtime=3s

// input params -- total amount , times
func TestCatsLuck(t *testing.T) {
	totalAmount := 500000
	initAmount := 10000
	amount := initAmount
	times := 2500
	probabilityA := 9524
	probabilityB := 5000
	var clk int
	var runTimes int

	rand.Seed(time.Now().Unix())
	for i := 1; i <= times; i++ {
		if i == times {
			fmt.Println("last amount : ", amount)
		}
		// check amount
		if amount > totalAmount {
			t.Log("余额不足,已终止")
			break
		}
		// update totalAmount
		totalAmount = totalAmount - amount
		// runTimes ++
		runTimes++
		// add clk
		clk = clk + amount
		// select A or B
		if amount == initAmount {
			if rand.Intn(10000) < probabilityA {
				// win
				reward := (amount - amount/50) * 105 / 100
				totalAmount = totalAmount + reward
			}
		} else {
			if rand.Intn(10000) < probabilityB {
				// win
				reward := (amount - amount/50) * 200 / 100
				totalAmount = totalAmount + reward
				amount = initAmount
			} else {
				// lose
				amount = amount * 2
			}
		}
	}

	fmt.Println("total amount : ", totalAmount)
	fmt.Println("run times : ", runTimes)
	fmt.Println("clk : ", clk)
	cats := clk / 50
	fmt.Println("clk eq cats : ", cats)
	fmt.Println("total : ", cats+totalAmount)
}

func TestCLK(t *testing.T) {
	totalAmount := 500000
	amount := 1000
	times := 100000
	probabilityB := 9524
	var clk int
	var runTimes int

	rand.Seed(time.Now().Unix())
	for i := 1; i <= times; i++ {
		// check amount
		if amount > totalAmount {
			t.Log("余额不足,已终止")
			break
		}
		// update totalAmount
		totalAmount = totalAmount - amount
		// runTimes ++
		runTimes++
		// add clk
		clk = clk + amount

		if rand.Intn(10000) < probabilityB {
			// win
			reward := (amount - amount/50) * 105 / 100
			totalAmount = totalAmount + reward
		}
	}

	fmt.Println("total amount : ", totalAmount)
	fmt.Println("run times : ", runTimes)
	fmt.Println("clk : ", clk)
	cats := clk / 50
	fmt.Println("clk eq cats : ", cats)
	fmt.Println("total : ", cats+totalAmount)
}
