package main

import "fmt"

func tampilkanGanjil(n, i int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	tampilkanGanjil(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)
	tampilkanGanjil(n, 1)
}
