package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"

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

			// Do caching stuff here
			//
			// This will only work if using a webserver to host the HTML. Using S3 will require a service
			// object to communicate with the bucket, check the version, and then update the cache accordingly.
			//
			// For the webserver approach, we'd check whether the response was a 304 not modified. If so, retreive
			// the HTML from the body of the response, updated the cache, and forward the response to the client.

			w := NewResponseWriter()

			data := struct {
				Name string
			}{
				Name: "Matt",
			}

			html := string(copy)
			tmpl, err := template.New("test").Parse(html)
			if err != nil {
				return err
			}

			if err := tmpl.Execute(w, &data); err != nil {
				return err
			}

			body := w.Body

			res.Header["Content-Length"] = []string{strconv.Itoa(body.Len())}
			res.Body = ioutil.NopCloser(w.Body)

			return nil
		},
	}

	proxy, err := NewReverseProxy(config)
	if err != nil {
		return nil, err
	}
	return proxy, nil
}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{
		Body: new(bytes.Buffer),
	}
}

type ResponseWriter struct {
	Body *bytes.Buffer
}

func (rw *ResponseWriter) Header() http.Header {
	panic("not implemented")
}

func (rw *ResponseWriter) Write(buf []byte) (int, error) {
	if rw.Body != nil {
		rw.Body.Write(buf)
	}
	return len(buf), nil
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	panic("not implemented")
}
