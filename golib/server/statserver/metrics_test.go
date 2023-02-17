package statserver

import (
	"context"
	"net/http"
	"testing"
)

func TestServerRun(t *testing.T) {
	srv := New(
		WithProfiling(true),
		WithPromMetric(true),
		WithMonitoringAddr(":6443"),
	)

	// Run runs the server in a sepreate routine
	srv.Start(context.Background())

	resp, err := http.Get("http://localhost:6443/healthz")
	if err != nil {
		t.Fatalf("http.Get: got %v, expented nil", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("resp.StatusCode: got %v, want 200", resp.StatusCode)
	}

	srv.Stop(context.Background())
}
