package main

import (
	"EnigmaGo"
	"syscall/js"
)

var GoVars map[string]string

func main() {
	c := make(chan bool)
	
	println("WASM Go Initialized")

	GoVars = make(map[string]string)
	GoVars["bigtext"] = "If you can't create it, you don't understand it"
	GoVars["myvar"] = "value of myvar"
	GoVars["int1"] = "2021"

	//1. Adding an <h1> element in the HTML document
	document := js.Global().Get("document")
	// p := document.Call("createElement", "div")

	// var buff bytes.Buffer
	// for i := 0; i < 10; i++ {
	// 	buff.WriteString("Hey " + fmt.Sprint(i) + " ")
	// }

	// p.Set("innerHTML", "Hello from Golang! <h1> Hey there</h1> "+buff.String())

	wasmbody := document.Get("body").Get("children").Get("wasmbody")

	hbody := wasmbody.Get("innerHTML").String()
	// println(hbody)

	rbody := Render(hbody)
	// println(rbody)

	wasmbody.Set("innerHTML", rbody)

	// wasmbody.Call("appendChild", p)

	//2. Exposing go functions/values in javascript variables.
	js.Global().Set("goVar", "I am a variable set from Go")
	js.Global().Set("sayHello", js.FuncOf(sayHello))

	go RunTests()

	//3. This channel will prevent the goprog ram to exit
	<-c
}

func RunTests() {

	println("Render test")
	println(Render("<div id='wasmbody'>    <h1> Var is <govar  var=\"myvar\">    </h1>  </div>"))
	println()

}

func sayHello(this js.Value, inputs []js.Value) interface{} {
	firstArg := inputs[0].String()
	r := "Hi " + firstArg + " from Go!"
	println(r)
	return r
}
