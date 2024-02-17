package main

import (
	"crypto/tls"
	"flag"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	source string
	target string
)

func init() {
	flag.StringVar(&source, "source", ":8080", "Source to listen on")
	flag.StringVar(&target, "target", "https://localhost", "Target URL")
	flag.Parse()
}
func getTransport() *http.Transport {
	// Create a new TLS configuration that ignores self-signed certificates
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	// Create a new HTTP transport with the custom TLS configuration
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	return transport
}

func main() {
	remote, err := url.Parse(target)
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
      // reset remote address to avoid errors due to proxy detection (e.g. home assistant) 
      r.RemoteAddr = ""
			r.Host = remote.Host
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Transport = getTransport()
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(source, nil)
	if err != nil {
		panic(err)
	}
}
