package main;

import (
	"fmt"
	"flag"
)

func main() {
	// flag.String() itu berarti akan diparsing menjadi string
	var host *string = flag.String("host", "localhost", "Put your database host"); //value 1: konfig, value 2: nilai default
	var username *string = flag.String("username", "root", "Put your database username");
	var password *string = flag.String("password", "root", "Put your database password");
	var number *int = flag.Int("number", 1996, "Put your database number");

	flag.Parse(); // Sebelum menggunakan data diatas, maka lakukan flag.Parse() untuk memproses parsing

	fmt.Println("Host :", *host); // Karena hasil dari parsing adalah pointer, maka panggil dengan menggunakan pointer
	fmt.Println("username :", *username);
	fmt.Println("Password :", *password);
	fmt.Println("Number :", *number);

	// jika kita eksekusi di terminal go run 41.package-flag.go -host=localhost3000 -username=rafi -password=12345678, maka nilai default akan ditimpa
}

/*
	Package Flag
	1. Package fag berisikan fungsionalitas untuk memparsing command line argument
	2. https://golang.org/pkg/flag/
*/ 
