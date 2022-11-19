package main;

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("r[a-z][a-z]i");

	fmt.Println(regex.MatchString("rafi")); // hasilnya true
	fmt.Println(regex.MatchString("ravi")); // hasilnya true
	fmt.Println(regex.MatchString("Xavi")); // hasilnya false
	fmt.Println(regex.MatchString("Rafi")); // hasilnya false


	result := regex.FindAllString("rafi rapi ravi rafa xavi refo", -1); // parameter ada 2. -1 berarti mencari semua
	fmt.Println(result); // hasilnya [rafi rapi ravi]
}

/*
	Package Regexp
	1. Adalah utilitas di golang untuk melakukan pencarian regular expression
	2. Regular expression di golang menggunakan library C yang dibuat oleh google bernama RE2
	3. https://github.com/google/re2/wiki/Syntax
	4. https://golang.org/pkg/regexp/

	Beberapa function di regexp
	function							Kegunaan
	regexp.MustCompile(string)			Membuat regexp
	regexp.MatchString(string) bool		Mengecek apakah regexp match dengan string
	regexp.FindAllString(string, max)	Mencari string yang match dengan maksimum jumlah hasil
*/ 