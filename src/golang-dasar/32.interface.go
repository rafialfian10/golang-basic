package main;

import "fmt";

type HasName interface {
	GetName() string
}

func sayHello(hasname HasName){
	fmt.Println("Hello", hasname.GetName())
}

// Person
type Person struct {
	Name string
}

func (person Person) GetName() string { // struct method ini telah mengikuti kontrak dari function GetName() maka nantinya dapat dijadikan parameter untuk sayHello
	return person.Name
}
//----------------------------

// Animal
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
//------------------------------------------

func main() {
	
	// sayHello() 
	// data parameter tidak bisa dimabil dari interface secara langsung. Seperti rafi := Hasname lalu rafi.GetName(), lalu rafi dimasukkan sebagai parameter itu tidak bisa

	rafi := Person {
		Name: "Rafi Alfian",
	}

	cow := Animal {
		Name: "Cow",
	}

	sayHello(rafi) // jika seperti ini baru bisa itu karena function sayHello kontraknya adalah HasName. Itu artinya siapapun / tipe data apapun yang mempunyai function GetName() string, maka dia berhak untuk mengikuto kontrak
	sayHello(cow);
}

/*
	Interface
	1. Adalah sebuah tipe data abstract, dia tidak implementasi secara langsung.
	2.Sebuah interface berisikan definisi-definisi method
	3.Biasanya interface digunakann sebagai kontrak

	Implementasi Interface
	1. Setiap tipe data yang sesuai kontrak interface, secara otomatis dianggap sebagai interface tersebut.
	2. Sehingga kita tidak perlu mengimplementasikan interface secara manual.
	3. Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface yang mana
*/ 


