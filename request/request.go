package Request

import (
	"bytes"
	Tokenizer "cloudflare-ddns/tokenizer"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Request struct {
	Mail    string
	Key     string
	Zone    string
	Records []Record
}

func InitRequest(tokenizer *Tokenizer.Tokenizer) Request {
	req := Request{}
	done := false

	for done == false {
		t := tokenizer.NextToken()

		switch t.Tp {
		case Tokenizer.Mail:
			req.Mail = t.Val
		case Tokenizer.Key:
			req.Key = t.Val
		case Tokenizer.Zone:
			req.Zone = t.Val
		case Tokenizer.Record:
			req.Records = append(req.Records, InitRecord(tokenizer))
		case Tokenizer.Done:
			done = true
		default:
			fmt.Printf("Failed to parse token: %d in request\n", tokenizer.GetTokenNum())
			os.Exit(1)
		}
	}

	if !req.isValid() {
		fmt.Println("Missing element in request")
		os.Exit(1)
	}

	return req
}

func (req *Request) isValid() bool {
	return len(req.Key) > 0 && len(req.Mail) > 0 && len(req.Zone) > 0 && len(req.Records) > 0
}

func (req *Request) Print() {
	fmt.Printf("Mail: %s\nKey: %s\nZone: %s\n", req.Mail, req.Key, req.Zone)

	for _, rec := range req.Records {
		rec.Print()
	}
}

func (req *Request) getRecords() []ResultItem {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", req.Zone)

	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Set("X-Auth-Email", req.Mail)
	r.Header.Set("X-Auth-Key", req.Key)

	client := http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	records := ""
	buff := make([]byte, 1024)

	for {
		i, err := res.Body.Read(buff)
		if err != nil || i == 0 {
			break
		}

		records += string(buff[:i])
	}

	var response RecordsResponse
	json.Unmarshal([]byte(records), &response)

	return response.Result
}

func (req *Request) FindExistingRecordIds() {
	recordRes := req.getRecords()

	for _, excRec := range recordRes {
		for idx, newRec := range req.Records {
			if newRec.Name == excRec.Name && newRec.Tp == excRec.Type {
				req.Records[idx].ID = excRec.ID
				req.Records[idx].Exists = compareRecords(newRec, excRec)
			}
		}
	}
}

func compareRecords(rec1 Record, rec2 ResultItem) bool {
	if (rec1.Comment == "") != (rec2.Comment == nil) {
		return false
	}
	if rec2.Comment != nil {
		return rec1.Comment == *(rec2.Comment) && rec1.Content == rec2.Content && rec1.Proxied == rec2.Proxied && rec1.TTL == rec2.TTL
	}
	return rec1.Content == rec2.Content && rec1.Proxied == rec2.Proxied && rec1.TTL == rec2.TTL
}

func (req *Request) UpdateDnsRecords() {

	client := http.Client{}
	buff := make([]byte, 4096)

	for _, rec := range req.Records {
		if rec.Exists {
			fmt.Println("Identical record already exist")
			continue
		}
		recordRequest := req.createRequest(rec, rec.ID != "")
		r, err := client.Do(recordRequest)
		if err != nil {
			fmt.Println(err)
			continue
		}

		i, err := r.Body.Read(buff)

		fmt.Println((string(buff[:i])))
	}
}

func (req *Request) createRequest(rec Record, update bool) *http.Request {
	var url string
	var method string

	if update {
		url = fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", req.Zone, rec.ID)
		method = "PATCH"
	} else {
		url = fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", req.Zone)
		method = "POST"
	}

	body := struct {
		Comment string `json:"comment"`
		Content string `json:"content"`
		Name    string `json:"name"`
		Proxied bool   `json:"proxied"`
		TTL     int    `json:"ttl"`
		Type    string `json:"type"`
	}{
		Comment: rec.Comment,
		Content: rec.Content,
		Name:    rec.Name,
		Proxied: rec.Proxied,
		TTL:     rec.TTL,
		Type:    rec.Tp,
	}

	jsonBody, _ := json.Marshal(body)

	r, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Auth-Email", req.Mail)
	r.Header.Set("X-Auth-Key", req.Key)

	return r
}
