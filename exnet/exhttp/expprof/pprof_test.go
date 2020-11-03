// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package expprof

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"runtime/pprof"
	"testing"
)

// TestDescriptions checks that the profile names under runtime/pprof package
// have a key in the description map.
func TestDescriptions(t *testing.T) {
	for _, p := range pprof.Profiles() {
		_, ok := profileDescriptions[p.Name()]
		if ok != true {
			t.Errorf("%s does not exist in profileDescriptions map\n", p.Name())
		}
	}
}

func TestServeHTTP(t *testing.T) {
	testCases := []struct {
		path               string
		statusCode         int
		contentType        string
		contentDisposition string
		resp               []byte
	}{
		{"/pprof/<script>scripty<script>", http.StatusNotFound, "text/plain; charset=utf-8", "", []byte("Unknown profile\n")},
		{"/pprof/heap", http.StatusOK, "application/octet-stream", `attachment; filename="heap"`, nil},
		{"/pprof/heap?debug=1", http.StatusOK, "text/plain; charset=utf-8", "", nil},
		{"/pprof/cmdline", http.StatusOK, "text/plain; charset=utf-8", "", nil},
		{"/pprof/profile?seconds=1", http.StatusOK, "application/octet-stream", `attachment; filename="profile"`, nil},
		{"/pprof/symbol", http.StatusOK, "text/plain; charset=utf-8", "", nil},
		{"/pprof/trace", http.StatusOK, "application/octet-stream", `attachment; filename="trace"`, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			req := httptest.NewRequest("GET", "http://example.com"+tc.path, nil)
			w := httptest.NewRecorder()
			ServeHTTP(w, req)

			resp := w.Result()
			if got, want := resp.StatusCode, tc.statusCode; got != want {
				t.Errorf("status code: got %d; want %d", got, want)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("when reading response body, expected non-nil err; got %v", err)
			}
			if got, want := resp.Header.Get("X-Content-Type-Options"), "nosniff"; got != want {
				t.Errorf("X-Content-Type-Options: got %q; want %q", got, want)
			}
			if got, want := resp.Header.Get("Content-Type"), tc.contentType; got != want {
				t.Errorf("Content-Type: got %q; want %q", got, want)
			}
			if got, want := resp.Header.Get("Content-Disposition"), tc.contentDisposition; got != want {
				t.Errorf("Content-Disposition: got %q; want %q", got, want)
			}

			if resp.StatusCode == http.StatusOK {
				return
			}
			if got, want := resp.Header.Get("X-Go-Pprof"), "1"; got != want {
				t.Errorf("X-Go-Pprof: got %q; want %q", got, want)
			}
			if !bytes.Equal(body, tc.resp) {
				t.Errorf("response: got %q; want %q", body, tc.resp)
			}
		})
	}
}
