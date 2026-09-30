package main

import "fmt"

func main() {

	fmt.Println(0.1 + 0.2)
	fmt.Println(0.1+0.2 == 0.3)

	var fareCents int64 = 3 * 300 // 3 miles at $ 3.00 mile
	fmt.Printf("$%d.%02d\n", fareCents/100, fareCents%100)

}
