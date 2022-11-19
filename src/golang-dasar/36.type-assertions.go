package main;

import "fmt";


// func Random() interface{} {
// 	return "Tipe Data";
// }


// func main(){
// 	result := Random();
// 	resultString := result.(string); // jika misal type data string diganti dengan type data lain maka akan terjadi panic. Dan bisanya untuk menghindari kasus tersebut maka gunakan switch case
// 	fmt.Println(resultString);

// 	switch value := result.(type) {
// 	case string: 
// 		fmt.Println("String", value);
// 	case int: 
// 		fmt.Println("Int", value);
	
// 	case bool: 
// 		fmt.Println("Boolean", value);
	
// 	default: 
// 		fmt.Println("Type data unknown")
// 	}

// }


func Random(data interface{}) interface{} {
	return data;
}

func main() {
	result := Random("Rafi Alfian");

	switch value := result.(type) {
		case string: {
			fmt.Println("String", value);
		}
		case int: {
			fmt.Println("Int", value);
		}
		default: {
			fmt.Println("Boolean", value);
		}
	}
}

/*
	Type Assertions
	1. Type assersions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan.
	2. Fitur ini sering kali digunakan ketika kita bertemu dengan data interface kosong.

	Type Assertions menggunakan switch
	1. Saat salah menggunakan type assertions, maka bisa berakibat terjasi panic di aplikasi kita.
	2. Jika panic tidak ter-recover, maka otomatis program akan berhenti
	3. Agar lebih aman, sebaiknya gunakan switch expression untuk melakukan type assertions
*/ 