package handlers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/mjosc/rp-cache/pkg/shared"
)

func NewUIProxy(dst string) (shared.UIProxy, error) {
	config := &ReverseProxyConfig{
		Destination: dst,
		ModifyResponse: func(res *http.Response) error {

			copy, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return nil
			}

			// Do caching stuff here.

			str := string(copy)
			updated := strings.Replace(str, "{{ .Name }}", "Matt", 1)
			data := []byte(updated)

			res.Header["Content-Length"] = []string{strconv.Itoa(len(data))}
			res.Body = ioutil.NopCloser(bytes.NewBuffer(data))

			return nil
		},
	}

	proxy, err := NewReverseProxy(config)
	if err != nil {
		return nil, err
	}
	return proxy, nil
}

func NewResponseWriter() http.ResponseWriter {
	return &ResponseWriter{}
}

type ResponseWriter struct {
	Body io.ReadCloser
}

func (w *ResponseWriter) Header() http.Header {
	return nil
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	fmt.Println("hello")
	w.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	return len(data), nil
}

func (w *ResponseWriter) WriteHeader(statusCode int) {

}
