package main;

import "fmt";

func getHello() (string, string, int) { // jika return lebih dari satu, maka harus memakai kurung ()
	return "Hello, nama saya " + "Rafi", "Alfian" + " umur saya", 25;
}

func getFullName()(string, string, string) {
	return "Rezqy", "Ade", "Chandra";
}

func biodata()(string, int, bool){
	return "Rafi Alfian", 21, true;
}


func main(){
	firstName, lastName, age := getHello();

	name1, name2, age := getHello();
	fmt.Println(name1, name2, age);
	
	// Dapat juga menghiraukan return value
	// firstName, middleName, lastName := getFullName();
	firstName, _, lastName := getFullName();
	fmt.Println(firstName, lastName);


	name, nim, lulus :=biodata();
	fmt.Println("Nama :", name);
	fmt.Println("NIM : ", nim);
	fmt.Println("Lulus : ", lulus)

}

// Returning Multiple values
// 1. Function tidak hanya dapat mengembalikan satu value, tapiu juga bisa multiple value
// 2. Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return valu-nya di function