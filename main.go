package main

import "fmt"

func main() {
	var x, y, hasil float64
	fmt.Scanln(&x, &y)
	hasil = 1/(3*(x*x)+10) + 10*y + 7
	fmt.Println(hasil)
}

