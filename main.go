package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func main() {
	// 设置路由，当访问根路径时，调用 handleRequest 函数
	http.HandleFunc("/", handleRequest)

	// 设置监听的端口
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// handleRequest 处理入站的 HTTP 请求
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 获取请求者的 IP 地址
	ip := getIPAddress(r)

	// 如果 IP 地址不是私有的，将其返回给请求者
	if !isPrivateIP(ip) {
		fmt.Fprintf(w, "%s", ip)
	} else {
		fmt.Fprint(w, "Private IP addresses are not displayed.")
	}
}

// isPrivateIP 检查 IP 地址是否属于私有地址段
func isPrivateIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
	_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
	_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
	return private24BitBlock.Contains(ip) || private20BitBlock.Contains(ip) || private16BitBlock.Contains(ip)
}

// getIPAddress 从 HTTP 请求中提取 IP 地址
func getIPAddress(r *http.Request) string {
	// 尝试从 X-Forwarded-For 头部获取 IP 地址
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// 取 X-Forwarded-For 中的第一个 IP（可能有多个 IP，由代理添加）
		return strings.Split(xForwardedFor, ",")[0]
	}

	// 否则，直接从远程地址获取 IP（并去除可能附加的端口信息）
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
