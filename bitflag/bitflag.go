package main

import (
	"fmt"
	"os"
	"strings"
)

type bitFlag int

const (
	active bitFlag = 1 << iota
	send
	receive
)

func main() {

	var f bitFlag
	a := os.Args
	l := len(a)

	if l > 1 {
		digit := a[1][0] - 0
		f = bitFlag(digit)
	} else {
		os.Exit(0)
	}

	/*flag := active | send | receive
	fmt.Println(flag)
	fmt.Println(active)
	fmt.Println(send)
	fmt.Println(receive)*/

	fmt.Println(f.String())

}
func (flag bitFlag) String() string {

	var flags []string

	if flag&active == active {
		flags = append(flags, "Active")
	}

	if flag&send == send {
		flags = append(flags, "Send")
	}

	if flag&receive == receive {
		flags = append(flags, "Receive")
	}

	if len(flags) > 0 {
		// приведение типа int(flag) совершенно необходимо,
		// чтобы избежать бесконечной рекурсии!
		return fmt.Sprintf("%d (%s)", int(flag), strings.Join(flags, " | "))
	}

	return "0()"

}
