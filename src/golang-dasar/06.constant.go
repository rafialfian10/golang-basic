package main;

import "fmt";

func main() {

	const name string = "Rafi Alfian";
	const email = "A710170021";
	const nim = 21;

	// name = "Alfian Rafi" // nama tidak bisa ditimpa karena tipe data const
	// nim = 22;

	const (
		firstName = "Rafi";
		lastName = "Alfian";
	)

	fmt.Println(name);
	fmt.Println(email);
	fmt.Println(nim);
	fmt.Println(firstName);
	fmt.Println(lastName);
}