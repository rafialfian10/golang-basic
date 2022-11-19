package main

import (
	"fmt"
	"golang-dasar/helper"
	_ "golang-dasar/database" // ketika memanggil package ini, maka wajib untuk mengeksekusi semua function yang ada dipackage, tapi jika hanya ingin mengeksekusi salah satunya maka gunakan blank identifier
)

func main() {
	result := helper.SayHello("Rafi Alfian");
	fmt.Println(result);

	// resultDatabase := database.Database();
	// fmt.Println(resultDatabase);
}

/*
	Blank Identifier
	1. Kadang kita ingin menjalankan init function dipackage tanpa harus mengeksekusi salah satu function yang ada di package
	2. Secara default, Go-lang akan komplen ketika ada package yang diimport namun tidak digunakan
	3. Untuk menangani hal tersebut, kita bisa menggunakan blank identifier( _ ) sebelum nama package ketika melakukan import
*/ 