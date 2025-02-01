package Request

import (
	Tokenizer "cloudflare-ddns/tokenizer"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	Comment string
	Name    string
	Tp      string
	Content string
	ID      string
	TTL     int
	Exists  bool
	Proxied bool
}

func InitRecord(tokenizer *Tokenizer.Tokenizer) Record {
	rec := Record{Comment: "", TTL: 300, Proxied: false, ID: "", Exists: false, Tp: ""}
	done := false

	for done == false {
		t := tokenizer.NextToken()

		switch t.Tp {
		case Tokenizer.Comment:
			rec.Comment = t.Val
		case Tokenizer.Name:
			rec.Name = t.Val
		case Tokenizer.Tp:
			if strings.ToLower(t.Val) == "aaaa" {
				rec.Tp = "AAAA"
				rec.Content = getIP(IPv6)
			} else {
				rec.Tp = "A"
				rec.Content = getIP(IPv4)
			}
		case Tokenizer.Ttl:
			v, err := strconv.Atoi(t.Val)
			if err != nil {
				panic(err)
			}
			rec.TTL = v
		case Tokenizer.Proxied:
			rec.Proxied = t.Val == "true"
		case Tokenizer.End:
			done = true
		default:
			fmt.Printf("Token \"%s\" was unexpected at position %d in record\n", Tokenizer.TokenTypeToString(t.Tp), tokenizer.GetTokenNum())
			os.Exit(1)
		}
	}

	if !rec.isValid() {
		fmt.Println("Missing element in record")
		os.Exit(1)
	}

	return rec
}

func (rec *Record) Print() {
	fmt.Printf("\nRecord:\n")
	fmt.Printf("Comment: %s\nName: %s\nType: %s\nTTL: %d\nProxied: %s\nContent: %s\nID: %s\n", rec.Comment, rec.Name, rec.Tp, rec.TTL, rec.Proxied, rec.Content, rec.ID)
}

func (rec *Record) isValid() bool {
	return len(rec.Name) > 0 && len(rec.Tp) > 0
}

type IPv int

const (
	IPv4 IPv = iota
	IPv6
)

var ipv4 string = ""
var ipv6 string = ""

func getIP(v IPv) string {
	var url string
	var res string

	switch v {
	case IPv4:
		if ipv4 != "" {
			return ipv4
		}
		url = "https://ipinfo.io/ip"
	case IPv6:
		if ipv6 != "" {
			return ipv6
		}
		url = "https://v6.ipinfo.io/ip"
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	buff := ReadToBuffer(resp.Body)

	switch v {
	case IPv4:
		ipv4 = string(buff)
		res = ipv4
	case IPv6:
		ipv6 = string(buff)
		res = ipv6
	}

	return res
}
