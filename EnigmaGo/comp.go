package main

import "strings"

type Enigma struct {
	Plugboard Connector
	Rotors    []Rotor
	Reflector Connector
}

var EnigmaReplacer *strings.Replacer = strings.NewReplacer(" ", "x")

func (E *Enigma) InputString(s string) string {

	rs := []rune(EnigmaReplacer.Replace(strings.ToLower(s)))
	Es := make([]rune, len(rs))

	for i, c := range rs {
		ro := []rune(E.InputCharact(string(c), true))

		if len(ro) == 0 {
			Es[i] = '-'
		} else {
			Es[i] = ro[0]
		}

	}
	return string(Es)
}

func (E *Enigma) InputCharact(s string, ShouldAdvance bool) string {
	si, found := AlphaToInt[s]
	if !found {
		si = SymbolListSize - 1
	}
	return IntToAlpha[E.InputInt(si, ShouldAdvance)]
}

func (E *Enigma) InputInt(inp int, ShouldAdvance bool) int {
	if IsDebug {
		if ShouldAdvance {
			print(">> ")
		} else {
			print("<< ")
		}
	}

	if IsDebug {
		print("Input " + IntToAlpha[inp])
	}

	inp = E.Plugboard.Forward(inp)
	if IsDebug {
		print(" -> " + IntToAlpha[inp])
	}

	lenrots := len(E.Rotors)
	// ORots := make([]int, lenrots)

	for i := lenrots - 1; i > -1; i-- {
		R := &E.Rotors[i]
		inp = R.Forward(inp)
		if IsDebug {
			print(" -> " + IntToAlpha[inp])
		}
		// ORots[i] = inp

	}

	inp = E.Reflector.Forward(inp)
	if IsDebug {
		print(" ->>- " + IntToAlpha[inp])
	}

	// ORotsRev := make([]int, lenrots)

	for i := 0; i < lenrots; i++ {
		R := &E.Rotors[i]
		inp = R.Backward(inp)
		if IsDebug {
			print(" -> " + IntToAlpha[inp])
		}
		// ORotsRev[i] = inp
	}

	inp = E.Plugboard.Backward(inp)
	if IsDebug {
		println(" -> " + IntToAlpha[inp] + " ::")
	}
	if ShouldAdvance {

		E.InputInt(inp, false)
	}

	ChainAdvance := ShouldAdvance

	if ChainAdvance {
		for i := lenrots - 1; i > -1; i-- {
			R := &E.Rotors[i]

			if R.AdvanceOffset() {
				ChainAdvance = false
				break
			}
		}
	}

	return inp
}

func (E *Enigma) ToString() string {

	rt := "\n    Rotors : \n"
	for _, R := range E.Rotors {
		rt += R.ToString() + "\n"
	}

	return "Enigma \n    Plugboard : " + E.Plugboard.ToString() + rt + "\n    Reflector : " + E.Reflector.ToString() + "\n"
}

func (E *Enigma) PrintRotorOffsets() {
	rt := "Rotors : "
	for _, R := range E.Rotors {
		rt += IntToAlpha[R.offset] + "  "
	}

	println(rt)
}

func (E *Enigma) ResetOffset() {
	for i := 0; i < len(E.Rotors); i++ {
		R := &E.Rotors[i]
		R.ResetOffset()
	}

}

func (E *Enigma) Check() bool {
	IsOK := true
	abcd := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(abcd); i++ {
		c := abcd[i : i+1]

		ce := E.InputCharact(c, false)
		cd := E.InputCharact(ce, false)

		if c == cd {
			if IsDebug {
				println("Check " + c + " -> " + ce)
			}
		} else {
			println("Enigma Machine Error : Checking letter failed : Input [" + c + "]  encrypted [" + ce + "]  decrypted [" + cd + "] ")
			IsOK = false
		}
	}

	E.ResetOffset()
	return IsOK
}

// -------------------------------------- //

// Reflector / Plugboard
type Connector struct {
	fwd  map[int]int
	bckw map[int]int
}

func (c *Connector) Forward(pin int) int {
	return c.fwd[pin]
}
func (c *Connector) Backward(pin int) int {
	return c.bckw[pin]
}

func (c *Connector) MakeConfig(A string, B string) {

	c.fwd = make(map[int]int, SymbolListSize)
	c.bckw = make(map[int]int, SymbolListSize)

	for i := 0; i < SymbolListSize; i++ {
		c.fwd[i] = i
		c.bckw[i] = i
	}

	for i := 0; i < len(A); i++ {
		a := AlphaToInt[A[i:i+1]]
		b := AlphaToInt[B[i:i+1]]
		c.fwd[a] = b
		c.fwd[b] = a
		c.bckw[a] = b
		c.bckw[b] = a
	}
}

func (c *Connector) ToString() string {

	l1 := "\n          "
	l2 := "\n          "

	for k, v := range c.fwd {
		l1 += IntToAlpha[k]
		l2 += IntToAlpha[v]
	}

	m1 := "\n          "
	m2 := "\n          "

	for k, v := range c.bckw {
		m1 += IntToAlpha[k]
		m2 += IntToAlpha[v]
	}

	return "    Connector : \n        Fwrd  : " + l1 + l2 + "\n        Bckwd : " + m1 + m2
}

// -------------------------------------- //

type Rotor struct {
	fwd      map[int]int
	bckw     map[int]int
	offset0  int
	offset   int
	advancer int
}

func (c *Rotor) Forward(pin int) int {
	return WrapAround(c.fwd[WrapAround(pin+c.offset)] - c.offset)
}
func (c *Rotor) Backward(pin int) int {
	return WrapAround(c.bckw[WrapAround(pin+c.offset)] - c.offset)
}

func (c *Rotor) MakeConfig(code string, advancer string, offset string) {
	c.offset0 = AlphaToInt[offset]
	c.offset = AlphaToInt[offset]
	c.advancer = AlphaToInt[advancer]

	c.fwd = make(map[int]int, SymbolListSize)
	c.bckw = make(map[int]int, SymbolListSize)

	for i := 0; i < len(code); i++ {
		a := AlphaToInt[code[i:i+1]]
		c.fwd[i] = a
		c.bckw[a] = i
	}

}

func (c *Rotor) ToString() string {

	l1 := "\n          "
	l2 := "\n          "

	for k, v := range c.fwd {
		l1 += IntToAlpha[k]
		l2 += IntToAlpha[v]
	}

	m1 := "\n          "
	m2 := "\n          "

	for k, v := range c.bckw {
		m1 += IntToAlpha[k]
		m2 += IntToAlpha[v]
	}

	return "    Rotor :  Offset " + IntToAlpha[c.offset] + "  Advancer " + IntToAlpha[c.advancer] + "  : \n        Fwrd  : " + l1 + l2 + "\n        Bckwd : " + m1 + m2
}

func (c *Rotor) ResetOffset() {
	c.offset = c.offset0
}

func (R *Rotor) AdvanceOffset() bool {
	R.offset = WrapAround(R.offset + 1)
	return R.offset != R.advancer
}

func WrapAround(x int) int {
	if x < 0 {
		x = SymbolListSize + x
	}

	x = x % SymbolListSize
	return x
}
