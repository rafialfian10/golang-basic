package main;

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("     Rafi Alfian     ", " ")); // ini akan menghapus karakter kiri dan kanan
	fmt.Println(strings.ToLower("RAFI ALFIAN"));
	fmt.Println(strings.ToUpper("rafi alfian"));
	fmt.Println(strings.Split("Rafi Alfian Ervin Alfianto Ali Imron", " ")); // nilai yang dikembalikan adalah slice of strings
	fmt.Println(strings.Split("Rafi Alfian Ervin Alfianto Ali Imron", "Alfian")); // maka strings Alfian akan dipotong
	fmt.Println(strings.Contains("Rafi Alfian", "Rafi")); // akan mengambalikan boolean, jika ada string Rafi maka hasilnya true
	fmt.Println(strings.ReplaceAll("Rafi Rafi Rafi", "Rafi", "Alfian")); // maka strings Rafi akan diganti dengan Alfian
	fmt.Println(strings.ReplaceAll("Rafi Ervin Rafi", "Rafi", "Alfian")); // yang akan direplace hanya Rafi saja, Ervin tidak akan direplace
}

/*
	Package Strings
	1. Adalah package yangberisikan function-function untuk memanipulasi tipe data string
	2. Ada banyak sekali function yang bisa digunakan
	3. https://golang.org/pkg/strings

	contohnya
	1. Strings.Trim(string, cutset) berfungsi memotong cutset diawal dan akhir string
	2. Strings.ToLower(string) berfungsi membuat smua karakter string menjadi lower case
	3. Strings.ToUpper(string) berfungsi membuat smua karakter string menjadi upper case
	4. Strings.Split(string, separator) berfungsi memotong string berdasarkan separator
	5. Strings.Contains(string, search) berfungsi mengecek apakah string mengandung string lain
	6. Strings.ReplaceAll(string, from, to) berfungsi mengubah semua string dari from ke to
*/ 