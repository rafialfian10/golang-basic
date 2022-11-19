package main;

import (
	"fmt"
	"strconv" 
)

func main() {
	// strconv.ParseBool
	boolean, err := strconv.ParseBool("true");
	if err == nil {
		fmt.Println("Boolean :", boolean);
	} else {
		fmt.Println("Error :", err.Error());
	}
	//---------------------------------------------

	// strconv.ParseInt
	number, err := strconv.ParseInt("1000000", 10, 64); // parameter ada 3: 1.string, 2.base int(bisa 10, 8, 16), 3.bitSize int (ukuran bisa 32 / 64)
	if err == nil {
		fmt.Println("Int :", number);
	} else {
		fmt.Println("Error :", err.Error());
	}
	//------------------------------------------------

	// strconv.FormatInt
	value := strconv.FormatInt(1000000, 10) // tidak mengembalikan err. 2: binary, 10: desimal, 8: oktal, 16: heksa desimal
	fmt.Println("Format Int :", value);
	//------------------------------------------------

	//strconv.FormatBool
	boolean2 := strconv.FormatBool(true);
	fmt.Println("Format Bool :", boolean2);
	//------------------------------------------------

	// strconv.Itoa
	itoa := strconv.Itoa(5000000) // mengubah int ke string
	fmt.Println("Hasil Itoa :", itoa)
	//------------------------------------------------

	// strconv.Atoi
	atoi, _ := strconv.Atoi("7000000") // mengubah string ke int
	fmt.Println("Hasil Atoi :", atoi)
	//------------------------------------------------

}

/*
	Package strconv
	1. Sebelumnya sudah belajar cara konversi tipe data, misal int32 ke int64
	2. Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda?, misalnya dari int ke string, atau sebaliknya
	3. Hal tersebut bisa kita lakukan dengan bantuan package strconv(string conversion)
	4. https://golang.org/pkg/strconv/

	Beberapa function di package strconv
	Function							Kegunaan
	strconv.parseBool(string)			Mengubah string ke bool
	strconv.parseFloat(string)			Mengubah string ke float
	strconv.parseInt(string)			Mengubah string ke int
	strconv.FotmatFloat(float, ...) 	Mengubah float64 ke string
	strconv.FormatBool(bool) 			Mengubah boolean ke string
	strconv.FormatInt(int, ...)			Mengubah int64 ke string
	strconv.Itoa(i int)					Mengubah int menjadi string
	strconv.Atoi(s string)				Mengubah string menjadi int
*/ 