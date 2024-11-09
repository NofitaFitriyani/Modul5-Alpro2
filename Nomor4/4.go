package main

import "fmt"

func tampilkanBilangan(n, i int, turun bool) {
	if turun {
		if i < 1 {
			tampilkanBilangan(n, 2, false)
			return
		}
		fmt.Print(i, " ")
		tampilkanBilangan(n, i-1, true)
	} else {
		if i > n {
			return
		}
		fmt.Print(i, " ")
		tampilkanBilangan(n, i+1, false)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)
	tampilkanBilangan(n, n, true)
}
