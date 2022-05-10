package main

import "fmt"

const langEnglish = "english"
const langSpanish = "spanish"
const langFranch = "franch"

const EnglishHelloText = "hello"
const SpanishHelloText = "hola"
const FranchHelloText = "bonjure"

const name = "maxyc"
const world = "world"

func main() {
	fmt.Println(hello(name, langEnglish))
}

func hello(name, lang string) string {
	if name == "" {
		name = world
	}

	return GetHelloTextByLang(lang) + " " + name
}

func GetHelloTextByLang(lang string) string {
	switch lang {
	case langSpanish:
		return SpanishHelloText
	case langFranch:
		return FranchHelloText
	}

	return EnglishHelloText
}
