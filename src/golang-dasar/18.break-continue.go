package main;

import "fmt";

func main(){

	Break
	counter := 10;
	for i := 0; i < counter; i++ {
		if i == 5 {
			break; // Break akan menghentikan perulangan total
		}
		fmt.Println("Perulangan ke" ,i) // Hasil index 0 - 4 karena berhenti pada index 4
	}

	// Continue
	counter2 := 10;

	for i := 0; i < counter2; i++ {
		if i % 2 == 0 {
			continue; // continue akan menghentikan perulangan tapi akan melanjutkanke perulangan selanjutnya
		}
		fmt.Println("Index ke", i, "adalah bilangan ganjil"); 
	}

}

// Break digunakan untuk menghentikan seluruh perulangan
// Continue digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya