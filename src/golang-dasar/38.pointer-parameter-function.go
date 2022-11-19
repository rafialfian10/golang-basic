package main;

import "fmt";

// type Address struct {
// 	City, Province, Country string
// }

// func ChangeCountryToIndonesia(address *Address) { // jika (address Address) tidak akan berubah. Jika ingin merubah data maka tambahkan pointer (address *Address) agar data dapat diubah
// 	address.Country = "Indonesia";
// 	address.City = "Kudus";
// }

// func main() {
// 	var address1 Address = Address{"Jepara", "Jawa Tengah", ""};
// 	// ChangeAddressToIndonesia(address1); // data tidak akan berubah karena address1 bukan pointer, agar data berubah, maka ubah address1 menjadi pointer dengan menambahkan &
// 	ChangeCountryToIndonesia(&address1);

// 	fmt.Println(address1);
// }

/*
	Poniter di Function
	1. Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan copy, lalu dikirm ke function tersebut
	2. Oleh karena itu, jika kita mengubah data didalam function, data yang aslinya tidak akan pernah berubah.
	3. Hal ini akan membuat variabel mejadi akan, karena tidak akan bisa berubah.
	4. Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
	5. Untuk melakukan ini, kita juga bisa menggunakan pointer di function
	6. Untuk menjadikan sebuah parameter sebagai function, kita bisa menggunakan operator * di parameternya
	*/ 

	type Address struct {
		City, Province, Country string
	}

	func ChangeAddress(address *Address) {
		address.Country = "Indonesia"
	}

	func ChangeAddressCountry(address *Address) {
		address.Country = "Japan"
	}

	

	func main() {
		var address1 Address = Address{"Jepara", "Jawa Tengah", ""};
		var address2 *Address = &address1

		address2.City = "Kudus";
		ChangeAddress(&address1)

		var address3 Address = Address{"Bandung", "Jawa barat", "Indonesia"};
		var address4 *Address = &address3;
		// *address4 = Address{"Malang", "Jawa Timur", "Indonesia"};

		ChangeAddressCountry(address4)

		fmt.Println(address1);
		fmt.Println(address2);
		fmt.Println(address3);
		fmt.Println(address4);
	}