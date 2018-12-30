package main

import (
	"container/list"
	"regexp"
	"strings"
	"unicode"
)

// Parser Object
type Parser struct {
	pos   int
	input []rune
}

func (prsr Parser) nextChar() rune {
	nextCharacter := prsr.input[prsr.pos]
	return nextCharacter
}

func (prsr Parser) startsWith(s string) bool {
	remainingString := string(prsr.input[prsr.pos:len(prsr.input)])
	return strings.HasPrefix(remainingString, s)
}

func (prsr Parser) eof() bool {
	return prsr.pos >= len(prsr.input)
}

func (prsr *Parser) consumeChar() rune {
	nextCharacter := prsr.input[prsr.pos]
	prsr.pos++
	return nextCharacter
}

type test func(rune) bool

func (prsr *Parser) consumeWhile(t test) string {
	var result strings.Builder
	for !prsr.eof() && t(prsr.nextChar()) {
		result.WriteString(string(prsr.consumeChar()))
	}
	return result.String()
}

func (prsr *Parser) consumeWhitespace() {
	prsr.consumeWhile(unicode.IsSpace)
}

func (prsr *Parser) parseTagName() string {
	return prsr.consumeWhile(func(x rune) bool {
		matched, _ := regexp.MatchString("[A-Za-z0-9]", string(x))
		return matched
	})
}

func (prsr *Parser) parseNode() node {
	if prsr.nextChar() == '<' {
		return prsr.parseElement()
	}
	return prsr.parseText()
}

func (prsr *Parser) parseText() node {
	return text(prsr.consumeWhile(func(x rune) bool {
		return x != '<'
	}))
}

func (prsr *Parser) parseElement() node {
	if prsr.consumeChar() != '<' {
		println("error didn't start with <")
	}

	tagName := prsr.parseTagName()
	attrs := prsr.parseAttributes()

	if prsr.consumeChar() != '>' {
		println("error didn't end with >")
	}

	children := prsr.parseNodes()

	if prsr.consumeChar() != '<' {
		println("error")
	}
	if prsr.consumeChar() != '/' {
		println("error")
	}
	if prsr.parseTagName() != tagName {
		println("error")
	}
	if prsr.consumeChar() != '>' {
		println("error")
	}

	return elem(tagName, attrs, children)
}

func (prsr *Parser) parseAttr() (string, string) {
	name := prsr.parseTagName()
	if prsr.consumeChar() != '=' {
		println("error")
	}
	value := prsr.parseAttrValue()
	return name, value
}

func (prsr *Parser) parseAttrValue() string {
	openQuote := prsr.consumeChar()
	if openQuote != '"' && openQuote != '\'' {
		println("error")
	}
	value := prsr.consumeWhile(func(x rune) bool {
		return x != openQuote
	})

	if prsr.consumeChar() != openQuote {
		println("error")
	}
	return value
}

func (prsr *Parser) parseAttributes() attrMap {
	var attributes map[string]string

	for true {
		prsr.consumeWhitespace()
		if prsr.nextChar() == '>' {
			break
		}
		name, value := prsr.parseAttr()
		attributes[name] = value
	}
	return attributes
}

func (prsr *Parser) parseNodes() *list.List {
	nodes := list.New()

	for true {
		prsr.consumeWhitespace()
		if prsr.eof() || prsr.startsWith("</") {
			break
		}
		nodes.PushBack(prsr.parseNode())
	}
	return nodes
}

func (prsr *Parser) parse(source string) node {
	parser := (Parser{pos: 0, input: []rune(source)})
	nodes := parser.parseNodes()

	if nodes.Len() == 1 {
		return nodes.Front().Value.(node)
	}
	return elem("html", nil, list.New())
}
