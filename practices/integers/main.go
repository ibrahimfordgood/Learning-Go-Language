package main

import "fmt"

func main() {
	var stock uint = 2
	stock = stock - 3
	fmt.Println(stock) // 18446744073709551615

	var small int8 = 127
	small++
	fmt.Println(small) // -128: overflow wraps silently too
}
