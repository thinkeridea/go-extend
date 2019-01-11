// Copyright (C) 2018  Qi Yin <qiyin@thinkeridea.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package expprof

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"runtime/pprof"
	"sync/atomic"
	"testing"
	"time"
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
		{"/debug/pprof/<script>scripty<script>", http.StatusNotFound, "text/plain; charset=utf-8", "", []byte("Unknown profile\n")},
		{"/debug/pprof/heap", http.StatusOK, "application/octet-stream", `attachment; filename="heap"`, nil},
		{"/debug/pprof/heap?debug=1", http.StatusOK, "text/plain; charset=utf-8", "", nil},
		{"/debug/pprof/cmdline", http.StatusOK, "text/plain; charset=utf-8", "", nil},
		{"/debug/pprof/profile?seconds=1", http.StatusOK, "application/octet-stream", `attachment; filename="profile"`, nil},
		{"/debug/pprof/symbol", http.StatusOK, "text/plain; charset=utf-8", "", nil},
		{"/debug/pprof/trace", http.StatusOK, "application/octet-stream", `attachment; filename="trace"`, nil},
	}

	atomic.StoreInt32(&apiState, int32(OpenApi))
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

func TestServeHTTPApiClose(t *testing.T) {
	testCases := []struct {
		path       string
		statusCode int
	}{
		{"/debug/pprof/<script>scripty<script>", http.StatusNotFound},
		{"/debug/pprof/heap", http.StatusNotFound},
		{"/debug/pprof/heap?debug=1", http.StatusNotFound},
		{"/debug/pprof/cmdline", http.StatusNotFound},
		{"/debug/pprof/profile?seconds=1", http.StatusNotFound},
		{"/debug/pprof/symbol", http.StatusNotFound},
		{"/debug/pprof/trace", http.StatusNotFound},
	}

	atomic.StoreInt32(&apiState, int32(CloseApi))
	for _, tc := range testCases {
		req := httptest.NewRequest("GET", "http://example.com"+tc.path, nil)
		w := httptest.NewRecorder()
		ServeHTTP(w, req)

		resp := w.Result()
		if got, want := resp.StatusCode, tc.statusCode; got != want {
			t.Errorf("status code: got %d; want %d", got, want)
		}
	}
}

func TestSetApiState(t *testing.T) {
	OpenTime = 100 * time.Millisecond
	atomic.StoreInt32(&apiState, int32(CloseApi))

	SetApiState(OpenApi)
	if atomic.LoadInt32(&apiState) != int32(OpenApi) {
		t.Errorf("apiState != %d", OpenApi)
	}

	time.Sleep(1 * time.Second)
	if atomic.LoadInt32(&apiState) != int32(CloseApi) {
		t.Errorf("apiState != %d", CloseApi)
	}

	SetApiState(OpenApi)
	SetApiState(OpenApi)
	SetApiState(CloseApi)
	if atomic.LoadInt32(&apiState) != int32(CloseApi) {
		t.Errorf("apiState != %d", CloseApi)
	}
}
