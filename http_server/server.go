package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/web", handleRequest)
	http.HandleFunc("/healthz", handleLive)
	http.ListenAndServe(":80", nil)
}

func handleRequest(res http.ResponseWriter, req *http.Request) {
	reWriteRespose(res, req)

	getEnv(res, "VERSION")

	stdOut(res, req, http.StatusOK)
}

func handleLive(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("200"))
}

func reWriteRespose(res http.ResponseWriter, req *http.Request) {
	for i := range req.Header {
		val := ""
		for j, v := range req.Header[i] {
			if j == 0 {
				val = v
				continue
			}
			val += v
		}
		res.Header().Set(i, val)
	}
}

func getEnv(res http.ResponseWriter, key string) {
	res.Header().Add(key, os.Getenv(key))
}

func stdOut(res http.ResponseWriter, req *http.Request, statusCode int) {
	ip := clientIP(req)
	res.WriteHeader(statusCode)
	fmt.Printf("ClientIP:%s   Status Code:%v\n", ip, statusCode)
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
