package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading", err)
		http.Error(w, "Unable to read request", http.StatusBadRequest)
		return
	}

	fmt.Println(w, "Hello %s", string(data))
}
