package main;

import "fmt";

func getName(name string)string {
return "Hello " + name;
}

func main(){
	name := getName; // function dapatdisimpan kedalam variabel tapi jangan gunakan ()
	fmt.Println(name("Rafi Alfian"));

	name2 := getName;
	result := name2("Ervin Alfianto");
	fmt.Println(result);
}


/*
	Function Value
	1. function adalah first class citizen
	2. Function juga merupakan tipe data, dan bisa disimpan didalam variable
*/