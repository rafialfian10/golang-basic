// package main;

// import "fmt";

// func Newmap(name string) map[string]string {
// 	if name == "" {
// 		return nil;
// 	} else {
// 		return map[string]string {
// 			"name": name,
// 		}
// 	}
// }

// func biodata(nama string, nim int, lulus bool) map[string]interface{} {
// 	if nama == "" {
// 		return nil;
// 	} else if nim == 0 {
// 		return nil;
// 	} else if lulus == false {
// 		return nil;
// 	} else {
// 		return map[string] interface{} {
// 			"Nama": nama,
// 			"Nim": nim,
// 			"Lulus": lulus,
// 		}
// 	}
// }

// func main(){
// 	data1 := Newmap("");
// 	data2 := Newmap("Rafi Alfian");
// 	fmt.Println(data1);
// 	fmt.Println(data2);
// 	//----------------------------------

// 	bio := biodata("Rafi", 21, true);
// 	fmt.Println(bio);
// }

// /*
// 	Nil
// 	1. Biasanya dalam bahasa pmrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null / nil.
// 	2. Berbeda dengan go-lang, di go-lang saat kita membuat veriabel dengan tipe data tertentu, makasecara otomatis akan dibuatkan default valuenya
// 	3. namun di go-lang ada data nil, yaitu data kosong
// 	4.Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel.
// */ 

package main
import "fmt"

func GetBiodata(name string, email string, nim int, lulus bool)map[string] interface{} {
	if name == "" {
		return map[string]interface{} {
			"Nama": nil,
			"Email": email,
			"Nim": nim,
			"Status Kuliah": lulus,
		}
	} else if email == "" {
		return map[string]interface{} {
			"Nama": name,
			"Email": nil,
			"Nim": nim,
			"Status Kuliah": lulus,
		}
	} else if nim == 0 {
		return map[string]interface{} {
			"Nama": name,
			"Email": email,
			"Nim": nil,
			"Status Kuliah": lulus,
		}
	} else if lulus == false {
		return map[string]interface{} {
			"Nama": name,
			"Email": email,
			"Nim": nim,
			"Status Kuliah": nil,
		}
	} else {
		return map[string]interface{} {
			"Nama": name,
			"Email": email,
			"Nim": nim,
			"Status Kuliah": lulus,
		}
	}
}

func main() {
	result := GetBiodata("Rafi Alfian", "rafialfian770@gmail.com", 21, true);
	fmt.Println(result);
}