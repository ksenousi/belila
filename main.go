package main

import (
	"fmt"
	"container/list"
)

func main() {
	textNode := text("My first Text Node!\n")
	fmt.Printf(textNode.text)

	elemNode := elem("My first Element Node!\n", nil, list.New())
	fmt.Printf(elemNode.elementData.tagName)
}
