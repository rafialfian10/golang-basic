package main;

import "fmt";

type Biodata struct {
	name, email, nohp string
	nim int
}

func main(){
	//cara membuat data struct
	// cara 1
	rafi := Biodata {
		name: "Rafi Alfian",
		nim: 21,
		email: "rafialfian770@gmail.com",
		nohp: "08979638899",
	}
	fmt.Println(rafi)

	// cara 2
	var ervin Biodata;
	ervin.name = "Ervin Alfianto";
	ervin.nim = 10;
	ervin.email = "ervin@gmail.com";
	ervin.nohp = "085465788632";

	fmt.Println(ervin)
}


/*
	Struct
	1. Adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
	2. Struct biasanya representasi data dalam program aplikasi yang kita buat
	3. Data si struct disimpan didalam field
	4. Sederhananya struct dalah kumpulan dari field
*/ 