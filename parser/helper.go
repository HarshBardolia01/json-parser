package parser

import (
	"fmt"

	"example.com/json-parser/lexer"
)

func (par *Parser) unexpextedToken(token lexer.Token) {
	panic(fmt.Sprintf("Parser Error: Unexpected token %s", lexer.TOKEN_TYPE_NAMES[token.Type]))
}

func (par *Parser) getStackPeekElement() StackElement {
	return par.stack[len(par.stack)-1]
}

func (par *Parser) pushStack(item StackElement) {
	par.stack = append(par.stack, item)
}

func (par *Parser) popStack() StackElement {
	element := par.getStackPeekElement()
	par.stack = par.stack[:len(par.stack)-1]
	return element
}

func (par *Parser) handleJson(token lexer.Token) {
	switch token.Type {
	case lexer.BEGIN_OBJECT, lexer.BEGIN_ARRAY, lexer.STRING, lexer.NUMBER, lexer.BOOLEAN, lexer.NULL:
		par.popStack()
		par.pushStack(StackElement{JSON_ELEMENT, VALUE})
	default:
		par.unexpextedToken(token)
	}
}

func (par *Parser) handleValue(token lexer.Token) {
	switch token.Type {
	case lexer.BEGIN_OBJECT:
		par.popStack()
		par.pushStack(StackElement{JSON_ELEMENT, OBJECT})
	case lexer.BEGIN_ARRAY:
		par.popStack()
		par.pushStack(StackElement{JSON_ELEMENT, ARRAY})
	case lexer.STRING, lexer.NULL, lexer.NUMBER, lexer.BOOLEAN:
		par.popStack()
		par.pushStack(StackElement{TOKEN, token.Type})
	default:
		par.unexpextedToken(token)
	}
}

func (par *Parser) hanldeObject(token lexer.Token) {
	switch token.Type {
	case lexer.BEGIN_OBJECT:
		par.popStack()
		par.pushStack(StackElement{TOKEN, lexer.END_OBJECT})
		par.depth++
		par.pushStack(StackElement{JSON_ELEMENT, MEMBERS})
		par.pushStack(StackElement{TOKEN, lexer.BEGIN_OBJECT})
		par.depth++
	default:
		par.unexpextedToken(token)
	}
}

func (par *Parser) handleArray(token lexer.Token) {
	switch token.Type {
	case lexer.BEGIN_ARRAY:
		par.popStack()
		par.pushStack(StackElement{TOKEN, lexer.END_ARRAY})
		par.depth++
		par.pushStack(StackElement{JSON_ELEMENT, VALUES})
		par.pushStack(StackElement{TOKEN, lexer.BEGIN_ARRAY})
		par.depth++
	default:
		par.unexpextedToken(token)
	}
}

func (par *Parser) handleMembers(token lexer.Token) {
	switch token.Type {
	case lexer.STRING:
		par.popStack()
		par.pushStack(StackElement{JSON_ELEMENT, MORE_MEMBERS})
		par.pushStack(StackElement{JSON_ELEMENT, MEMBER})
	case lexer.END_OBJECT:
		par.popStack()
	default:
		par.unexpextedToken(token)
	}
}

func (par *Parser) handleMemeber(token lexer.Token) {
	switch token.Type {
	case lexer.STRING:
		par.popStack()
		par.pushStack(StackElement{JSON_ELEMENT, VALUE})
		par.pushStack(StackElement{TOKEN, lexer.NAME_SEPERATOR})
		par.depth++
		par.pushStack(StackElement{TOKEN, lexer.STRING})
		par.depth++
	default:
		par.unexpextedToken(token)
	}
}

func (par *Parser) handleMoreMembers(token lexer.Token) {
	switch token.Type {
	case lexer.VALUE_SEPERATOR:
		par.popStack()
		par.pushStack(StackElement{JSON_ELEMENT, MORE_MEMBERS})
		par.pushStack(StackElement{JSON_ELEMENT, MEMBER})
		par.pushStack(StackElement{TOKEN, lexer.VALUE_SEPERATOR})
		par.depth++
	case lexer.END_OBJECT:
		par.popStack()
	default:
		par.unexpextedToken(token)
	}
}

func (par *Parser) handleValues(token lexer.Token) {
	switch token.Type {
	case lexer.BEGIN_OBJECT, lexer.BEGIN_ARRAY, lexer.STRING, lexer.NUMBER, lexer.BOOLEAN, lexer.NULL:
		par.popStack()
		par.pushStack(StackElement{JSON_ELEMENT, MORE_VALUES})
		par.pushStack(StackElement{JSON_ELEMENT, VALUE})
	case lexer.END_ARRAY:
		par.popStack()
	default:
		par.unexpextedToken(token)
	}
}

func (par *Parser) handleMoreValues(token lexer.Token) {
	switch token.Type {
	case lexer.VALUE_SEPERATOR:
		par.popStack()
		par.pushStack(StackElement{JSON_ELEMENT, MORE_VALUES})
		par.pushStack(StackElement{JSON_ELEMENT, VALUE})
		par.pushStack(StackElement{TOKEN, lexer.VALUE_SEPERATOR})
		par.depth++
	case lexer.END_ARRAY:
		par.popStack()
	default:
		par.unexpextedToken(token)
	}
}

func (par *Parser) parseJsonElement(token lexer.Token, jsonElement JsonElementType) {
	switch jsonElement {
	case JSON:
		par.handleJson(token)
	case VALUE:
		par.handleValue(token)
	case OBJECT:
		par.hanldeObject(token)
	case ARRAY:
		par.handleArray(token)
	case MEMBERS:
		par.handleMembers(token)
	case MEMBER:
		par.handleMemeber(token)
	case MORE_MEMBERS:
		par.handleMoreMembers(token)
	case VALUES:
		par.handleValues(token)
	case MORE_VALUES:
		par.handleMoreValues(token)
	}
}
