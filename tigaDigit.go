package main

import "fmt"

func main() {
	var x, d1, d2, d3, hb int
	fmt.Scanln(&x)
	d1 = x / 100
	hb = x % 100
	d2 = hb / 10
	d3 = hb % 10

	fmt.Println(d1, d2, d3)
}
