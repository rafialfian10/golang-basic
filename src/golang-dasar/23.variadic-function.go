package main;

import "fmt";

func getNumber1(numbers ...int) int{ // varargs parameter harus berapa dibelakang/final, jika getNumbers(numbers ...int, name string) maka ini error, jika seperti ini masih bisa dilakukan getNumbers(nama string, numbers ...int).
	total := 0;
	for _, number := range numbers {
		total += number;
	}
	return total;
}

func getNumber2(numbers ...int)int { // int diluar kurung artinya kita wajib mereturn sesuatu didalam function
	total := 0;
	for i := 0; i < len(numbers); i++ {
		total += numbers[i];
	}
	return total;
}

func getNameDatas(datas ...string) {	
	for i, data := range datas {
		fmt.Println(i+1, data)
	}
}

func main(){

	getTotal1 := getNumber1(10, 10, 10, 10, 10);
	fmt.Println(getTotal1)
	//--------------------------

	// Slice Parameter
	// 1. Kadang ada kasus dimana kita menggunakan variadic functio, namun memiliki berupa slice
	// 2. Kita dapat menjadikan slice sebagai varargs parameter. contoh kasus:

	number := []int{10, 20, 30, 40, 50} // misal jika kita sudah punya slice maka tinggal kita masukkan kedalam parameter dengan menambahkan ...
	getTotal2 := getNumber2(number...);
	fmt.Println(getTotal2)
	//--------------------------------------

	nameDatas := []string {
		"Rafi", 
		"Ervin",
		"Ali",
	}
	getNameDatas(nameDatas...);

}


/*
	Variadic Function

	1. Parameter yang berada diposisi terakhir, memiliki kemampuan dijasikan varargs.
	2. varargs artinya datanya menerima lebih dari 1 input, atau dapat dianggap semacam array.
	3. Apa bedanya dengan parameter biasa dengan tipe data array
		a. Jika parameter data array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function.
		b. Jika parameter menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satu. cukup gunakan tanda koma (,)
*/