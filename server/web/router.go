package web

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"runtime/debug"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/go-playground/log/v7"

	"github.com/Chiiruno/daijoubu/server/config"
	// "github.com/Chiiruno/daijoubu/server/imager"
	"github.com/Chiiruno/daijoubu/server/util"
)

var (
	healthCheckMsg = []byte("God's in His heaven, all's right with the world")
	webRoot        = "www" // Used for overriding during tests.
)

// Create the monolithic router for routing HTTP requests. Separated into own
// function for easier testability.
func createRouter() http.Handler {
	r := httptreemux.NewContextMux()

	r.NotFoundHandler = func(w http.ResponseWriter, _ *http.Request) {
		text404(w)
	}

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		http.Error(w, fmt.Sprintf("500 %s", err), 500)
		ip, ipErr := util.GetIP(r)

		if ipErr != nil {
			ip = net.IPv4zero
		}

		log.Errorf("server: %s: %#v\n%s\n", ip, err, debug.Stack())
	}

	r.GET("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		buf.WriteString("User-agent: *\n")

		if config.Get().DisableRobots {
			buf.WriteString("Disallow: /\n")
		}

		w.Header().Set("Content-Type", "text/plain")
		buf.WriteTo(w)
	})

	api := r.NewGroup("/api")
	api.GET("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Write(healthCheckMsg)
	})
	/* api.GET("/socket", func(w http.ResponseWriter, r *http.Request) {
		httpError(w, r, websockets.Handle(w, r))
	}) */

	// All upload images
	// api.POST("/upload", imager.NewImageUpload)
	// api.POST("/upload-hash", imager.UploadImageHash)

	/* assets := r.NewGroup("/assets")
	assets.GET("/media/*path", serveImages)
	assets.GET("/*path", serveAssets) */
	return r
}
