package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		inputHeaders http.Header
		wantString   string
		wantErr      bool
	}{
		"Valid API Key": {
			inputHeaders: http.Header{
				"Authorization": []string{"ApiKey valid_api_key"},
			},
			wantString: "valid_api_key",
			wantErr:    false,
		},
		"Missing Authorization header": {
			inputHeaders: http.Header{},
			wantString:   "",
			wantErr:      true,
		},
		"Malformed Authorization header": {
			inputHeaders: http.Header{
				"Authorization": []string{"InvalidApiKey api_key"},
			},
			wantString: "",
			wantErr:    true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gotApiKey, err := GetAPIKey(tt.inputHeaders)
			if err != nil && !tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotApiKey != tt.wantString {
				t.Errorf("GetAPIKey() = %v, want %v", gotApiKey, tt.wantString)
			}
		})
	}
}
