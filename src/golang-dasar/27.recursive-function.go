package main;

import "fmt";

func factorialLoop(value int)int { // cara biasa
	hasil := 1;
	for i := value; i > 0; i--{
		hasil *= i;
	}
	return hasil;
}

func factorialRecursive(value int)int { // cara recursive
	if value == 1 {
		return 1 // artinya return berhenti
	} else {
		return value * factorialRecursive(value - 1)
	}
}

func main(){

	factorial := factorialLoop;
	fmt.Println("Factorial(Looping) = ", factorial(5));

	fmt.Println("Factorial(Recursive) dari 5 = ", factorialRecursive(5))
}

/*
	Recursive Function
	1. Adalah function yang memanggil function dirinya sendiri
	2. Kadang dalam pekerjaan, kita sering menemui dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
	3. Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah factorial
*/