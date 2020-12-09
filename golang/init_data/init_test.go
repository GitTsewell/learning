package main

import (
	"fmt"
	"testing"
)

func TestNeWData(t *testing.T) {
	type user struct {
		Name  string
		Class int
	}

	u1 := new(user)
	u1.Name = "xiao ming"
	u1.Class = 1

	u2 := user{
		Name:  "xiao fu",
		Class: 2,
	}

	u3 := &user{
		Name:  "pang hu",
		Class: 2,
	}

	c1 := u1
	c1.Name = "tutu"

	c2 := u2
	c2.Name = "xiao hong"

	c3 := u3
	c3.Class = 5
	fmt.Println(u1, u2, u3)
	fmt.Println(c1, c2, c3)
}
