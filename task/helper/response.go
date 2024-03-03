package helper

import (
	"bytes"
	"net/http"

	"github.com/joho/godotenv"
)

// response implements http.ResponseWriter
type Response struct {
	buf    *bytes.Buffer
	header http.Header
	status int
}

func (r *Response) Header() http.Header {
	return r.header
}

func (r *Response) WriteHeader(statusCode int) {
	r.status = statusCode
}

func (r *Response) Write(p []byte) (n int, err error) {
	return r.buf.Write(p)
}

func (r *Response) ReadBodyAsMap() (map[string]string, error) {
	return godotenv.Parse(r.buf)
}

func NewResponseWriter() *Response { //nolint:revive
	return &Response{header: map[string][]string{}, buf: &bytes.Buffer{}}
}
