package config

import "testing"

func TestInitializeViper(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitializeViper(); (err != nil) != tt.wantErr {
				t.Errorf("InitializeViper() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
