package main

import "fmt"

func main() {
	var x, d1, d2, d3, hb int
	// note : hb itu buat nyimpan hasil bagi-nya
	fmt.Scanln(&x)
	d1 = x / 100
	// 546/100 = 5,46 = 5
	hb = x % 100
	// 546 % 100 = 46
	d2 = hb / 10
	// 46/10 = 4.6 = 4
	d3 = hb % 10
	// 46 % 10 = 6

	fmt.Println(d1, d2, d3)
}
