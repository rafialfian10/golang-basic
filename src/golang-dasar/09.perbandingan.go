package main;

import "fmt";

func main(){
	var name1 = "Rafi";
	var name2 = "Rafi";
	var result bool = name1 == name2;

	var angka1 = 100;
	var angka2 = 200;

	fmt.Println(result);
	fmt.Println(angka1 > angka2);
	fmt.Println(angka1 < angka2);
	fmt.Println(angka1 == angka2);
	fmt.Println(angka1!= angka2);
}

// Operator Perbandingan (<, >, <=, >=, ==, !=)