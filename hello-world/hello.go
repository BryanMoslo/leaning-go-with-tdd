package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "

func main() {
	fmt.Println(Hello("Bryan", ""))
}

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	case portuguese:
		return portugueseHelloPrefix
	default:
		return englishHelloPrefix
	}
}
