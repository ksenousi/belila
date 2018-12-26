package main

import "container/list"

type TextNode struct {
	text string
}

type ElementNode struct {
	children    *list.List
	elementData ElementData
}

type ElementData struct {
	tagName    string
	attributes AttrMap
}

type AttrMap = map[string]string

func text(data string) TextNode {
	return TextNode{data}
}

func elem(name string, attrs AttrMap, children *list.List) ElementNode {
	elementData := ElementData{tagName: name, attributes: attrs}
	return ElementNode{children, elementData}
}
