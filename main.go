package main

import (
	Request "cloudflare-ddns/request"
	Tokenizer "cloudflare-ddns/tokenizer"
	"os"
	"path/filepath"
)

func main() {
	changeWorkingDir()
	tokenizer := Tokenizer.Init("./config.conf")
	req := Request.InitRequest(tokenizer)
	req.UpdateDnsRecords()
}

func changeWorkingDir() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dirName := filepath.Dir(exePath)

	err = os.Chdir(dirName)
	if err != nil {
		panic(err)
	}
}
