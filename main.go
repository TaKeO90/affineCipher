package main

import (
	"flag"
	"fmt"
	"log"
)

/*
RULES:

Encryption => E(x) = (ax + b) mod m
Decryption => D(x) = a**-1 (x - b) mod m
*/

// ModuloNumber is set to 26 because we are using English alphabet here.
const ModuloNumber = 26

type Mod int

const (
	nTol Mod = iota
	lTon
)

func bindAlphaNumber(w string, n int, mod Mod) (string, int) {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	switch mod {
	case lTon:
		for i, n := range alphabet {
			if string(n) == w {
				return "", i
			}
		}
		return "", -1
	case nTol:
		return string(alphabet[n]), 0
	default:
		return "", 0
	}
}

type Key struct {
	A int
	B int
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func encProcess(l string, key Key, result *string) {
	var nRes int
	_, num := bindAlphaNumber(l, 0, lTon)
	nRes = (key.A*num + key.B) % ModuloNumber
	w, _ := bindAlphaNumber("", nRes, nTol)
	*result += w
}

func (k Key) Encrypt(word string) (result string) {
	for _, l := range word {
		encProcess(string(l), k, &result)
	}
	return
}

func modMulInverse(n int) int {
	for i := 0; i < ModuloNumber; i++ {
		if ((n%ModuloNumber)*(i%ModuloNumber))%ModuloNumber == 1 {
			return i
		}
	}
	return 1
}

func decProcess(l string, key Key, result *string) {
	_, num := bindAlphaNumber(l, 0, lTon)
	dRes := modMulInverse(key.A) * (num - key.B) % ModuloNumber
	str, _ := bindAlphaNumber("", dRes, nTol)
	*result += str
}

func (k Key) Decrypt(word string) (result string) {
	for _, l := range word {
		decProcess(string(l), k, &result)
	}
	return
}

func main() {
	a := flag.Int("a", 0, "key a")
	b := flag.Int("b", 0, "key b")

	enc := flag.Bool("enc", false, "set enc to false to decrypt and vice versa")

	word := flag.String("word", "", "word to encrypt")

	flag.Parse()

	if *a != 0 && *b != 0 && *word != "" {
		key := Key{
			A: *a,
			B: *b,
		}
		if *enc {
			fmt.Println(key.Encrypt(*word))
		} else {
			fmt.Println(key.Decrypt(*word))
		}
	} else {
		flag.PrintDefaults()
	}
}
