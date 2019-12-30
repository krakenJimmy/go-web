package tools

import (
	"net"

	"go.uber.org/zap"
)

func GetIP(remoteAddr string) string {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	ip, port, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return ""
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return ""
	}

	if port == "" {
		zap.S().Error("port is nil")
	}

	return ip
}
