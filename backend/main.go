package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	LoopbackAddr   = "127.0.0.1:8585" // Zero LAN/WAN exposure
	SecretFirewall = "OmniStream_HMAC_Super_Secret_Key_2026_!#"
)

var bufferPool = sync.Pool{New: func() interface{} { return make([]byte, 65536) }}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/stream", handleSecureStream)
	
	server := &http.Server{
		Addr:    LoopbackAddr,
		Handler: PanicRecoveryMiddleware(LocalhostOnlyMiddleware(mux)),
	}
	log.Println("🚀 Solo Production Daemon Running on", LoopbackAddr)
	server.ListenAndServe()
}

func LocalhostOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RemoteAddr[:9] != "127.0.0.1" && r.RemoteAddr[:5] != "[::1]" {
			http.Error(w, "🚫 WAN Access Blocked", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil { log.Printf("🛡️ Recovered: %v", rec) }
		}()
		next.ServeHTTP(w, r)
	})
}

func verifyHMAC(trackID, ts, sig string) bool {
	mac := hmac.New(sha256.New, []byte(SecretFirewall))
	mac.Write([]byte(trackID + ":" + ts))
	return hmac.Equal([]byte(sig), []byte(hex.EncodeToString(mac.Sum(nil))))
}

func handleSecureStream(w http.ResponseWriter, r *http.Request) {
	if !verifyHMAC(r.URL.Query().Get("id"), r.URL.Query().Get("ts"), r.URL.Query().Get("sig")) {
		http.Error(w, "🚫 Cryptographic Firewall Verification Failed", http.StatusUnauthorized)
		return
	}
	req, _ := http.NewRequest("GET", r.URL.Query().Get("src"), nil)
	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil { http.Error(w, "Upstream Timeout", 502); return }
	defer resp.Body.Close()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	buf := bufferPool.Get().([]byte)
	defer bufferPool.Put(buf)
	io.CopyBuffer(w, resp.Body, buf)
}