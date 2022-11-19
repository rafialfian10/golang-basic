package main;

import "fmt";

func main() {
	type noKTP string;
	type statusKuliah bool;

	var KtpRafi noKTP = "3320122408960005";
	var lulus statusKuliah = true;
	var tidakLulus statusKuliah = false;

	fmt.Println(KtpRafi);
	fmt.Println(KTP("0000000000000000"));

	fmt.Println(lulus);
	fmt.Println(tidakLulus);

}