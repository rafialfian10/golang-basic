package main;

import "fmt";

func main(){
	days := [...]string {
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jum'at",
		"Sabtu",
		"Minggu",
	}

	var slice1 = days[5:];
	slice1[0] = "Sabtu baru";
	slice1[1] = "Minggu baru";

	fmt.Println(days); // Hasil [Senin Selasa Rabu Kamis Jum'at Sabtu baru Minggu baru] ini juga ikut berubah
	fmt.Println(slice1); // Hasil [Sabtu baru Minggu baru]

	days2 := append(slice1, "Libur baru");
	days2[0] = "Opps"
	fmt.Println(days2); // Hasil [Opps Minggu Baru Libur baru]
	fmt.Println(days); // Hasil [Senin Selasa Rabu Kamis Jum'at Sabtu baru Minggu baru]

	var mounths = [...]string {
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice2 = mounths[4:7];
	var slice3 = mounths[:3];
	fmt.Println(slice2); // hasilnya [Mei, Juni, Juli]
	fmt.Println(slice3); // Hasilnya [Januari Februari Maret]
	
	mounths[0] = "Januari lagi";
	slice4 := mounths[:11];
	mounths2 := append(slice4, "Desember baru")
	fmt.Println(mounths); // [Januari Februari Maret April Mei Juni Juli Agustus September Oktober November Desember baru] ini juga ikut berubah
	fmt.Println(slice4) // Hasil [Januari Februari Maret April Mei Juni Juli Agustus September Oktober November]
	fmt.Println(mounths2) // [Januari Februari Maret April Mei Juni Juli Agustus September Oktober November Desember baru]
//----------------------------------------------------------


// Make Slice
newSlice := make([]string, 2, 5); // panjang isi array hanya boleh 2, sedangkan maksimal kapasitas adalah 5
newSlice[0] = "Rafi";
newSlice[1] = "Alfian";

fmt.Println(newSlice);
fmt.Println(len(newSlice));
fmt.Println(cap(newSlice));

// Copy Slice
copySlice := make([]string, len(newSlice), cap(newSlice)); //misal jika len(1), maka hanya akan mengcopy rafi
copy(copySlice, newSlice); // copy newSlice ke copySlice

fmt.Println(copySlice);


// Penting
// iniArray := [...]int {1,2,3,4,5}; // ini adalah array
// iniSlice := []int {1,2,3,4,5}; // ini adalah slice
}