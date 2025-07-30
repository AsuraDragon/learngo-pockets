package main

import (
	"flag"
	"fmt"
)

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
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, el, fr, he, ur, vi") //#1
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}
