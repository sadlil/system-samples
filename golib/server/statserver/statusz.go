package statserver

import (
	"bytes"
	"net/http"
	"os"
	"runtime"
	"strings"
)

// HealthzHandler is an HTTP handler function that returns a 200 OK response to
// indicate the health of the service. It sets the content type to "text/plain"
// and writes the string "ok" to the response writer.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

// StatuszHandler returns a http.HandlerFunc that writes the Go version, Go OS
// and environment variables to an http.ResponseWriter in the HTTP response.
// The function takes a bytes.Buffer as input, and prefix anything the buffer contains.
// The http.ResponseWriter is then set to have a content type of "text/plain", and
// a status code of http.StatusOK is written to it. Finally, the contents of the buffer
// are written to the http.ResponseWriter using the w.Write method.
func StatuszHandler(buf bytes.Buffer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buf.WriteString("\nGo Version: " + runtime.Version())
		buf.WriteString("\nGo OS: " + runtime.GOOS)
		buf.WriteString("\n\nEnvironment Vars:\n")
		for _, e := range os.Environ() {
			buf.WriteString(e)
			buf.WriteString("\n")
		}

		buf.WriteString("Commandline:\n")
		buf.WriteString(strings.Join(os.Args, "\n"))

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(buf.Bytes())
	}
}
