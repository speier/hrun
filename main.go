package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	filename = flag.String("f", "", "filename")
	methods  = map[string]interface{}{
		"HEADER": setHeader,
		"GET":    httpGet,
		"POST":   httpPost,
	}
	header  = http.Header{}
	cookies = []*http.Cookie{}
)

func init() {
	flag.Parse()
}

func main() {
	// default headers
	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json")

	// run script
	vm := NewInterpreter(methods)
	_, err := vm.RunFile(*filename)
	if err != nil {
		println(err.Error())
	}
}

func setHeader(args ...interface{}) {
	for _, a := range args {
		s, ok := a.(string)
		if ok {
			h := strings.Split(s, ":")
			if len(h) == 2 {
				k, v := h[0], h[1]
				header.Set(k, v)
			}
		}
	}
}

func httpGet(host string, args ...interface{}) {
	res, err := httpreq(http.MethodGet, host, args...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}

func httpPost(host string, args ...interface{}) {
	res, err := httpreq(http.MethodPost, host, args...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}

func httpreq(method string, host string, args ...interface{}) (string, error) {
	payload := &bytes.Buffer{}
	for _, a := range args {
		// headers
		s, ok := a.(string)
		if ok {
			h := strings.Split(s, ":")
			if len(h) == 2 {
				k, v := h[0], h[1]
				header.Set(k, v)
			}
		}
		// payload
		m, ok := a.(map[string]interface{})
		if ok {
			b, err := json.Marshal(m)
			if err != nil {
				return "", err
			}
			_, err = payload.Write(b)
			if err != nil {
				return "", err
			}
		}
	}

	req, err := http.NewRequest(method, host, payload)
	if err != nil {
		return "", err
	}

	// set request headers and cookies
	req.Header = header
	for _, c := range cookies {
		req.AddCookie(c)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// save response cookies
	for _, c := range res.Cookies() {
		cookies = append(cookies, c)
	}

	return string(body), nil
}
