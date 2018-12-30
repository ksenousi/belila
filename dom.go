package main

import "container/list"

type node struct {
	children    *list.List
	text        string
	elementData elementData
}

type elementData struct {
	tagName    string
	attributes attrMap
}

type attrMap = map[string]string

func text(data string) node {
	return node{text: data}
}

func elem(name string, attrs attrMap, children *list.List) node {
	elementData := elementData{tagName: name, attributes: attrs}
	return node{children: children, elementData: elementData}
}
