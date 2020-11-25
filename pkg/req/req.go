package req

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	header  = http.Header{}
	cookies = []*http.Cookie{}

	Methods = map[string]interface{}{
		"HEADER": setHeader,
		"GET":    httpGet,
		"POST":   httpPost,
	}
)

func init() {
	// default headers
	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json")
}

func printres(host string, status int, body string, err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("\n[%d] %s\n%s\n", status, host, body)
}

func httpGet(host string, args ...interface{}) int {
	status, body, err := httpreq(http.MethodGet, host, args...)
	printres(host, status, body, err)
	return status
}

func httpPost(host string, args ...interface{}) int {
	status, body, err := httpreq(http.MethodPost, host, args...)
	printres(host, status, body, err)
	return status
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

func httpreq(method string, host string, args ...interface{}) (statusCode int, body string, err error) {
	setHeader(args...)

	payload := &bytes.Buffer{}
	for _, a := range args {
		// payload
		m, ok := a.(map[string]interface{})
		if ok {
			var b []byte
			b, err = json.Marshal(m)
			if err != nil {
				return
			}
			_, err = payload.Write(b)
			if err != nil {
				return
			}
		}
	}

	req, err := http.NewRequest(method, host, payload)
	if err != nil {
		return
	}

	// set request headers and cookies
	req.Header = header
	for _, c := range cookies {
		req.AddCookie(c)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	var bodyBytes []byte
	bodyBytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	// save response cookies
	for _, c := range res.Cookies() {
		cookies = append(cookies, c)
	}

	return res.StatusCode, string(bodyBytes), nil
}
