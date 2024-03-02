package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

func (lex *Lexer) getStateName() string {
	return STATE_NAMES[lex.state]
}

func (lex *Lexer) unexpectedCharacter(ch rune) {
	panic(fmt.Sprintf("Lexer Error: Unexpected character '%c' at position %d in %s State.", ch, lex.position, lex.getStateName()))
}

func (lex *Lexer) printTokens() {
	fmt.Println("Tokens:")

	for _, token := range lex.readTokens {
		fmt.Printf("%-25s: %s\n", TOKEN_TYPE_NAMES[token.Type], token.Value)
	}
}

func (lex *Lexer) appendToken(tokenType TokenType, value string) {
	lex.readTokens = append(lex.readTokens, Token{tokenType, value})
}

func (lex *Lexer) readNormal(ch rune) {
	switch {
	case ch == '(':
		lex.readTokens = append(lex.readTokens, Token{OPENING_PARENTHESIS, "("})
	case ch == ')':
		lex.readTokens = append(lex.readTokens, Token{CLOSING_PARENTHESIS, ")"})
	case ch == '{':
		lex.readTokens = append(lex.readTokens, Token{BEGIN_OBJECT, "{"})
	case ch == '}':
		lex.readTokens = append(lex.readTokens, Token{END_OBJECT, "}"})
	case ch == '[':
		lex.readTokens = append(lex.readTokens, Token{BEGIN_ARRAY, "["})
	case ch == ']':
		lex.readTokens = append(lex.readTokens, Token{END_ARRAY, "["})
	case ch == ':':
		lex.readTokens = append(lex.readTokens, Token{NAME_SEPERATOR, ":"})
	case ch == ',':
		lex.readTokens = append(lex.readTokens, Token{VALUE_SEPERATOR, ","})
	case ch == '"':
		lex.state = READING_STRING
		lex.buf = append(lex.buf, ch)
	case ch == 't' || ch == 'f':
		lex.state = READING_BOOLEAN
		lex.buf = append(lex.buf, ch)
	case unicode.IsDigit(ch) || ch == '-':
		lex.state = READING_NUMBER
		lex.buf = append(lex.buf, ch)
	case ch == 'n':
		lex.state = READING_NULL
		lex.buf = append(lex.buf, ch)
	case unicode.IsSpace(ch):
	default:
		lex.unexpectedCharacter(ch)
	}

}

func isStringEnd(word []rune) bool {
	size := len(word)
	isEnd := true

	if word[size-1] != '"' {
		return false
	}

	for i := size - 2; i >= 0; i-- {
		if word[i] == '\\' {
			isEnd = !isEnd
		} else {
			break
		}
	}

	return isEnd
}

func (lex *Lexer) readString(ch rune) {
	lex.buf = append(lex.buf, ch)

	if isStringEnd(lex.buf) {
		lex.appendToken(STRING, string(lex.buf))
		lex.buf = make([]rune, 0)
		lex.state = READING_NORMAL
	}
}

func (lex *Lexer) readNumber(ch rune) {
	isEnd := (ch == ',' || ch == '}' || ch == ']' || unicode.IsSpace(ch))
	isValid := (ch == '.' || ch == '-' || ch == '+' || ch == 'e' || ch == 'E')

	if unicode.IsDigit(ch) {
		lex.buf = append(lex.buf, ch)
	} else if isValid {
		lex.buf = append(lex.buf, ch)
	} else if isEnd {
		lex.appendToken(NUMBER, string(lex.buf))
		lex.buf = make([]rune, 0)
		lex.state = READING_NORMAL
		lex.readNormal(ch)
	} else {
		lex.unexpectedCharacter(ch)
	}
}

func (lex *Lexer) readBoolean(ch rune) {
	isBooleanRead := string(lex.buf) == "true" || string(lex.buf) == "false"

	if isBooleanRead {
		isEnd := (ch == ',' || ch == '}' || ch == ']' || unicode.IsSpace(ch))

		if isEnd {
			lex.appendToken(BOOLEAN, string(lex.buf))
			lex.buf = make([]rune, 0)
			lex.state = READING_NORMAL
			lex.readNormal(ch)
		} else {
			lex.unexpectedCharacter(ch)
		}
	} else {
		lex.buf = append(lex.buf, ch)
		isInTrue := strings.Contains("true", string(lex.buf))
		isInFalse := strings.Contains("false", string(lex.buf))

		if !(isInFalse || isInTrue) {
			lex.unexpectedCharacter(ch)
		}
	}
}

func (lex *Lexer) readNull(ch rune) {
	isNullRead := string(lex.buf) == "null"

	if isNullRead {
		isEnd := (ch == ',' || ch == '}' || ch == ']' || unicode.IsSpace(ch))

		if isEnd {
			lex.appendToken(NULL, string(lex.buf))
			lex.buf = make([]rune, 0)
			lex.state = READING_NORMAL
			lex.readNormal(ch)
		} else {
			lex.unexpectedCharacter(ch)
		}
	} else {
		lex.buf = append(lex.buf, ch)
		isInNull := strings.Contains("null", string(lex.buf))

		if !isInNull {
			lex.unexpectedCharacter(ch)
		}
	}
}

func (lex *Lexer) readRune(ch rune) {
	switch {
	case lex.state == READING_NORMAL:
		lex.readNormal(ch)
	case lex.state == READING_STRING:
		lex.readString(ch)
	case lex.state == READING_NUMBER:
		lex.readNumber(ch)
	case lex.state == READING_BOOLEAN:
		lex.readBoolean(ch)
	case lex.state == READING_NULL:
		lex.readNull(ch)
	}
}
