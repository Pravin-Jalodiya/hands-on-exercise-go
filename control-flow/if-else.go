package main

import (
	"fmt"
	"math/rand"
)

func ifElse() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	if x < y {
		fmt.Println("x<y")
	} else if x > y {
		fmt.Println("x>y")
	} else {
		fmt.Println("x=y")
	}
}
