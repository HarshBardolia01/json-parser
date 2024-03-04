package parser

import (
	"example.com/json-parser/lexer"
)

type JsonElementType int
type StackElementType int

const (
	JSON         JsonElementType = iota // has: value
	VALUE                               // has: object, array, string, number, true, false, null
	OBJECT                              // has: {} or { members }
	MEMBERS                             // has: member or more-member
	MEMBER                              // has: string : value
	ARRAY                               // has: [] or [ values ]
	VALUES                              // has: value or more-value
	MORE_MEMBERS                        // , member
	MORE_VALUES                         // , value
)

const (
	TOKEN StackElementType = iota
	JSON_ELEMENT
)

var JSON_ELEMENT_TYPE_NAMES = []string{
	JSON:         "JSON",
	VALUE:        "Value",
	OBJECT:       "Object",
	MEMBERS:      "Members",
	ARRAY:        "Array",
	VALUES:       "Values",
	MORE_MEMBERS: "More-Members",
	MORE_VALUES:  "More-Values",
}

type StackElement struct {
	elementType StackElementType
	value       interface{}
}

type Parser struct {
	tokens   []lexer.Token
	stack    []StackElement
	depth    int
	position int
	isValid  bool
	isEnd    bool
}

func (par *Parser) parse() {
	par.stack = append(par.stack, StackElement{JSON_ELEMENT, JSON})

	for par.position < len(par.tokens) && !par.isEnd {
		curToken := par.tokens[par.position]

		if par.depth > 20 {
			par.unexpextedToken(curToken)
		}

		if len(par.stack) == 0 {
			if curToken.Type == lexer.EOF {
				par.isValid = true
				return
			} else {
				par.unexpextedToken(curToken)
			}
		}

		if par.isEnd {
			break
		}

		curStackPeek := par.getStackPeekElement()

		switch curStackPeek.elementType {
		case TOKEN:
			if curStackPeek.value == curToken.Type {
				// fmt.Println(curToken, curStackPeek)
				par.depth--
				par.position++
				par.popStack()
			} else {
				par.unexpextedToken(curToken)
			}
		case JSON_ELEMENT:
			par.parseJsonElement(curToken, curStackPeek.value.(JsonElementType))
		}
	}

	par.unexpextedToken(lexer.Token{Type: lexer.EOF, Value: "EOF"})
}

func ParseTokens(tokens []lexer.Token) bool {
	parser := Parser{
		tokens:   tokens,
		stack:    make([]StackElement, 0),
		position: 0,
		isValid:  false,
		isEnd:    false,
	}

	parser.parse()
	return parser.isValid
}
