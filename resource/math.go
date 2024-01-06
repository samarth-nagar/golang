package resource

func Add(i1 int, i2 int) int {
	return i1 + i2
}

func Multiply(i1 int, i2 int) int {
	return i1 * i2
}
func Yoooo() string {
	return "sup"

}
func sort(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a[0]
}

func matrixaddition(a [][]int, b [][]int) [][]int {
	var c [][]int
	for i := 0; i < len(a); i++ {
		c = append(c, make([]int, len(a[i])))
		for j := 0; j < len(a[i]); j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

// compare yearly values and find the percentage change of 4 years and give a4 percentage change between every year ie 4 output 4 input
func PercentageChange(a []int) []int {
	var b []int
	for i := 0; i < len(a); i++ {
		if i == 0 {
			b = append(b, 0)
		} else {
			b = append(b, (a[i]-a[i-1])/a[i-1]*100)
		}
	}
	return b
}
