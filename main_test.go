package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestUploadFile(t *testing.T) {
	t.Run("returns message", func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/upload", nil)
			response := httptest.NewRecorder()

			uploadFile(response, request)

			got := response.Body.String()
			want := "Uploading File"

			if got != want {
					t.Errorf("got %q, want %q", got, want)
			}
	})
}