package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEndpoints(t *testing.T) {
	srv := httptest.NewServer(newMux())
	defer srv.Close()

	tests := []struct {
		path     string
		wantCode int
		wantBody string
	}{
		{"/", http.StatusOK, "sample app"},
		{"/healthz", http.StatusOK, "ok"},
		{"/version", http.StatusOK, `"version"`},
		{"/missing", http.StatusNotFound, ""},
	}

	for _, tc := range tests {
		resp, err := http.Get(srv.URL + tc.path)
		if err != nil {
			t.Fatalf("GET %s: %v", tc.path, err)
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode != tc.wantCode {
			t.Errorf("GET %s: status %d, want %d", tc.path, resp.StatusCode, tc.wantCode)
		}
		if !strings.Contains(string(body), tc.wantBody) {
			t.Errorf("GET %s: body %q does not contain %q", tc.path, string(body), tc.wantBody)
		}
	}
}
