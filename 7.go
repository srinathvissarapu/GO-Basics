package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 3 / 5
	y = sum - x
	return
}

func main() {
	
	a, b := split(100)
	c, d := split(50)

	
	fmt.Println(a, b)
	fmt.Println(c, d)
}
