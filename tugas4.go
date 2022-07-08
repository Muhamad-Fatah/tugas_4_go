package main

import "fmt"

func main() {
	var tinggi = map[string]int{"Aldo": 182, "Yosep": 178}
	var t1, t2 = tampil_mahasiswa()
	fmt.Println("Aldo : ", tinggi[t1], "cm")
	fmt.Println("Yosep : ", tinggi[t2], "cm")
}

func tampil_mahasiswa() (string, string) {
	var a = "Aldo"
	var y = "Yosep"
	return a, y
}
