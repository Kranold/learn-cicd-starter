package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func createHttpHeader(header, key string) http.Header {
	newHeader := http.Header{}
	newHeader.Set(header, key)
	return newHeader
}

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header     http.Header
		wantOutput string
		wantErr    error
	}{
		"validHeader": {
			header:     createHttpHeader("Authorization", "ApiKey test-key"),
			wantOutput: "test-key",
			wantErr:    nil},
		"invalidHeader": {
			header:     createHttpHeader("Authorization", "noKey lol"),
			wantOutput: "",
			wantErr:    errors.New("malformed authorization header"),
		},
	}

	for name, tc := range tests {
		gotValue, gotError := GetAPIKey(tc.header)
		if !reflect.DeepEqual(gotValue, tc.wantOutput) {
			t.Fatalf("test %s failed: got %v, %v; want %v, %v", name, gotValue, gotError, tc.wantOutput, tc.wantErr)
		}
	}

}
