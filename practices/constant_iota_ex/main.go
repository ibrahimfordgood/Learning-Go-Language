package main

import "fmt"

type Status string

const (
	statusPending   Status = "Pending"
	statusPaid      Status = "Paid"
	statusConfirmed Status = "Confirmed"
)

func main() {
	fmt.Println(statusConfirmed, statusPaid, statusPending)
}
