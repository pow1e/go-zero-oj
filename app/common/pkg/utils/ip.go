package utils

import (
	"github.com/thinkeridea/go-extend/exnet"
	"net/http"
)

func GetIP(req *http.Request) string {
	ip := ""
	if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	if ip = req.Header.Get("X-Real-Ip"); ip != "" {
		return ip
	}
	if ip = exnet.ClientIP(req); ip != "" {
		return ip
	}
	if ip = exnet.ClientPublicIP(req); ip != "" {
		return ip
	}
	if ip = exnet.RemoteIP(req); ip != "" {
		return ip
	}
	return ""
}
