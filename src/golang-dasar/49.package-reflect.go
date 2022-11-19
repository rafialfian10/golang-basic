package main
import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` // ini adalah structTag, cara menambahkan hanya menggunakan backtick(``)
}


// Contoh Kode program Validation Library

func IsValid(datas interface{}) bool {
	data := reflect.TypeOf(datas);
	for i := 0; i < data.NumField(); i++ {
		field := data.Field(i);
		if field.Tag.Get("required") == "true" { // jika true maka lakukan validasi
			value := reflect.ValueOf(datas).Field(i).Interface();  // cara mengambil value: reflect.ValueOf(datas), Field ke i, konversi menjadi interface()
			if value == "" {
				return false;
			}
		}
	}
	return true;
}

func main() {
	sample := Sample{"Rafi Alfian"};
	var sampleType reflect.Type = reflect.TypeOf(sample);


	fmt.Println(sampleType.NumField()); // Hasil 2, karena field pada struct ada 2
	fmt.Println(sampleType.Field(0).Name); // Hasil Name, mirip seperti akses array

	// StructTag
	fmt.Println(sampleType.Field(0).Tag.Get("required")); // Hasil true
	fmt.Println(sampleType.Field(0).Tag.Get("max")); // Hasil 10

	// Kode Program Validasi
	fmt.Println(IsValid(sample)); // maka hasilnya adalah true
	sample.Name = ""
	fmt.Println(IsValid(sample)); // maka hasilnya dalah false karena Name = ""
	
}

/*
	Package Reflect
	1. Dalam bahasa pemrograman, biasanya ada fitur reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan
	2. Hal ini bisa dilakukan di golang dengan menggunakan package reflect
	3. reflection ini sangat berguna ketika ingin membuat library yang general sehingga mudah digunakan
	4. https://golang.org/pkg/reflect
*/ 