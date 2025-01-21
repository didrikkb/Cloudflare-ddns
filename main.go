package main

import (
	Request "cloudflare-ddns/request"
	Tokenizer "cloudflare-ddns/tokenizer"
)

func main() {
	tokenizer := Tokenizer.Init("./config.conf")
	req := Request.InitRequest(tokenizer)
	req.UpdateDnsRecords()
}
