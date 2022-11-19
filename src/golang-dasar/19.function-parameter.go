package main;

import "fmt";


func sayHello(firtsParameter string, lastParameter string){
	fmt.Println(lastParameter, firtsParameter,)
}

func biodata(name string, age int, lulus bool ){
	fmt.Println("Nama :", name,"\nUmur :", age,"\nStatus :", lulus);
}


func main(){
	sayHello("Rafi", "Good Night")

	name := "Rafi Alfian";
	age := 25;
	lulus := true;
	
	biodata(name, age, lulus);
}