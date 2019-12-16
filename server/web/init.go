package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/log/v7"

	"github.com/Chiiruno/daijoubu/server/config"
	"github.com/Chiiruno/daijoubu/server/util"
)

// Init initializes the web server.
func Init() (err error) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	closing := false
	done := make(chan bool, 1)
	c := config.Server.Server
	server := &http.Server{
		Addr:    c.Address,
		Handler: createRouter(),
	}
	// server.RegisterOnShutdown(db.Close)
	prettyAddr := c.Address

	if len(c.Address) != 0 && c.Address[0] == ':' {
		prettyAddr = fmt.Sprintf("127.0.0.1%s", prettyAddr)
	}

	go func() {
		log.Infof("listening on http://%s", prettyAddr)
		err = util.WrapError("error starting web server", server.ListenAndServe())
		log.Infof("no longer listening on http://%s", prettyAddr)
		done <- true
	}()

	for {
		select {
		case s := <-sigs:
			if !closing {
				closing = true
				log.Warnf("shutting down by signal: %s", s)
				server.Shutdown(context.Background())
			}
		case <-done:
			return
		}
	}
}
