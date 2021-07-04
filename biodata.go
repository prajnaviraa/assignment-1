package main

import (
	"fmt"
	"os"
)

type student struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var st student
	abs := os.Args[1]
	st = input(abs)
	printMessage(st)
}

func printMessage(sf student) {
	fmt.Println("Nama       :", sf.nama)
	fmt.Println("Alamat     :", sf.alamat)
	fmt.Println("Pekerjaan  :", sf.pekerjaan)
	fmt.Println("Alasan     :", sf.alasan)
}

func input(absen string) (s student) {
	if absen == "1" {
		s.nama = "Dianda"
		s.alamat = "alamat 1"
		s.pekerjaan = "Pekerjaan 1"
		s.alasan = "alasan 1"
	}
	if absen == "2" {
		s.nama = "Namaz Al Fattah"
		s.alamat = "alamat 2"
		s.pekerjaan = "Pekerjaan 2"
		s.alasan = "alasan 2"
	}
	if absen == "3" {
		s.nama = "Arya"
		s.alamat = "alamat 3"
		s.pekerjaan = "Pekerjaan 3"
		s.alasan = "alasan 3"
	}
	if absen == "4" {
		s.nama = "Muhammad Naufal Rachfian Djamhur"
		s.alamat = "alamat 4"
		s.pekerjaan = "Pekerjaan 4"
		s.alasan = "alasan 4"
	}
	if absen == "5" {
		s.nama = "Wirsal Djamaluddin"
		s.alamat = "alamat 5"
		s.pekerjaan = "Pekerjaan 5"
		s.alasan = "alasan 5"
	}
	if absen == "6" {
		s.nama = "Deny Prasetyo"
		s.alamat = "alamat 6"
		s.pekerjaan = "Pekerjaan 6"
		s.alasan = "alasan 6"
	}
	if absen == "7" {
		s.nama = "Prajnavira Taslim"
		s.alamat = "alamat 7"
		s.pekerjaan = "Pekerjaan 7"
		s.alasan = "alasan 7"
	}
	return
}
