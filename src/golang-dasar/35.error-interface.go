package main;

import (
	"fmt";
	"errors";
)
// type error interface { // cara memanggil error secara manual
// 	Error() string
// }

func pembagian(nilai int, pembagi int) (int, error) { // return parameter kedua(pembagi) adalah interface error
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0");
	} else {
		result := nilai / pembagi;
		return result, nil; // karena error adalah interface maka bisa gunakan nil 
	}

}

func main() {
	hasil, err := pembagian(100, 10);
	if err == nil { // tidak error
		fmt.Println("Hasil :", hasil)
	} else {
		fmt.Println(err.Error())
	}
}

/*
	Error Interface
	1. go-lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface-nya adalah error

	Membuat Error
	1. Untuk membuat error, tidak perlu manual
	2. Go-lang telah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors
*/ 

