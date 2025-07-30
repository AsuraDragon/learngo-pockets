package main

import "fmt"

type language string //#1

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",      // Greek
	"en": "Hello world",        // English
	"fr": "Bonjour le monde",   // French
	"he": "םלוע םולש",          // Hebrew
	"ur": "ﺎﯿﻧد ﻮﻠﯿہ",          // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

func greet(l language) string { //#1
	greeting, ok := phrasebook[l] //#2
	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}
	return greeting
}

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}
