package statserver

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthzHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()

	// Call healthz handler
	HealthzHandler(w, req)

	if w.Code != 200 {
		t.Errorf("w.Code: got %v, want %v", w.Code, 200)
	}

	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("io.ReadAll(Body): got %v, expected nil", err)
	}
	if string(data) != "ok" {
		t.Errorf("Body: got %v, expected Ok", string(data))
	}
}

func TestStatuszHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/statusz", nil)
	w := httptest.NewRecorder()

	buf := bytes.Buffer{}
	buf.WriteString("Hello World!")

	// Call StatuszHandler handler
	StatuszHandler(buf)(w, req)

	if w.Code != 200 {
		t.Errorf("w.Code: got %v, want %v", w.Code, 200)
	}

	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("io.ReadAll(Body): got %v, expected nil", err)
	}

	t.Logf("Body:\n%v", string(data))
	if idx := strings.Index(string(data), "Hello World!"); idx != 0 {
		t.Errorf("Body(Hello World!): got %v, expected 0", idx)
	}

	if ok := strings.Contains(string(data), "Go Version"); !ok {
		t.Errorf("Body(Go Version): got %v, expected true", ok)
	}
}
