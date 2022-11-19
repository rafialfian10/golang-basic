package main;

import "fmt";

func main() {
	var array1[10] string; // Maksimal data yang dapat ditampung adalah 3
	array1[0] = "Universitas";
	array1[1] = "Muhammadiyah";
	array1[2] = "Surakarta";

	fmt.Println(array1[0]);
	fmt.Println(array1[1]);
	fmt.Println(array1[2]);

	var array2 = [3] int { // Membuat array secara langsung saat deklarasi variabel dengan membungkus dengan {} dan beri = sebelum nama var
		10,
		20,
		30,
	}

	fmt.Println(array2[0]);
	fmt.Println(array2[1]);
	fmt.Println(array2[2]);

	//Operasi
	fmt.Println(len(array1)); // len, hasilnya adalah 10 meskipun isi array hanya 3
	fmt.Println(array1[1]); // Hasilnya muhammadiyah

	var array3 = [...]int { // Data yang ditampung array bebas
		10, 
		20,
		30,
		40,
		50,
	}

	fmt.Println(array3);
	fmt.Println(len(array3));
	array3[0] = 100;
	fmt.Println(array3);
}


