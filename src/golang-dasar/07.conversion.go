package main;

import "fmt";

func main(){
	// Konversi int
	var nilai32 int32 = 128;
	var nilai64 int64 = int64(nilai32);
	var nilai8 int8 = int8(nilai32); // hasilnya akan menjadi -128, karena int8 nilai maksimal hanya sampai 127, maka nilainya akan kembali ke awal lagi

	// Konversi String
	var name = "RAFI";
	var r byte = name[0]; // Jika mengkonversi karakter dari string, maka akan menjadi byte, oleh karena itu konversikan kembali menjadi string.
	var rName string = string(r)
	byte dan string boleh tidak ditulis


	fmt.Println(nilai32);
	fmt.Println(nilai64);
	fmt.Println(nilai8);
	fmt.Println(name);
	fmt.Println(rName)
}