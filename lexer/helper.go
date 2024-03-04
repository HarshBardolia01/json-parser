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
	lex.isError = true
	// panic(fmt.Sprintf("Lexer Error: Unexpected character '%c' at row %d and col %d in %s State.", ch, lex.row, lex.col, lex.getStateName()))
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
		lex.readTokens = append(lex.readTokens, Token{END_ARRAY, "]"})
	case ch == ':':
		lex.readTokens = append(lex.readTokens, Token{NAME_SEPERATOR, ":"})
	case ch == ',':
		lex.readTokens = append(lex.readTokens, Token{VALUE_SEPERATOR, ","})
	case ch == '"':
		if len(lex.readTokens) == 0 {
			lex.unexpectedCharacter(ch)
		}
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

	if word[size-1] != '"' {
		return false
	}

	isEnd := true
	pos := size - 2

	for pos >= 0 {
		if word[pos] == '\\' {
			isEnd = !isEnd
		} else {
			break
		}
		pos--
	}

	return isEnd
}

func isHex(word []rune) bool {
	if len(word) != 4 {
		return false
	}

	for _, ch := range word {
		if !(unicode.IsNumber(ch) || unicode.IsLetter(ch)) {
			return false
		}
	}

	return true
}

func isValidString(word []rune) (bool, int) {
	size := len(word)
	pos := 0

	for pos < size {
		if word[pos] == '\\' {
			if pos+1 >= size {
				return false, pos
			}

			ch := word[pos+1]

			if ch == '"' || ch == '\\' || ch == '/' || ch == 'b' || ch == 'f' || ch == 'n' || ch == 'r' || ch == 't' {
				pos += 2
				continue
			}

			if ch == 'u' && pos+4 < size && isHex(word[pos+1:pos+5]) {
				pos += 5
				continue
			}

			return false, pos
		}

		pos++
	}

	return true, -1
}

func (lex *Lexer) readString(ch rune) {
	lex.buf = append(lex.buf, ch)

	if isStringEnd(lex.buf) {
		isValid, pos := isValidString(lex.buf)

		if !isValid {
			lex.unexpectedCharacter(lex.buf[pos])
		}

		lex.appendToken(STRING, string(lex.buf))
		lex.buf = make([]rune, 0)
		lex.state = READING_NORMAL
	}
}

func isValidInteger(num []rune) bool {
	size := len(num)

	if size == 0 {
		return true
	}

	if size == 1 {
		return unicode.IsDigit(num[0])
	}

	if num[0] == '0' {
		return false
	}

	cntDigits := 0

	for _, ch := range num {
		if unicode.IsDigit(ch) {
			cntDigits++
		}
	}

	if cntDigits == size {
		return true
	}

	if cntDigits == size-1 && num[0] == '-' && num[1] != '0' {
		return true
	}

	return false
}

func isValidFraction(num []rune) bool {
	cntDigits := 0

	for _, ch := range num {
		if unicode.IsDigit(ch) {
			cntDigits++
		}
	}

	return len(num) == cntDigits
}

func isValidExponent(num []rune) bool {
	size := len(num)

	if size == 0 {
		return false
	}

	cntDigits := 0

	for _, ch := range num {
		if unicode.IsDigit(ch) {
			cntDigits++
		}
	}

	if size == cntDigits {
		return true
	}

	startsWithSign := num[0] == '+' || num[0] == '-'

	if cntDigits+1 == size && cntDigits != 0 && startsWithSign {
		return true
	}

	return false
}

func isValidNumber(num []rune) bool {
	dotPos := -1
	expPos := -1
	cntExp := 0
	cntDot := 0

	for ind, ch := range num {
		if ch == '.' {
			dotPos = ind
			cntDot++
		}

		if ch == 'e' || ch == 'E' {
			expPos = ind
			cntExp++
		}
	}

	if cntDot > 1 || cntExp > 1 {
		return false
	}

	var integer, fraction, exponent []rune

	if dotPos != -1 {
		integer = num[0:dotPos]
	} else {
		if expPos != -1 {
			integer = num[0:expPos]
		} else {
			integer = num
		}
	}

	if dotPos != -1 {
		if expPos != -1 {
			fraction = num[dotPos+1 : expPos]
		} else {
			fraction = num[dotPos+1:]
		}
	}

	if expPos != -1 {
		exponent = num[expPos+1:]
	}

	if !isValidInteger(integer) {
		return false
	}

	if dotPos != -1 && !isValidFraction(fraction) {
		return false
	}

	if expPos != -1 && !isValidExponent(exponent) {
		return false
	}

	return true
}

func (lex *Lexer) readNumber(ch rune) {
	isEnd := (ch == ',' || ch == '}' || ch == ']' || unicode.IsSpace(ch))
	isValid := (ch == '.' || ch == '-' || ch == '+' || ch == 'e' || ch == 'E')

	if unicode.IsDigit(ch) {
		lex.buf = append(lex.buf, ch)
	} else if isValid {
		lex.buf = append(lex.buf, ch)
	} else if isEnd {
		if isValidNumber(lex.buf) {
			lex.appendToken(NUMBER, string(lex.buf))
			lex.buf = make([]rune, 0)
			lex.state = READING_NORMAL
			lex.readNormal(ch)
		} else {
			lex.unexpectedCharacter(ch)
		}
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
		if ch == '\n' || ch == '\t' {
			lex.unexpectedCharacter(ch)
		}
		lex.readString(ch)
	case lex.state == READING_NUMBER:
		lex.readNumber(ch)
	case lex.state == READING_BOOLEAN:
		lex.readBoolean(ch)
	case lex.state == READING_NULL:
		lex.readNull(ch)
	}
}
