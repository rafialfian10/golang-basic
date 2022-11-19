package main;

import "fmt";

type Man struct {
	Name string
}

func (man *Man) married() { // Agar data dapat berubah maka tambahkan pointer pada struct
	man.Name = "Mr. " + man.Name;
	fmt.Println("Di method :", man.Name)
}
//-------------------------------------------------

type Biodata struct {
	Nama, Email string
	Nim int
}

func (biodata *Biodata) cetakBiodata() {
	biodata.Nama = "Ervin Alfianto";
	biodata.Email = "ervin@gmail.com";
	biodata.Nim = 10;
	fmt.Println("Hello nama saya adalah", biodata.Nama, "email", biodata.Email, "NIM saya adalah", biodata.Nim);
}
//----------------------------------------------------

func main() {
	rafi := Man{"Rafi"}; // struct
	rafi.married(); // ini akan memanggil baris 10;

	fmt.Println(rafi.Name) // Tambahkan Name agar {} hilang saat di eksekusi 
	//-----------------------------------------------------

	var r Biodata = Biodata {"Rafi Alfian", "rafi@gmail.com" , 21};
	r.cetakBiodata();
	fmt.Println(r.Nama, r.Email, r.Nim); // Hssil tetap Ervin Alfianto ervin@gmail.com 10
}

// /*
// 	Pointer di Method
// 	1. Walaupun method akan menempel di struct, tapi sebenarnya data structyang diakses didalam method adalah pass by value
// 	2. Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method.
// */ 


