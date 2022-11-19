package main;

import (
	"fmt"
	"os" // digunakan untuk menangkap argumen dibelakang layar, jika ingin mengambil argumen, maka mulai dari index 1, karena index 0 adalah lokasi file binari-nya
)

func main() {
	// Package OS
	args := os.Args;
	fmt.Println("Ini adalah argumen :", args);
	// misal jika running diterminal 40.package-os.go Rafi Alfian
	fmt.Println(args[1]) // ini akan menangkap Rafi
	fmt.Println(args[2]) // ini akan menagkap Alfian
	//-----------------------------------------


	// Package Hostname
	hostname, err := os.Hostname() // hostname mengembalikan 2 value os.Hostname(name string, err error)
	if err == nil {
		fmt.Println("Hostname :", hostname);
	} else {
		fmt.Println("Error :", err.Error());
	}
	//-------------------------------------------


	// Package Get Environment
	username := os.Getenv("USERNAME");
	password := os.Getenv("PASSWORD");
	// jika kita export di terminal
	// export USERNAME=root
	// export PASSWORD=root
	// go run 40.package-os.go Rafi Alfian
	// maka hasilnya root dan root

	fmt.Println(username);
	fmt.Println(password)

}


/*
	Package OS
	1. Golang telah menyediakan banyak ekali package bawaan, salah satunya adalah package os
	2. Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan disemua sistem operasi)
	3. https://golang.org/pkg/os/
*/

/*
	Package Hostname
	1. Berfungsi untuk mengambil nama sistem operasi yang kita pakai
	2. Hostname mengembalikan 2 value (name string, err error)
*/ 

/*
	Package Get Environment
	1. Digunakan untuk setting konfigurasi pada aplikasi
*/ 