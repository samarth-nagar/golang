package resource

import (
	"alderaan/resource"
	"fmt"
)

func MatixAdd(a [3][3]int, b [3][3]int) ([3][3]int, [3][3]int, [3][3]int) {
	var c [3][3]int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	fmt.Println(c)
	return c, a, b
}
func GenRand3X3Matrix(integer int) [3][3]int {
	var mat [3][3]int
	for index, value := range mat {
		for index2 := range value {
			mat[index][index2] = resource.GenRandInt(integer)
		}

	}
	return mat

}
func AvgOfElements(array [3]int) {
	var sum int
	for _, value := range array {
		sum += value
	}
	fmt.Println(sum / 3)
}

func Main() {
	var a, b [3][3]int
	a = GenRand3X3Matrix(10)
	b = GenRand3X3Matrix(10)
	_, _, c := MatixAdd(a, b)
	AvgOfElements(c[0])
	AvgOfElements(c[1])
	AvgOfElements(c[2])
}
