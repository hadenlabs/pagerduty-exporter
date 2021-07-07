package config

import "testing"

func TestReadConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadConfig(); (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
