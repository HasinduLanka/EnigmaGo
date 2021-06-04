package main

import (
	"fmt"
)

var IsDebug bool = false

func main() {

	fmt.Println("Github.com/hasindulanka/EnigmaGo/EnigmaGo")

	E := MakeEnigmaM3("abc", "hjk", "s", "f", "a")

	if E.Check() {
		println("Enigma internal Check complete")
	}

	tstE := E.InputString("I solemnly swear that i am upto no good, you bastard!")
	println("Encrypted : " + tstE)
	E.PrintRotorOffsets()
	E.ResetOffset()

	tstD := E.InputString(tstE)
	println("Decrypted : " + tstD)
	E.PrintRotorOffsets()
	E.ResetOffset()

}

func MakeEnigmaM3(plugBoardConfigFrom string, plugBoardConfigTo string, Rotor1Offset string, Rotor2Offset string, Rotor3Offset string) Enigma {
	E := Enigma{
		Plugboard: Connector{},
		Rotors:    []Rotor{{}, {}, {}},
		Reflector: Connector{},
	}

	E.Plugboard.MakeConfig(plugBoardConfigFrom, plugBoardConfigTo)
	E.Rotors[0].MakeConfig(EnigmaM3Rotor1Config1, EnigmaM3Rotor1Config2, Rotor1Offset)
	E.Rotors[1].MakeConfig(EnigmaM3Rotor2Config1, EnigmaM3Rotor2Config2, Rotor2Offset)
	E.Rotors[2].MakeConfig(EnigmaM3Rotor3Config1, EnigmaM3Rotor3Config2, Rotor3Offset)
	E.Reflector.MakeConfig(EnigmaM3ReflectorConfig1, EnigmaM3ReflectorConfig2)

	return E
}

const EnigmaM3Rotor1Config1 string = "ekmflgdqvzntowyhxuspaibrcj"
const EnigmaM3Rotor1Config2 string = "r"

const EnigmaM3Rotor2Config1 string = "ajdksiruxblhwtmcqgznpyfvoe"
const EnigmaM3Rotor2Config2 string = "f"

const EnigmaM3Rotor3Config1 string = "bdfhjlcprtxvznyeiwgakmusqo"
const EnigmaM3Rotor3Config2 string = "w"

const EnigmaM3ReflectorConfig1 string = "abcdefgijkmtv"
const EnigmaM3ReflectorConfig2 string = "yruhqslpxnozw"

var IntToAlpha map[int]string
var AlphaToInt map[string]int

const SymbolListSize int = 26

func init() {
	AlphaToInt1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26}

	AlphaToInt = map[string]int{}
	for k, v := range AlphaToInt1 {
		AlphaToInt[k] = v + -1
	}

	IntToAlpha = map[int]string{}
	for k, v := range AlphaToInt {
		IntToAlpha[v] = k
	}

}

// abcdefghijklmnopqrstuvwxyz
