package main

import (
	"container/list"
	"fmt"
)

func main() {
	textNode := text("My first Text Node!\n")
	println(textNode.text)

	elemNode := elem("My first Element Node!\n", nil, list.New())
	println(elemNode.elementData.tagName)

	parser := Parser{0, []rune("汉Hello字")}
	fmt.Printf("\n%t\n", parser.startsWith("汉"))
	println(string(parser.consumeChar()))
	println(string(parser.consumeChar()))
	fmt.Printf("%t\n", parser.startsWith("ello"))
	parser.consumeWhitespace()
	println(parser.pos)
	fmt.Printf("%t\n", parser.eof())

}
