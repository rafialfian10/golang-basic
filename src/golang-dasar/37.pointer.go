package main;

import "fmt";

type Address struct {
	City, Province, Country string
}

func main() {
	// Pass by Value
	// Jika address2 diambil dari address1, maka address1 akan benar-benar dicopy. Meskipun field address2 diganti,maka adress1 tidak berpengaruh.
	// address1 := Address{"Jepara", "Jawa tengah", "Indonesia"};
	// address2 := address1;

	// Pass by Reverence
	// jika address2 diambil dari address1, maka address2 masih mengacu ke address1, Jika field address2 diubah, maka address1 juga akan berubah
	var address1 Address = Address{"Jepara", "Jawa tengah", "Indonesia"}; // deklarasi variabel lengkap
	var address2 *Address = &address1; // tambahkan & (pointer). tanda bintang adalah pointer
	var address3 *Address = &address1;
	var address4 *Address = &address1;
	
	address2.City = "Semarang"; // Ini hanya mengubah satu field-nya saja, jika kita ingin mengubah seluruh variabel maka gunakan operator *
	// address3 = &Address{"Bandung", "jawa Barat", "Indonesia"}; // jika kita pointer langsung ke Address(&Address), maka yang lain tidak berubah, jika ingin semua ikut berubah, maka operator * seperti dibawah
	*address3 = Address{"Bandung", "Jawa Barat", "Indonesia"}; // jika misal ada address baru (sampai 10 misalnya) jika mengacu ke Address maka akan ikut berubah / atau benar-benar di timpa
	
	address4.City = "Surabaya"; // Jika field address4 diganti maka semua juga aka ikut diganti
	
	
	fmt.Println(address1);
	fmt.Println(address2);
	fmt.Println(address3);
	fmt.Println(address4);

	// Function new. operator & dapat diganti dengan new (fungsi sama saja)
	var address5 *Address = new(Address) // jika menggunakan new, maka harus *Address (ada pointer)
	address5.City = "Lampung";
	address5.Province = "Bandar Lampung";
	address5.Country = "Indonesia";
	fmt.Println(address5);
}

/*
	Pointer (Pass by Reference)
	1. Pointer adalah kemampuan membuat reference ke lokasi data memory yang sama, tanpa menduplikasi data yang sudah ada.
	2. Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

	Pass by Value
	1. secara default di go-langsemua variabel itu dipassing by value, bukan by reference
	2.Artinya jika kita mengirim sebuah variabel ke dalam function, method / variabel lain, sebenarnya yang dikirim adalah duplikasi value-nya

	Operator *
	1. Saat kita mengubah operator pointer, maka yang berubah hanya variabel tersebut
	2. Semua variabel yang mengacu kedata yang sama tidak akan berubah
	3. Jika kita ingin mengubah seluruh variabel yang mengacu ke data tersebut, kita bisa menggunakan operator *

	Function New
	1. Sebelumnya untuk membuat pointer dengan menggunakan operator &
	2. Go-lang juga memiliki function new yang bisa digunakan untuk membuat pointer
	3. Namun function new hanya mengembalikan pointer ke data yang kosong, artinya tidak ada data awal
*/ 
