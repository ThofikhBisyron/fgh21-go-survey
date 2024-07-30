package main

import (
	"fmt"
)

type Respondent struct {
	fullName     string
	age          int
	gender       string
	isSmoker     bool
	cigarVariant []string
}

func main() {
	respondent := []Respondent{
		Respondent{
			fullName: "RS",
			age:      24,
			gender:   "Laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Surya",
				"Esse",
				"Manchester",
			},
		},
		Respondent{
			fullName: "Yahud",
			age:      55,
			gender:   "Laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Surya",
				"Esse",
				"Manchester",
			},
		},
	}
	for i := 0; i < len(respondent); i++ {

	}
	var name string
	var age int
	var gender string
	var isSmoker bool
	var cigarVariant []string

	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan Umur Anda: ")
	fmt.Scanln(&age)
	fmt.Print("Apa Jenis Kelamin anda: ")
	fmt.Scanln(&gender)
	fmt.Print("Apakah Anda Perokok?")
	fmt.Scanln(&isSmoker)
	fmt.Print("Rokok Apa Saja yang Pernah Anda Coba?")
	fmt.Scanln(&isSmoker)
	fmt.Printf("Nama :%s\n", name)
	fmt.Printf("Umur :%d\n", age)
	fmt.Printf("Jenis Kelamin :%s\n", gender)
	fmt.Printf("Anda Anda Perokok :%t\n", isSmoker)
	fmt.Printf("Anda Anda Perokok :%s\n", cigarVariant)
}
