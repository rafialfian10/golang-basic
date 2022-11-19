package main;

import "fmt";


func sayHello(){
	name := "Rafi Alfian"
	fmt.Println("Hello", name);
}

func main(){
	sayHello();
	sayHello();
	sayHello();



	for i := 0; i < 10; i++ {
	 	sayHello();
	}
}