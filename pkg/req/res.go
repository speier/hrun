package req

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// minimal http.Response
type Response struct {
	Host             string      `json:"host"`
	Status           string      `json:"status"`     // e.g. "200 OK"
	StatusCode       int         `json:"statusCode"` // e.g. 200
	Proto            string      `json:"proto"`      // e.g. "HTTP/1.0"
	ProtoMajor       int         `json:"protoMajor"` // e.g. 1
	ProtoMinor       int         `json:"protoMinor"` // e.g. 0
	Header           http.Header `json:"header"`
	Body             string      `json:"body"`
	ContentLength    int64       `json:"contentLength"`
	TransferEncoding []string    `json:"transferEncoding"`
	Uncompressed     bool        `json:"uncompressed"`
	Trailer          http.Header `json:"trailer"`
}

func (r *Response) String() string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "\n[%d] %s\n", r.StatusCode, r.Host)
	fmt.Fprintf(sb, "%s\n", r.Body)
	return sb.String()
}

func (r *Response) Map() map[string]interface{} {
	b, err := json.Marshal(&r)
	if err != nil {
		panic(err)
	}
	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	return m
}
