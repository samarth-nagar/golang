package main

import (
	"fmt"
	"go-tut/resource"
)

func main() {
	fmt.Println("hrllo  world")

	ans := resource.Add(1, 2)
	multans := resource.Multiply(1, 2)

	fmt.Println(ans, multans)
}
