package main;

import "fmt";

func getData(input int) interface{} {
	if input == 10 {
		return input;
	} else if input == 20 {
		return true;
	} else {
		return "string";
	}
}

func main() {
	var data interface{} = getData(10);
	fmt.Println(data);
}

// Interface kosong biasanya digunakan untuk menerima parameter yang tipe data-nya lebih dari satu 
