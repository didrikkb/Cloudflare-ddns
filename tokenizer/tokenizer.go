package Tokenizer

import (
	"fmt"
	"os"
	"strings"
)

type Tokenizer struct {
	data    []string
	counter int
}

func Init(file string) *Tokenizer {
	data, err := os.ReadFile(file)
	if err != nil {
		os.Exit(1)
	}

	return &Tokenizer{
		data:    splitInputIntoLines(data),
		counter: 0,
	}
}

func splitInputIntoLines(input []byte) []string {
	data := make([]string, 0)

	for _, l := range strings.Split(string(input), "\n") {
		if len(l) > 0 {
			data = append(data, l)
		}
	}

	return data
}

func (t *Tokenizer) NextToken() Token {
	if len(t.data) <= 0 {
		return Token{
			Tp: Done,
		}
	}

	line := t.data[0]
	t.data = t.data[1:]
	t.counter++

	arr := strings.Split(line, ":")
	tp, val := removeWhitespaces(arr[0]), ""

	if len(arr) > 1 {
		val = arr[1]
	}

	var tokenType TokenType
	var tokenVal string

	switch strings.ToLower(tp) {
	case "mail":
		tokenType = Mail
		tokenVal = removeWhitespaces(val)
	case "key":
		tokenType = Key
		tokenVal = removeWhitespaces(val)
	case "zone":
		tokenType = Zone
		tokenVal = removeWhitespaces(val)
	case "record":
		tokenType = Record
	case "name":
		tokenType = Name
		tokenVal = removeWhitespaces(val)
	case "type":
		tokenType = Tp
		tokenVal = removeWhitespaces(val)
	case "ttl":
		tokenType = Ttl
		tokenVal = removeWhitespaces(val)
	case "proxied":
		tokenType = Proxied
		tokenVal = removeWhitespaces(val)
	case "comment":
		tokenType = Comment
		tokenVal = val
	case "end":
		tokenType = End
	default:
		fmt.Printf("Failed to parse token \"%s\" in config\n", tp)
		os.Exit(1)
	}

	return Token{Val: tokenVal, Tp: tokenType}
}

func removeWhitespaces(input string) string {
	return strings.ReplaceAll(input, " ", "")
}

func (t *Tokenizer) GetTokenNum() int {
	return t.counter
}
