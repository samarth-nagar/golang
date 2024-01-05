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
