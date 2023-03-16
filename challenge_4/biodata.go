package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
}

var daftarTeman = map[int]Teman{
	1: Teman{
		Nama:      "Ahmad",
		Alamat:    "Jl. Ahmad Yani 123, Surabaya",
		Pekerjaan: "Pengusaha",
	},
	2: Teman{
		Nama:      "Budi",
		Alamat:    "Jl. Sudirman 456, Jakarta",
		Pekerjaan: "Programmer",
	},
	3: Teman{
		Nama:      "Cindy",
		Alamat:    "Jl. Pahlawan 789, Bandung",
		Pekerjaan: "Web Developer",
	},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go <nomor absen>")
		os.Exit(1)
	}

	nomorAbsen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Nomor absen harus berupa angka: %s\n", os.Args[1])
		os.Exit(1)
	}

	teman, ok := daftarTeman[nomorAbsen]
	if !ok {
		fmt.Printf("Teman dengan nomor absen %d tidak ditemukan\n", nomorAbsen)
		os.Exit(1)
	}

	fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\n", teman.Nama, teman.Alamat, teman.Pekerjaan)
}
