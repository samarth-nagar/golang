package main

import (
	"fmt"
	"go-tut/resource"
)

func main() {

	fmt.Println("hello  world")
	ans, b, c := resource.MatixAdd(resource.GenRand3X3Matrix(10), resource.GenRand3X3Matrix(10))
	fmt.Println(ans, b, c)
}
