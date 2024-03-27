package server

import (
	"crypto/tls"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/handler"
	"github.com/shiningrush/droplet"
)

func (s *server) setupAPI() {
	// orchestrator
	droplet.Option.Orchestrator = func(mws []droplet.Middleware) []droplet.Middleware {
		var newMws []droplet.Middleware
		// default middleware order: resp_reshape, auto_input, traffic_log
		// We should put err_transform at second to catch all error
		newMws = append(newMws, mws[0], &handler.ErrorTransformMiddleware{})
		newMws = append(newMws, mws[1:]...)
		return newMws
	}

	// routes
	r := internal.SetUpRouter()

	// HTTP
	addr := net.JoinHostPort(conf.ServerHost, strconv.Itoa(conf.ServerPort))
	s.server = &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  time.Duration(1000) * time.Millisecond,
		WriteTimeout: time.Duration(5000) * time.Millisecond,
	}

	// HTTPS
	if conf.SSLCert != "" && conf.SSLKey != "" {
		addrSSL := net.JoinHostPort(conf.SSLHost, strconv.Itoa(conf.SSLPort))
		s.serverSSL = &http.Server{
			Addr:         addrSSL,
			Handler:      r,
			ReadTimeout:  time.Duration(1000) * time.Millisecond,
			WriteTimeout: time.Duration(5000) * time.Millisecond,
			TLSConfig: &tls.Config{
				// Causes servers to use Go's default ciphersuite preferences,
				// which are tuned to avoid attacks. Does nothing on clients.
				PreferServerCipherSuites: true,
			},
		}
	}
}
