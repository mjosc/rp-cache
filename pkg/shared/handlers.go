package shared

import "net/http"

type ReverseProxy interface {
	http.Handler
}

type APIProxy interface {
	http.Handler
}

type UIProxy interface {
	http.Handler
}
