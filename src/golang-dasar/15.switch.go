package main;

import "fmt";

func main(){
	name := "Rafi";

	switch name {
		case "Rafi": {
			fmt.Println("Hello", name);
			fmt.Println("Hello", name);
		}
		case "Ervin": {
			fmt.Println("Hello", name);
		}
		case "Ali": {
			fmt.Println("Hello", name);
		}
		default: {
			fmt.Println("Nama yang anda masukkan salah!");
		}
	}


	// Switch dengan short statement
	name2 := "Chandra";

	switch length := len(name2); length > 5{ // buat variabel length (nama bebas) yang diisi dengan name2, lalu lakukan kondisi
		case true: {
			fmt.Println("Nama terlalu panjang");
		}
		case false: {
			fmt.Println("Nama sudah benar");
		}
	}



	// Switch tanpa kondisi
	name3 := "Dawam Mahfudz";

	length3 := len(name3);

	switch {
		case length3 > 15: {
			fmt.Println("Nama terlalu panjang");
		} 
		case length3 > 10: {
			fmt.Println("Nama lumayan panjang");
		}
		default: {
			fmt.Println("Nama sudah benar");
		}
	}

}