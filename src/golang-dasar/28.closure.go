package main;

import "fmt";

func main() {
	counter := 0;
	name := "Rafi"

	increment := func(){
		// name := "Alfian" // Solusi agar name tidak ditimpa dalah dengan membuat variabel baru
		fmt.Println("Increment");
		// fmt.Println(name);
		counter++;
	}


	increment();
	increment();
	fmt.Println(counter) // hasilnya akan 2 karena value pada counter ditimpa oleh function
	fmt.Println(name); // hasilnya akan menjadi Alfian karena ditimpa oleh name yang ada didalam function
}