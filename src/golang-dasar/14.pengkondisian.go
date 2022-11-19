package main;

import "fmt";

func main() {
	name := "Budi";

	if name == "Rafi" {
		fmt.Println("Hello", name);
	} else if name == "Ervin" {
		fmt.Println("Hello", name);
	} else if name == "Ali" {
		fmt.Println("Hello", name);
	} else {
		fmt.Println("Hello, kenalan donk");
	}  

	length := "Rafiaaaaaaaaaaaaaaaa";

	if len(length) > 5 && len(length) < 10 {
		fmt.Println("Nama cukup panjang");
	} else if len(length) > 10  {
		fmt.Println("Nama terlalu panjang");
	} else {
		fmt.Println("Nama sudah benar");
	}
}