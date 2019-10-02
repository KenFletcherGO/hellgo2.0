package main

import (
	"fmt"

	"github.com/artytheparty/hellgo2.0/makesure"
)

func main() {
	var codeHolder string
	var languages map = make(map[string]string{"en": "Hello", "es": "Hola", "ge": "Guten tag", "fr": "Bonjour", "ko": "Anyo Haseyo", "ru": "Privet", "ja": "Konnichiwa"})
	var options map := map[string]string{"en": "English", "es": "Espanol", "ge": "German", "fr": "French", "ko": "Korean", "ru": "Russian", "ja": "japanese"}

	fmt.Printf("Welcome please select a language of your choice by typing in the two letter code\n")
	fmt.Println("Layout:")
	fmt.Println("code : Language")
	for k, v := range options {
		fmt.Printf("%s : %s\n", k, v)
	}
	fmt.Scanln(&codeHolder)

	fmt.Printf(makesure.Check(codeHolder, languages) + "Let's Go(lang)!\n")
}
