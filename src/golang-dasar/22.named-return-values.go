package main;

import "fmt";

func getBiodata()(nama string, nim int, lulus bool){
	nama = "Rafi Alfian";
	nim = 21;
	lulus = true;

	// return nama, nim, lulus;
	return; // atau dapat ditulis return saja
}

func main(){	
	// a,b,c := getBiodata();
	// fmt.Println(a,b,c);

	name, nim, lulus := getBiodata();
	fmt.Println(name, nim, lulus);
}

// Named Return Values (Fitur yang hanya ada di golang)
// Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function
// Namun kita juga bisa membuat variabel secara langsung di tipe data return functionnya