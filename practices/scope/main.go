package main

import (
	"errors"
	"fmt"
)

func find(id int) (string, error) {
	if id == 0 {
		return "", errors.New("not found")
	}
	return "Aminata", nil
}

func main() {
	var err error
	name := "nobody"
	if true {
		name, err := find(0) // BUG: := declares NEW name and err inside this block
		_ = name
		_ = err
	}

	fmt.Println(name, err) // prints "nobody <nil>": the error vanished
}
