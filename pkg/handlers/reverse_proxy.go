package handlers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/mjosc/rp-cache/pkg/shared"

	"github.com/mjosc/rp-cache/pkg/utils"
)

type ReverseProxyConfig struct {
	Destination    string
	Director       func(*http.Request)
	ModifyResponse func(*http.Response) error
	ErrorHandler   func(http.ResponseWriter, *http.Request, error)
}

func NewReverseProxy(config *ReverseProxyConfig) (shared.ReverseProxy, error) {
	if config == nil {
		config = &ReverseProxyConfig{}
	}
	proxy := ReverseProxy{}
	director, err := proxy.director(config)
	if err != nil {
		return nil, err
	}
	proxy.Inner = &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: proxy.modifyResponse(config),
		ErrorHandler:   proxy.errorHandler(config),
	}
	return &proxy, nil
}

type ReverseProxy struct {
	Inner *httputil.ReverseProxy
}

func (p *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Inner.ServeHTTP(w, r)
}

func (p *ReverseProxy) director(config *ReverseProxyConfig) (func(*http.Request), error) {
	if config.Director != nil {
		return config.Director, nil
	}
	dst, err := url.ParseRequestURI(config.Destination)
	if err != nil {
		return nil, fmt.Errorf("reverse proxy unable to parse target url: %v", err)
	}
	return func(r *http.Request) {
		path := utils.RemoveDirFromPath(r.URL.Path, 0)
		r.URL.Scheme = dst.Scheme
		r.URL.Host = dst.Host
		r.URL.Path = path
	}, nil
}

func (p *ReverseProxy) modifyResponse(config *ReverseProxyConfig) func(*http.Response) error {
	return config.ModifyResponse
}

func (p *ReverseProxy) errorHandler(config *ReverseProxyConfig) func(http.ResponseWriter, *http.Request, error) {
	return config.ErrorHandler
}
