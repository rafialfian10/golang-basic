package main;

import "fmt";

func main(){
	var nilai_matematika = 70;
	var nilai_ipa = 80;

	var lulus_matematika = nilai_matematika > 60;
	var lulus_ipa = nilai_ipa > 80;

	var result1 = lulus_matematika && lulus_ipa;
	var result2 = lulus_matematika || lulus_ipa;

	fmt.Println(result1);
	fmt.Println(result2);
	fmt.Println(nilai_matematika > 70 && nilai_ipa > 70);
}