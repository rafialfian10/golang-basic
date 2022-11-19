package main;

import (
	"fmt"
	"container/ring"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5);

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10); // isi data akan diisi dengan konversi int menjadi string
		data = data.Next();
	}

	// // fmt.Println(*data); // hasilnya tidak bisa diprintln karena dia tidak punya function untuk println, untuk menanganinya, maka gunakan function bawaan container/ring

	data.Do(func(DataValue interface{}){ // function bawaan ring
		fmt.Println(DataValue)
	})
}

/*
	Package Container Ring
	1. Adalah implementasi struktur data circular list
	2. Circular list adalah struktur data ring, dimana akhir element akan kembali ke element awal (HEAD)
	3. https://golang.org/pkg/container/ring/
*/ 