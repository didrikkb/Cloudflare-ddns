package Tokenizer

import (
	"fmt"
)

type TokenType int

const (
	Mail TokenType = iota
	Key
	Zone
	Name
	Tp
	Ttl
	End
	Record
	Proxied
	Comment
	Done
)

type Token struct {
	Val string
	Tp  TokenType
}

func (t Token) Print() {
	fmt.Printf("Type: %s, Val: %s\n", TokenTypeToString(t.Tp), t.Val)
}

func TokenTypeToString(tt TokenType) string {
	var res string

	switch tt {
	case Mail:
		res = "Mail"
	case Key:
		res = "Key"
	case Zone:
		res = "Zone"
	case Name:
		res = "Name"
	case Tp:
		res = "Type"
	case Ttl:
		res = "TTL"
	case End:
		res = "End"
	case Record:
		res = "Record"
	case Proxied:
		res = "Proxied"
	case Comment:
		res = "Comment"
	case Done:
		res = "Done"
	}

	return res
}
