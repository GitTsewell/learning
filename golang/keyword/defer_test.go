package keyword

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			fmt.Println("panic again and again")
		}()
		fmt.Println("panic again")
	}()

	fmt.Println("panic once")
}
