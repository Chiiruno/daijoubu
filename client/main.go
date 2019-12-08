package main

import (
	"github.com/gascore/dom"
	"github.com/gascore/dom/js"
)

func main() {
	println("testo")
	dom.ConsoleError("oh noooooo")
	js.Get("console").Call("warn", "wait maybe I can fix this")
	js.Call("alert", "tehe pero")
}
