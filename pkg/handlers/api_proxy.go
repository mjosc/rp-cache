package handlers

import (
	"github.com/mjosc/rp-cache/pkg/shared"
)

func NewAPIProxy(dst string) (shared.APIProxy, error) {
	config := &ReverseProxyConfig{
		Destination: dst,
	}
	proxy, err := NewReverseProxy(config)
	if err != nil {
		return nil, err
	}
	return proxy, nil
}
