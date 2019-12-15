package main

import "syscall/js"

func main() {
	println("testo")
	js.Global().Get("console").Call("error", "oh noooooo")
	js.Global().Get("console").Call("warn", "wait maybe I can fix this")
	js.Global().Call("alert", "tehe pero")
}
