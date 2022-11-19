package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.4));
	fmt.Println(math.Floor(1.8));
	fmt.Println(math.Ceil(1.2));
	fmt.Println(math.Max(10, 20)); 
	fmt.Println(math.Min(10, 20));
}

/*
	Package Math
	1. merupakan package yang berisikan constant dan fungsi matematika
	2. https://golang.org/pkg/math/

	Contoh beberapa math
	math.Round(float64)			Membulatkan dengan yang paling dekat
	math.Floor(float64)			Membulatkan ke bawah
	math.Ceil(float64)			Membulatkan ke atas
	math.Max(float64, float64)	Mengembalikan nilai float64 yang paling besar
	math.Min(float64, float64)	Mengembalikan nilai float64 yang paling kecil
*/ 