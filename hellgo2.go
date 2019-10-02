package main

import (
	"fmt"
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
	var options = make(map[string]string)
	options["en"] = "English"
	options["es"] = "Espanol"
	options["ge"] = "German"
	options["fr"] = "French"
	options["ko"] = "Korean"
	options["ru"] = "Russian"
	options["ja"] = "Japanese"

	fmt.Printf("Welcome please select a language of your choice by typing in the two letter code\n")
	fmt.Println("Layout:")
	fmt.Println("code : Language")
	for k, v := range options {
		fmt.Printf("%s : %s\n", k, v)
	}
	fmt.Scanln(&codeHolder)

	fmt.Printf(Check(codeHolder, languages) + " Let's Go(lang)!\n")
}
func Check(s string, f map[string]string) string {
	var holder string
	holder = f[s]
	//if its not the key then the value would be
	//empty remeber this
	if holder == "" {
		holder = "Yo"
	}
	/*
		for k, v := range f {
			if len(s) == 2 && k == s {
				holder = v
			} else {
				holder = "Yo"
			}
		}
	*/
	return holder
}
