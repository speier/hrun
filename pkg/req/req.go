package req

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

var (
	header  = http.Header{}
	cookies = []*http.Cookie{}

	Methods = map[string]interface{}{
		"header": setHeader,
		"GET":    httpGet,
		"POST":   httpPost,
	}
)

func init() {
	// default headers (can be override from script)
	header.Set("User-Agent", "hrun/0.1")
	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json")
}

func printres(res *Response, err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}

func httpGet(host string, args ...interface{}) interface{} {
	res, err := httpreq(http.MethodGet, host, args...)
	printres(res, err)
	return res.Map()
}

func httpPost(host string, args ...interface{}) interface{} {
	res, err := httpreq(http.MethodPost, host, args...)
	printres(res, err)
	return res.Map()
}

func httpreq(method string, host string, args ...interface{}) (*Response, error) {
	setHeader(args...)

	payload, err := getPayload(args...)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, host, payload)
	if err != nil {
		return nil, err
	}

	// set request headers and cookies
	req.Header = header
	for _, c := range cookies {
		req.AddCookie(c)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var bodyBytes []byte
	bodyBytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// save response cookies
	for _, c := range res.Cookies() {
		cookies = append(cookies, c)
	}

	return &Response{
		Host:             host,
		Status:           res.Status,
		StatusCode:       res.StatusCode,
		Proto:            res.Proto,
		ProtoMajor:       res.ProtoMajor,
		ProtoMinor:       res.ProtoMinor,
		Header:           res.Header,
		Body:             string(bodyBytes),
		ContentLength:    res.ContentLength,
		TransferEncoding: res.TransferEncoding,
		Uncompressed:     res.Uncompressed,
		Trailer:          res.Trailer,
	}, nil
}

func setHeader(args ...interface{}) {
	for _, a := range args {
		s, ok := a.(string)
		if ok {
			h := strings.Split(s, ":")
			if len(h) == 2 {
				k, v := h[0], h[1]
				header.Set(strings.TrimSpace(k), strings.TrimSpace(v))
			}
		}
	}
}

func getPayload(args ...interface{}) (io.Reader, error) {
	payload := &bytes.Buffer{}

	for _, a := range args {
		m, ok := a.(map[string]interface{})
		if ok {
			if strings.Contains(header.Get("Content-Type"), "multipart/form-data") {
				// multipart
				w := multipart.NewWriter(payload)
				err := writeMultipartPayload(m, w)
				if err != nil {
					return nil, err
				}
			} else {
				// json (default)
				err := writeJsonPayload(m, payload)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return payload, nil
}

func writeMultipartPayload(m map[string]interface{}, w *multipart.Writer) error {
	for k, v := range m {
		s := fmt.Sprintf("%v", v)
		// file
		if strings.HasPrefix(s, "@") {
			s = s[1:]
			// read file
			f, err := os.Open(s)
			if err != nil {
				return err
			}
			defer f.Close()
			// copy file to form field part
			fp, err := w.CreateFormFile(k, f.Name())
			_, err = io.Copy(fp, f)
			if err != nil {
				return err
			}
		} else {
			err := w.WriteField(k, s)
			if err != nil {
				return err
			}
		}
	}
	header.Set("Content-Type", w.FormDataContentType())
	return w.Close()
}

func writeJsonPayload(m map[string]interface{}, w io.Writer) error {
	var b []byte
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
