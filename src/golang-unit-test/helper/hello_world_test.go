package helper

import (
	"testing"
	"fmt"
)

func TestHelloWorldRafi(t *testing.T) {
	result := HelloWorld("Rafi") // function dari file hello_world.go
	if result != "Hello Ervin" {
		// unit test failed
		// panic("Result is not 'Hello Rafi'") // Menggunakan panic bukanlah hal bagus, karena panic akan langsung menghentikan function

		// t.Fail() // jika terdapat error, maka program dibawahnya akan tetap dilanjutkan, tapi tetap dianggap gagal. Kekurangan Fail() / FailNow() tidak dapat memberikan pesan didalamnya

		t.Error("Result must be 'Hello Rafi'") // jika terdapat error, maka program dibawahnya akan tetap dilanjutkan, tapi tetap dianggap gagal
	}

	fmt.Println("TestHelloWorldRafi done")
}

func TestHelloWorldAlfian(t *testing.T) {
	result := HelloWorld("Alfian") // function dari file hello_world.go
	if result != "Hello Alfianto" {
		// t.FailNow() // jika terdapat error, maka program dibawahnya akan berhenti 
		
		t.Fatal("Result must be 'Hello Alfian'") // jika terdapat error, maka program dibawahnya akan berhenti 
	}

	fmt.Println("TestHelloWorldAlfian done")
}

/*
	Aturan File Test
	1. Go-lang memiliki aturan cara membuat file khusus untuk unit test
	2. Untuk membuat file unit test, kita harus menggunakan akhiran _test
	3. Jadi misalkan kita membuat file hello_world.go, artinya untuk membuat unit test-nya, kita harus membuat file hello_world_test.go. nama bebas, tapi akhiran harus ada _test dan harus berada di package yang sama
*/

/*
	Aturan Function Unit Test
	1. Selain aturan nama file, di Go-lang juga sudah diatur untuk nama function unit test
	2. Nama function untuk unit testharus diawali dengan nama Test
	3. Misal jika ingin mengetes function HelloWorld, maka kita akan membuat function unit test dengan nama function TestHelloWorld
	4. Tak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan ditest, yang penting harus diawali dengan kata Test, artinya nama belakang bebas tapi harus diawali dengan Test didepan
	5. Selanjutnya harus memiliki parameter (t *testing.T) dan tidak mengembalikan return value
*/

/*
	Menjalankan Unit Test
	1. Untuk menjalankan unit test kita bisa menggunakan perintah: go test
	2. Jika kita ingin lihat lebih detail function test apa saja yang sudah di running, kita bisa gunakan perintah: go test -v
	3. Dan jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah: go test -v -run=TestNamaFunction

	Menjalankan Semua Unit Test
	1. Jika kita ingin menjalankan semua unti test dari top folder module-nya, kita bisa gunakan perintah: go test -v ./...
*/ 

/*
	Mengagalkan Unit Test
	1. Menggagalkan unit test menggunakan panic bukanlah hal bagus
	2. Go-lang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
	3. Terdapat function fail(), FailNow(), Error() dan Fatal() jika ingin menggagalkan unit test

	t.Fail() dan t.FailNow()
	1. Terdapat dua funvtion untuk menggagalkan unit test yaitu Fail() dan FailNow(). Lantas apa bedanya?
	2. Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test (kode program selanjutnya bukan funtion test selanjutnya). Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal
	3. FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test

	t.Error(args...) dan t.Fatal(args...)
	1. Selain Fail() dan FailNow(), ada juga Error() dan Fatal()
	2. Error() function lebih seperti melakukan log(print) error, namun setelah malakukan log error, dia akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal
	3. Namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai
	4. Fatal() mirip degan Error(), hanya saja setelah melakukan log error, dia akan memanggil FailNow(), sehingga mengakibatkan sksekusi unit test berhenti
*/