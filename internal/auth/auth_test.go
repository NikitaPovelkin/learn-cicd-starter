package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		header     http.Header
		want       string
		wantErr    bool
		errMessage string
	}{
		{
			name:       "No Authorization Header",
			header:     http.Header{},
			want:       "",
			wantErr:    true,
			errMessage: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name: "Malformed Authorization Header",
			header: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			want:       "",
			wantErr:    true,
			errMessage: "malformed authorization header",
		},
		{
			name: "Valid ApiKey Header",
			header: http.Header{
				"Authorization": []string{"ApiKey my-secret-key"},
			},
			want:    "my-secret-key",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.errMessage {
				t.Errorf("GetAPIKey() error message = %v, want %v", err.Error(), tt.errMessage)
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
