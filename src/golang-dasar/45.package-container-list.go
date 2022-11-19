package main;

import (
	"fmt"
	"container/list"
)

func main() {
	data := list.New();
	data.PushBack("Rafi");
	data.PushBack("Ervin");
	data.PushBack("Ali");
	data.PushFront("Chanief");

	fmt.Println(data.Front().Value); // melihat data terdepan
	fmt.Println(data.Back().Value); // melihat data terbelakang

	//jika ingin melihat seluruh data dari depan
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value);
	}

	//jika ingin melihat seluruh data dari belakang
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value);
	}
}

/*
	package Container /List
	1. Package container / list adalah implementasi dari struktur data double linked list di Go-Lang
	2. https://golang.org/pkg/container/list
*/ 