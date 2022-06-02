package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

type Mux struct {
	http.ServeMux
}

func NewMux() *Mux { return &Mux{} }

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for i := range r.Header {
		val := ""
		for j, v := range r.Header[i] {
			if j == 0 {
				val = v
				continue
			}
			val += v
		}
		w.Header().Set(i, val)
	}
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	stdOut(r)

	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h, _ := m.Handler(r)
	h.ServeHTTP(w, r)
}

func stdOut(req *http.Request) {
	ip := clientIP(req)
	fmt.Printf("ClientIP:%s \n", ip)
}

func clientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {

		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
