package main

import "fmt"

func main() {
	var mahasiswa = map[string]int{"Aldo": 182, "Yosep": 178}
	tampil_mahasiswa(mahasiswa)
}

func tampil_mahasiswa(mahasiswa map[string]int) {

	fmt.Printf("Aldo : %d cm\n", mahasiswa["Aldo"])
	fmt.Printf("Yosep : %d cm\n", mahasiswa["Yosep"])

}
