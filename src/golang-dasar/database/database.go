package database;

import "fmt";

var connection string

func init() { // jika menggunakan init(), maka function ini akan otomatis dieksekusi meskipun tidak kita panggil
	connection = "Connect Database";
	fmt.Println("This is connection database");
}

func Database() string {
	return connection;
}

/* 
	Package Initialization
	1. Saat kita membuat package, kita bisa membuat function yang akan diakses ketika package kita diakses
	2. Ini sangat coock ketika contohnya, jika package berisi function-funvtion untuk berkomunikasi dengan database, kita membuat function initialisasi untuk membuka koneksi database
	3. Untuk membuat function secara otomatis ketika package diakses, kita cukup membuat function dengan nama init
*/ 