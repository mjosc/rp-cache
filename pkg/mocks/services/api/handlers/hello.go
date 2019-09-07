package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mjosc/rp-cache/pkg/mocks/services/api/shared"
)

func NewHello() shared.Hello {
	return &Hello{}
}

type Hello struct {
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	body, err := json.Marshal(&struct {
		Greeting string
	}{
		Greeting: "Hello",
	})
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Write(body)
}
