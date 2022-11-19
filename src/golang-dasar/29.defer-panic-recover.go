package main;

import "fmt";

// Defer
func logging() {
	fmt.Println("Aplikasi selesai");
}

func runApplication(value int) {
	defer logging(); // walaupun function ini error, maka function logging akan tetap dieksekusi karena menggunakan defer()
	fmt.Println("Aplikasi berjalan")
	result :=  10 / value; // 10 / 0 akan error, meskipun begitu, looging akan tetap dieksekusi
	fmt.Println(result);
}
//---------------------------------------------------

//Panic
func endApp() {
	message := recover(); // recover harus berada di defer(endAApp), jadi pesan pada panic("Error(panic)") akan ditangkap oleh recover(). Tetapi jika tidak terjadi error maka recover akan mengembalikan nilai <nil>, maka lakukan kondisi untuk menghilangkan nil
	if message != nil {
		fmt.Println("Terjadi error", message)
	} else {
		fmt.Println("End App");
	}
}

func runApp(error bool) {
	defer endApp();
	if error {
		panic("Error(panic)") // jika error bernilai false maka defer akan tetap dieksekusi, meskipun runApp masih error. Jika terjadi panic, maka code dibawah tidak akan dijalankan
	} else {
		fmt.Println("Applikasi berjalan")
	}
}
//------------------------------------------

func main() {
	//defer
	runApplication(10);

	//panic
	runApp(true)

}

/*
	Defer
	1. adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi.
	2. Defer function akan selalu dieksekusi walapun terjadi error di function yang dieksekui

	Panic
	1. panic function adalah function yang bisa digunakan untuk menghentikan program
	2. panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
	3. Saat panic function dipanggil, program akan terhenti, namun defer function akan tetap dieksekusi

	Recover
	1.adalah function yang bisa gunakan untuk menangkap data panic
	2.Dengan recover proses panic akan berhenti, sehingga program akan tetap berjalan
*/ 

