package auth

import (
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header1 := make(http.Header)
	header2 := make(http.Header)
	header3 := make(http.Header)

	header1.Set("Authorization", "ApiKey 1111111")
	header2.Set("Authorization", "APIKEY 2222222")
	header3.Set("Authorization", "bearer 3333333")

	t.Run("ApiKey", func(t *testing.T) {
		got, err := GetAPIKey(header1)
		if err != nil {
			t.Fatal(err)
		}
		diff := cmp.Diff("1111111", got)
		if diff != "" {
			t.Fatal(diff)
		}
	})

	t.Run("APIKEY", func(t *testing.T) {
		_, err := GetAPIKey(header2)
		if err == nil {
			t.Fatal("Expected an error")
		}
	})

	t.Run("bearer", func(t *testing.T) {
		_, err := GetAPIKey(header3)
		if err == nil {
			t.Fatal("Expected an error")
		}
	})
}
