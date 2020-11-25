package req

import (
	"fmt"
	"net/http"
	"strings"
)

// minimal http.Response
type Response struct {
	Host             string
	Status           string // e.g. "200 OK"
	StatusCode       int    // e.g. 200
	Proto            string // e.g. "HTTP/1.0"
	ProtoMajor       int    // e.g. 1
	ProtoMinor       int    // e.g. 0
	Header           http.Header
	Body             string
	ContentLength    int64
	TransferEncoding []string
	Uncompressed     bool
	Trailer          http.Header
}

func (r *Response) String() string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "\n[%d] %s\n", r.StatusCode, r.Host)
	fmt.Fprintf(sb, "%s\n", r.Body)
	return sb.String()
}
