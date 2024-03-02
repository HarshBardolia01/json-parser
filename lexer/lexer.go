package lexer

type TokenType int
type LexerState int

const (
	OPENING_PARENTHESIS TokenType = iota // (
	CLOSING_PARENTHESIS                  // )
	BEGIN_OBJECT                         // {
	END_OBJECT                           // }
	BEGIN_ARRAY                          // [
	END_ARRAY                            // ]
	NAME_SEPERATOR                       // :
	VALUE_SEPERATOR                      // ,
	NUMBER                               // 123
	MINUS                                // -
	DECIMAL                              // .
	STRING                               // "abc"
	BOOLEAN                              // true-flase
	NULL                                 // null
	EOF                                  // end of file
)

const (
	READING_NORMAL  LexerState = iota // not reading a complex token
	READING_STRING                    // reading a String
	READING_NUMBER                    // reading a Number
	READING_BOOLEAN                   // reading a Boolean
	READING_NULL                      // reading a Null
)

var TOKEN_TYPE_NAMES = []string{
	OPENING_PARENTHESIS: "Opening-Parenthesis",
	CLOSING_PARENTHESIS: "Closing-Parenthesis",
	BEGIN_OBJECT:        "Opening-Braces",
	END_OBJECT:          "Closing-Braces",
	BEGIN_ARRAY:         "Opening-Bracket",
	END_ARRAY:           "Closing-Bracket",
	NAME_SEPERATOR:      "Colon",
	VALUE_SEPERATOR:     "Comma",
	NUMBER:              "Number",
	MINUS:               "-",
	DECIMAL:             ".",
	STRING:              "String",
	BOOLEAN:             "Boolean",
	NULL:                "Null",
	EOF:                 "End-of-File",
}

var STATE_NAMES = []string{
	READING_NORMAL:  "Reading-Normal",
	READING_STRING:  "Reading-String",
	READING_NUMBER:  "Reading-Number",
	READING_BOOLEAN: "Reading-Boolean",
	READING_NULL:    "Reading-Null",
}

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	state      LexerState
	position   int
	buf        []rune
	readTokens []Token
}

func GetTokens(inputJSON string) []Token {
	lexer := Lexer{
		state:      READING_NORMAL,
		position:   0,
		buf:        make([]rune, 0),
		readTokens: make([]Token, 0),
	}

	for _, ch := range inputJSON {
		lexer.readRune(ch)
		lexer.position++
	}

	lexer.readTokens = append(lexer.readTokens, Token{EOF, "EOF"})
	// lexer.printTokens()

	return lexer.readTokens
}
