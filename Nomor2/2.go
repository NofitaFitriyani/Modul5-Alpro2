package main

import "fmt"

func cetakBintang(count int) {
	if count == 0 {
		return
	}
	for i := 0; i < count; i++ {
		fmt.Print("* ")
	}
	fmt.Println()
}

func cetakPola(baris int, barisSekarang int) {
	if barisSekarang > baris {
		return
	}
	cetakBintang(barisSekarang)
	cetakPola(baris, barisSekarang+1)
}

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	cetakPola(n, 1)
}
