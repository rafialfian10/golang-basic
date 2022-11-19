package main

import (
	"fmt"
	"time"
)

func main() {

	// Membuat waktu dengan time.Now()
	now := time.Now()

	fmt.Println("Time now :", now)
	fmt.Println("Year", now.Year());
	fmt.Println("Month", now.Month());
	fmt.Println("Day", now.Day());
	fmt.Println("Hour", now.Hour());
	fmt.Println("Minute", now.Minute());
	fmt.Println("Second", now.Second());
	fmt.Println("Nano Second", now.Nanosecond());


	// Membuat waktu secara manual menggunakan time.Date(...)
	utc := time.Date(1996, 8, 24, 23, 45, 10, 10, time.UTC)
	fmt.Println("Time Date", utc)

	// Membuat waktu dengan time.Parse(layout, string)
	layout := "2006-01-02" // ini adalah cara membuat layout pada golang, beda dengan bahasa pemrograman lainnya "yyyy-mm-dd"
	// parse, _ := time.Parse(layout, "2022-08-24") 
	
	parse, err := time.Parse(layout, "2022-08-24") // parse terdapat 2 parameter
	if err == nil {
		fmt.Println("Time Parse", parse)
	} else {
		fmt.Println("Error", err.Error())
	}
}

/*
	Package Time
	1. Adalah package yang berisikan fungsionalitas untuk management waktu di golang
	2. https://golang.org/pkg/time

	Beberapa Function di Package Time
	time.Now()						Untuk mendapatkan waktu saat ini
	time.Date(...) 					Untuk membuat waktu
	time.Parse(layout, string)		Untuk memparsing waktu daris string
*/ 