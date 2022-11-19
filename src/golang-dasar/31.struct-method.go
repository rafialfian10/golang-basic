package main;

import "fmt";

type BiodataMahasiswa struct {
	Nama string
	Email string
	Nim int
	Prodi string
}

// Cara yang biasa untuk memanggil struct (1)
// func sayHello(biodata BiodataMahasiswa, name string) { 
// 	fmt.Println("Hello", name, "my name is", biodata.Nama);
// }

// Cara memanggil struct dengan struct function / struct method (2)
func (biodata1 BiodataMahasiswa) sayHello(name string){
	fmt.Println("Hello", name, "my name is", biodata1.Nama)
}

func (biodata2 BiodataMahasiswa) sayHi(name string){
	fmt.Println("Hello", name, "my name is", biodata2.Nama)
}

func main() {
	rafi := BiodataMahasiswa {
		Nama: "Rafi Alfian",
		Email: "a710170021@student.ums.ac.id",
		Nim: 21,
		Prodi: "Teknik Informatika",
	}

	// Cara memanggil function cara 1
	// sayHello(rafi, "Ervin") // parameter pertama adalah struct 
	// Jika menggunakan cara 2 maka cara memanggilnya tidak bisa seperti dibawah


	// Cara memanggil function cara 2
	rafi.sayHello("Ervin"); // Ini seakan-akan struct mempunyai sebuah function
	rafi.sayHi("Ali");
}

/*
	Struct Method
	1. Struct bisanya tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
	2. Namun jika kita ingin menambahkan method kedalam struct, sehingga seakan-akan struct memiliki function
	3. Method adalah function
*/ 

