package main;

import "fmt";

// Parameter hanya 1
type Filter func(string)string // function type declarations. Ini bisa kita panggil di dalam parameter di dalam function

func sayHelloWithFilter(name string, filter Filter){
	nameFiltered :=	filter(name);
	fmt.Println("Hello,", nameFiltered)
}

func spamFilter(name string)string{
	if name == "Anjing"{
		return "..."
	} else {
		return name;
	}
}

// Parameter lebih dari 1
type Filter2 func(string, int, bool) (string, int, bool);

func sayHelloWithFilter2(name string, nim int, lulus bool, filter Filter2){
	a, b, c := filter(name, nim, lulus);
	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);
}

func spamFilter2(name string, nim int, lulus bool)(string, int, bool){
	return name, nim, lulus;
}


func main(){
	sayHelloWithFilter("Rafi Alfian", spamFilter);
	
	filter := spamFilter;
	sayHelloWithFilter("Ervin Alfianto", filter);

	//-------------------------------------

	filter2 := spamFilter2
	sayHelloWithFilter2("Rafi Alfian", 21, true, filter2);
}

/*
	Function as parameter
	1. Function tidak hanya bisa kita simpan didalam variabel sebagai value.
	2.Namun jiga bisa digunakan sebagai parameter untuk function lain.
*/

/*
	Function type declarations
	1. Kadang jika function terlalu panjang, agak ribet untuk menuliskannya didalam komputer
	2. Type Declaratiov juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter
*/


