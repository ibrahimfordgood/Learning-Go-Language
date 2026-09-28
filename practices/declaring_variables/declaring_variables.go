package main

import "fmt"

var appName = "Brian"

func main() {
	var fare int
	var rate float64 = 3.00
	var city = "Wilmington"
	passengers := 4
	pickup, dropoff := "New Castle", "PHL"

	fare = passengers * 10
	fmt.Println(appName, fare, rate, city, pickup, dropoff)

}
