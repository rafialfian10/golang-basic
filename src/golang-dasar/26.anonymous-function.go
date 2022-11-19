package main;

import "fmt";

type Blacklist func(string)bool;

func registerUser(name string, blacklist Blacklist){
	if blacklist(name) == true { // jika nama pada registerUser dan blacklist sama maka lakukan sesuatu
		fmt.Println("You are blocked", name)
	} else { // jika nama berbeda
		fmt.Println("Wellcome", name)
	}	
}

// biasanya kita membuat function dulu agar nanti dapat dipanggil, kita dapat membuat function anonymous dan menyimpan kedalam variabel.
// func blacklist(name string)bool { 
// 	return name == "admin";
// }

func main(){
	blacklist := func(name string)bool { // kita dapat langsung membuat function tanpa nama yang disimpan kedalam sebuah variabel
		return name == "admin 1";
	}
	registerUser("admin 1", blacklist); // Hasil: You are blocked Admin 1
	registerUser("Rafi", blacklist); // Hasil: Wellcome Rafi
	//---------------------------------------------------------------------

	// Atau langsung membuat function didalam parameter
	registerUser("admin 2", func(name string)bool {
		return name == "admin 2"; // Hasil: You are blocked Admin 2
	});
	registerUser("Rafi", func(name string)bool { // Hasil: Wellcome Rafi
		return name == "admin 2";
	});
	
}

