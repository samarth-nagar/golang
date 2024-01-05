package resource

import (
	"fmt"
	"math/rand"
)

var secretNum int
var guess int

func random() int {
	secretNum = rand.Intn(100)
	return secretNum
}
func Check() {
	var guess int
	fmt.Printf("guess an number between 1 and 100 :-")
	ans := random()
	for guess != ans {
		fmt.Scanf("%d", &guess)
		if guess > ans {
			fmt.Println("youre wrong the number is smaller")
		}
		if guess < ans {
			fmt.Println("youre wrong the number is bigger")
		}
		if guess == ans {
			fmt.Println("yea youre right the number was", ans)
		}

	}
}
