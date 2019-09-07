package handlers

import (
	"net/http"

	"github.com/mjosc/rp-cache/pkg/mocks/services/ui/shared"
)

func NewHello() shared.Hello {
	return &Hello{}
}

type Hello struct {
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write([]byte("<div>Hello, {{ .Name }}</div>"))
}
