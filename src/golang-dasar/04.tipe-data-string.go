package main;

import "fmt";

func main() {
	fmt.Println("Rafi Alfian");
	fmt.Println("Pendidikan Teknik Informatika");
	fmt.Println("Universitas Muhammadiyah Surakarta");

	fmt.Println(len("Rafi Alfian")); // Hasil 11
	fmt.Println("Pendidikan Teknik Informatika"[0]); // Hasil 80 (80 adalah hasil dari konversi byte. Byte-nya P)
	fmt.Println("Universitas Muhammadiyah Surakarta"[1]); // Hasil 110 (110 adalah hasil dari konversi byte. Byte-nya n)
}