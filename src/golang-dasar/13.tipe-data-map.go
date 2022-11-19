package main;

import "fmt";

func main(){
	person := map[string]string { // [string] ---> key sedangkan string ---> value
		"name": "Rafi Alfian",
		"age": "25",
		"lulus" : "true",
	}

	person["lulus"] = "false"; // false

	fmt.Println(person);
	fmt.Println(len(person));  // 2
	fmt.Println(person["name"]); // Rafi Alfian
	fmt.Println(person["age"]); // 25


	// Membuat map baru cara 2 (tanpa kurung kurawal)
	var anime map[string]string = make(map[string]string);
	anime["title"] = "Shingeki no Kyojin";
	anime["artist"] = "Hajime Isayama";
	anime["genre"] = "Action, Horror, Mistery, Military";

	fmt.Println(anime);
	delete(anime, "artist");
	fmt.Println(anime);

}

// Function Map
/*
Operasi						Keterangan
len(map)					Untuk mendapatkan jumlah data di map
map[key]					Untuk mengambil data di map dengan key
map[key] = value			Mengubah data di map dengan key
make([TypeKey]TypeValue)	Membuat map baru
delete(map, key)			Menghapus data di map dengan key
*/

