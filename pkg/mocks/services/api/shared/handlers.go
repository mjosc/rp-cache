package shared

import "net/http"

type Hello interface {
	http.Handler
}
