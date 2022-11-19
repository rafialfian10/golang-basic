package main;

import "fmt";

func main(){
	var name string; // Jika telah dideklarasikan string, maka variabel hanya dapat diisi dengan string
	name = "Rafi Alfian";
	fmt.Println(name);

	name = "Alfian Rafi" // variabel dapat ditimpa, tapi tipe data tetap string, jika misal diganti interger, maka akan error
	fmt.Println(name)

	var email = "rafialfian770@gmail.com"; // Atau dapat langsung mengisi variabel dengan data. Golang sudah dapat mengetahui tipe data tersebt.
	fmt.Println(email)

	age := 25; // Atau dapat juga tanpa menuliskan variable(var), tapi cukup dengan menambahkan :=
	fmt.Println(age);
	// PENTING(jika tipe data interger, lalu tidak menuliskan int berapa, maka akan secara default menjadi interger sesuai dengan SO komputer, ole karena itu, tipe data interger dapat dipaksakan menjadi int8, int32, int64, dll). Misal seperti dibawah

	var nim int8 = 21
	fmt.Println(nim)

	var ( // Multiple variabel, dengan membungkusnya dengan kurung buka ().
		firstname = "Rafi";
		lastName = "Alfian";
	)

	fmt.Println(firstname);
	fmt.Println(lastName);
}