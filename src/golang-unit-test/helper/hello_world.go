package helper

func HelloWorld(name string) string {
	return "Hello " + name
	// return "Hi " + name
}

/*
	Pengenalan Software Testing

	1. Software Testing adalah salah satu disiplin ilmu dalam software engineering
	2. Tujuan utama dari software testing adalah memastikan kualitas kode dan aplikasi kita baik
	3. Ilmu untuk software testing sendiri sangatlah luas, pada materi kali ini hanya fokus ke unit testing
	
	Test pyramid
	dalam unit testing terdapat istilah test pyramid mulai dari yang paling bawah test unit, test service, test UI / End to End Test
	test unit adalah yang paling murah implementasinya, mudah dalam implementasi, mudah untuk dibuat / simple, dan eksekusi semakin cepat,
	sedangkan jika semakin ke atas(Test Service dan Test UI), maka implementasnya akan lebih mahal / sulit, dan eksekusinya akan semakin lambat
	---------------------------------------------------------------------------------------------------------------------

	Unit Test

	1. Unit test akan fokus menguji bagian kode program terkecil, biasanya menguji sebuah function / method
	2. Unti test biasanya dibuat kecil dan cepat, oleh karena itu biasanya kadang kode unit test lebih banyak dari kode program aslinya, karena semua skenario pengujian akan dicoba di unti test
	3. Unit test biasa digunakan sebagai cara untuk meningkatkan kualitas kode program kita
*/ 



/*
	Pengenalan Testing Package

	1. Dibahasa pemrograman lain, biasanya untuk implementasi unit test, kita butuh library / framwork khusus untuk unit test
	2. Berbeda dengan golang, golang sudah ada unit test yang sudah disediakan sebuah package khusus bernama testing
	3. Selain itu untuk menjalankan unit test, golang juga sudah disediakan perintahnya
	4. Hal ini membuat implementasi unit testing digolang sangat mudah dibanding dengan bahasa pemrograman lain
	5. https://golang.org/pkg/testing


	testing.T
	1. Go-lang menyediakan sebuah struct yang bernama testing.T
	2. Struct ini digunakan untuk unit test di golang

	testing.M
	1. testing.M adalah struct yang disediakan Go-lang untuk mengatur life cicle testing
	2. Materi ini akan dibahas di chapter main

	testing.B
	1. testing.B adalah struct yang disediakan Go-lang untuk melakukan benchmarking (mengukur kecepatan kode program)
	2. di Go-lang untuk melakukan benchmarking(mengukur kecepatan kode program) pun sudah disediakan, sehingga kita tidak perlu menggunakan library / framework tambahan
*/ 


