package main

import (
	"fmt"

	"github.com/artytheparty/hellgo2.0/makesure"
)

func main() {
	var codeHolder string
	var languages = make(map[string]string)
	languages["en"] = "Hello"
	languages["es"] = "Hola"
	languages["ge"] = "Guten tag"
	languages["fr"] = "Bonjour"
	languages["ko"] = "Anyo Haseyo"
	languages["ru"] = "Privet"
	languages["ja"] = "Konnichiwa"
	languages["ph"] = "Kamusta"
	var options = make(map[string]string)
	options["en"] = "English"
	options["es"] = "Espanol"
	options["ge"] = "German"
	options["fr"] = "French"
	options["ko"] = "Korean"
	options["ru"] = "Russian"
	options["ja"] = "Japanese"
	options["ph"] = "Filipino"

	fmt.Printf("Welcome please select a language of your choice by typing in the two letter code\n")
	fmt.Println("Layout:")
	fmt.Println("code : Language")
	for k, v := range options {
		fmt.Printf("%s : %s\n", k, v)
	}
	fmt.Scanln(&codeHolder)

	fmt.Printf(makesure.Check(codeHolder, languages) + " Let's Go(lang)!\n")
}
