package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

// getClientIPInfo extracts all IP-related information from the HTTP request
func getClientIPInfo(r *http.Request) map[string]string {
	info := make(map[string]string)

	// Get X-Forwarded-For header
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		info["X-Forwarded-For"] = xff
	} else {
		info["X-Forwarded-For"] = "not present"
	}

	// Get X-Real-IP header
	xri := r.Header.Get("X-Real-IP")
	if xri != "" {
		info["X-Real-IP"] = xri
	} else {
		info["X-Real-IP"] = "not present"
	}

	// Extract IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		info["RemoteAddr"] = r.RemoteAddr
	} else {
		info["RemoteAddr"] = ip
	}

	return info
}

// ipHandler handles HTTP requests and returns detailed IP information
func ipHandler(w http.ResponseWriter, r *http.Request) {
	ipInfo := getClientIPInfo(r)

	// Return as plain text with all IP information
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "X-Forwarded-For: %s\n", ipInfo["X-Forwarded-For"])
	fmt.Fprintf(w, "X-Real-IP: %s\n", ipInfo["X-Real-IP"])
	fmt.Fprintf(w, "RemoteAddr: %s\n", ipInfo["RemoteAddr"])

	// Log access
	log.Printf("%s %s %s", r.Method, r.URL.Path, ipInfo["RemoteAddr"])
}

func main() {
	port := "8080"

	// Allow port specification via command line argument
	if len(os.Args) > 1 {
		if _, err := strconv.Atoi(os.Args[1]); err == nil {
			port = os.Args[1]
		}
	}

	// Set up route handler
	http.HandleFunc("/", ipHandler)

	addr := ":" + port
	fmt.Printf("Starting IP server on port %s...\n", port)
	fmt.Printf("Test commands:\n")
	fmt.Printf("  curl http://localhost:%s/\n", port)
	fmt.Printf("  curl --socks5 localhost:1080 http://localhost:%s/\n", port)
	fmt.Printf("Press Ctrl+C to stop\n\n")

	log.Fatal(http.ListenAndServe(addr, nil))
}
