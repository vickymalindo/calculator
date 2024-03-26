package main

import (
	"fmt"
	"os"
)

func add(k int, l int) {
	var result = k + l
	fmt.Printf("hasil dari %v + %v = %v \n \n", k, l, result)
	menus()
}

func substract(k int, l int) {
	var result = k - l
	fmt.Printf("hasil dari %v - %v = %v \n \n", k, l, result)
	menus()
}

func divide(k int, l int) {
	var result = k / l
	fmt.Printf("hasil dari %v : %v = %v \n \n", k, l, result)
	menus()
}

func multiply(k int, l int) {
	var result = k * l
	fmt.Printf("hasil dari %v x %v = %v \n \n", k, l, result)
	menus()
}

func menus() {
	var choosen, NumberOne, NumberTwo int

	fmt.Println("Pilihan menu :")
	fmt.Println("1. Tambah")
	fmt.Println("2. Kurang")
	fmt.Println("3. Kali")
	fmt.Println("4. Bagi")
	fmt.Println("5. Keluar")
	fmt.Print("Masukan pilhan anda : ")

	fmt.Scan(&choosen)

	if choosen == 5 {
		os.Exit(1)
	}

	fmt.Print("Masukan angka pertama : ")
	fmt.Scan(&NumberOne)
	fmt.Print("Masukan angka pertama : ")
	fmt.Scan(&NumberTwo)

	if choosen == 1 {
		add(NumberOne, NumberTwo)
	} else if choosen == 2 {
		substract(NumberOne, NumberTwo)
	} else if choosen == 3 {
		multiply(NumberOne, NumberTwo)
	} else {
		divide(NumberOne, NumberTwo)
	}
}

func main() {
	menus()
}
