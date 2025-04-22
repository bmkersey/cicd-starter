package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name       string
		headers    http.Header
		wantAPIKey string
		wantErr    bool
	}

	tests := []test{
		{
			name:       "valid api key",
			headers:    http.Header{"Authorization": []string{"ApiKey test-api-key"}},
			wantAPIKey: "test-api-key",
			wantErr:    true,
		},
		{
			name:       "missing authorization header",
			headers:    http.Header{},
			wantAPIKey: "",
			wantErr:    true,
		},
		{
			name:       "malformed authorization header - wrong prefix",
			headers:    http.Header{"Authorization": []string{"Bearer test-api-key"}},
			wantAPIKey: "",
			wantErr:    true,
		},
		{
			name:       "malformed authorization header - no space",
			headers:    http.Header{"Authorization": []string{"ApiKeytest-api-key"}},
			wantAPIKey: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request with headers
			req := &http.Request{
				Header: tt.headers,
			}

			// Call GetAPIKey
			gotAPIKey, err := GetAPIKey(req.Header)

			// Check results
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAPIKey != tt.wantAPIKey {
				t.Errorf("GetAPIKey() = %v, want %v", gotAPIKey, tt.wantAPIKey)
			}
		})
	}
}
