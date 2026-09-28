package main

import "fmt"

type Role int

const (
	RoleBuyer Role = iota
	RoleVendor
	RoleAdmin
)

func main() {
	fmt.Println(RoleBuyer, RoleAdmin, RoleVendor)

	x := 1
	fmt.Println(x)
	x = 2
	fmt.Println(x)
}
