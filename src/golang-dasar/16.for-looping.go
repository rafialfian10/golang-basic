package main;

import "fmt";

func main(){

// Cara 1
	counter1 := 1;
	for counter1 <= 10 {
		fmt.Println("Perulangan counter 1 ke ", counter1);
		counter1++;
	}
	
// Cara 2 (for dengan statement)
	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan counter 2 ke", counter2)
	}

// Cara 3 (for range) digunakan untuk melakukan iterasi terhadap semua data collection contoh (array, map, slice)
	slice := []string{"Rafi", "Ervin", "Ali"} //slice

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i]);
	}


// Cara 4 (for range)
	names := [...]string {"Rama", "Satria", "Arta"} //array

	for _, name := range names { //index dapat diganti dengan _ agar index tidak perlu dicetak
		// fmt.Println("Index ", index, "= ", name);
		fmt.Println("Nama :", name);
	}

// Cara 5 (for range)

	// biodata := map[string]string {
	// 	biodata["name"] = "Rafi Alfian";
	// 	biodata["age"] = "25";
	// 	biodata["email"] = "rafialfian770@gmail.com";
	// }
	
	var biodata map[string]string = make(map[string]string); //map
		biodata["name"] = "Rafi Alfian";
		biodata["age"] = "25";
		biodata["email"] = "rafialfian770@gmail.com";
	

	for i := 0; i < len(biodata); i++ {
		fmt.Println("Index ke", i+1, "adalah", biodata)
	}

}