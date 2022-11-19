package main;

import "fmt";

func getHello(name string)string{ // string diluar kurung berarti harus mereturn data string
	// return "Hello " + name + " Good Morning!"; 

	// atau dapat diberi kondisi

	if name == "" {
		return "Hello, what is your name?";
	} else {
		return "Hello, my name is " + name;
	}
}

func main() {	
	result := getHello("");
	fmt.Println(result);

	fmt.Println(getHello("Rafi Alfian"));
}