package handlers

import (
	"bytes"
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

			// Do caching stuff here
			//
			// This will only work if using a webserver to host the HTML. Using S3 will require a service
			// object to communicate with the bucket, check the version, and then update the cache accordingly.
			//
			// For the webserver approach, we'd check whether the response was a 304 not modified. If so, retreive
			// the HTML from the body of the response, updated the cache, and forward the response to the client.

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
