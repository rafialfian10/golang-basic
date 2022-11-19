package main;

import (
	"fmt"
	"sort"
)

type User struct {
	Nama, Email, Alamat string
	Age int
}

type UserSlice []User // aliaskan struct

func (userslice UserSlice) Len() int {
	return len(userslice);
}

func (userslice UserSlice) Less(i, j int) bool { //i: index 1, j: index 2
	return userslice[i].Age < userslice[j].Age
} 

func (userslice UserSlice) Swap(i, j int) {
	userslice[i], userslice[j] = userslice[j], userslice[i]
}


func main() {

	users := []User { 
		{"Rafi Alfian", "rafialfian770@gmail.com", "Jepara", 25},
		{"Ervin Alfianto", "ervinalfianto@gmail.com", "Jepara", 27},
		{"Ali Imron", "ali@gmail.com", "Kudus", 26},
		{"Chanief", "chanief@gmail.com", "Semarang", 28},
	}

	sort.Sort(UserSlice(users)) // lalu urutkan dengan sort.Sort(). Tapi tidak bisa langsung memanggil users, karena sort butuh data interface, jadi panggil dulu alias UserSlice karena alias ikut kontrak interface
	fmt.Println(users)


}

/*
	Package Sort
	1. Merupakan package yang berisikan utilitas untuk proses pengurutan
	2. Agar data dapat diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
	3. https://golang.org/pkg/sort
*/ 