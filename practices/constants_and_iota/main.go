package main

import "fmt"

const PlatformFeePercent = 5

type Status int

const (
	statusPending Status = iota
	statusPaid
	statusShipped
	statusDelievered
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println(statusPaid, statusDelievered, KB, MB, GB)
}
