package helper

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMain
func TestMain(m *testing.M) {
	fmt.Println("Sebelum unit dieksekusi")

	m.Run() // ini akan menjalankan / mengeksekusi semua unit test dibawah, meskipun terdapat unit test yang gagal, maka before dan after akan tetap dieksekusi. TestMain hanya berjalan di satu package ini saja. Jika kita punya package lain, maka buat TestMain lagi

	fmt.Println("Setelah unit dieksekusi")
}

// Assert
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorldTestify("Dawam");
	assert.Equal(t, "Hello Dawam", result, "Result must be 'Hello Dawam'") // par1: t, par2: nilai yang kita harapkan, par3: resultnya, par4: message, jika misal terjadi error, maka akan secara otomatis memanggil Fail() yang artinya eksekusi akan tetap dijalankan sehingga bisa kita beri pesan dibawahnya
	fmt.Println("TestHelloWorldTestify Done") // ini akan tetap di print
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorldAssertion("Dawam");
	require.Equal(t, "Hello Dawam", result, "Result must be 'Hello Dawam'") // par1: t, par2: nilai yang kita harapkan, par3: resultnya, par4: message, jika misal terjadi error, maka akan secara otomatis memanggil Fail() yang artinya eksekusi akan tetap dijalankan sehingga bisa kita beri pesan dibawahnya
	fmt.Println("TestHelloWorldRequire Done") // ini tidak akan di print

	// Menggunaka library Assertion / Require lebih baik daripada Fail() / FailNow() dan Error() / Fatal()
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" { // jika sistem operasi adalah darwin (Mas OS)
		t.Skip("Unit test tidak dapat berjalan di Mac OS")
	}

	result := HelloWorldTestify("Dawam");
	require.Equal(t, "Hello Dawam", result, "Result must be 'Hello Dawam'")
}

/*
	Assertion & Require

	Assetion
	1. Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan
	2. Apalagi jika result data yang harus di cek itu banyak
	3. Oleh karena itu, sangat disarankan untuk menggunakan assertion untuk melakukan pengecekan
	4. Sayangnya, Go-lang tidak menyediakan package untuk assertion, sehingga kita butuh menambahkan library untuk melakukan assertion ini

	Testify (Library Assertion)
	1. Salah satu library assertion yang paling populer di Go-lang adalah Testify
	2. Kita bisa menggunakan library ini untu melakukan assertion result data di unit test
	3.https://github.com/stretchr/testify
	4. Kita bisa menambahkan di Go module kita: go get github.com/stretchr/testify

 	Require
	1. Testify menyediakan dua package untuk asseertion, yaitu assert dan require
	2. Saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil Fail(), artinya eksekusi unit test akan tetap dilanjutkan
	3. Sedangkan jika kita menggunakan require, jika pengecekan gagal, maka require akan memanggil FailNow(), artinya eksekusi unit test tidak akan dilanjutkan
*/


/*
	Skip Test (Membatalkan Test)
	1. Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
	2. Di Go-lang juga kita bisa membatalkan eksekusi unit test jika kita mau
	3. Untuk membatalkan unit test kita bisa menggunakan function Skip()
*/


/*
	Before dan After Test
	1. Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan seteah sebuah unit test dieksekusi
	2. Jika kode yang kita lakukan sebelum dan setelah selalu sama antar unit test function, maka membuat manual di unit test function-nya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya
	3. Untungnya di Go-lang terdapat fitur yang bernama testing.M
	4. Fitur ini bernama Mani, dimana digunakan untuk mengatur aksekusi unit test, namun hal ini juga bisa kita gunakan unuk melakukan before dan after di unit test

	Testing.M
	1. Untuk mengatur eksekusi unit test, kita cukup membuat sebuah function bernama TestMain dengan parameter m *testing.M
	2. Jika terdapat function TestMain tersebut, maka secara otomatis Go-lang akan mengeksekusi function ini tiap kali akan menjalankan unit test disebuah package
	3. Dengan ii kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau
	4. Ingat, function TestMain itu hanya akan dieksekusi hanya sekali per Go-Lang package, bukan per tiap function unit test
*/ 