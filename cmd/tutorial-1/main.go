package main

import (
	"fmt"
	"go-tut/resource"
)

func main() {
	fmt.Println("hrllo  world")
	resource.MatixAdd(resource.GenRand3X3Matrix(9), resource.GenRand3X3Matrix(9))

}
