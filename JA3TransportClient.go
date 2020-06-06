package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"path"
	"github.com/CUCyber/ja3transport"
)

// JA3Response is the struct
type JA3Response struct {
	JA3Hash   string `json:"ja3_hash"`
	JA3       string `json:"ja3"`
	UserAgent string `json:"User-Agent"`
}

func main() {

	httpClient,err := ja3transport.New(ja3transport.SafariAuto)
	if err != nil{
		fmt.Println(err)
		panic(err)
	}

	/* First fetch the JA3 Fingerprint */
	resp, err := httpClient.Get("https://ja3er.com/json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// unmarshal the response
	var ja3Response JA3Response
	err = json.Unmarshal(body, &ja3Response)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	/* Fetch information about the ja3hash*/
	searchURL, _ := url.Parse("https://ja3er.com/search/")
	searchURL.Path = path.Join(searchURL.Path, ja3Response.JA3Hash)

	resp, err = httpClient.Get(searchURL.String())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, body, "", "\t")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(out.String())
}
