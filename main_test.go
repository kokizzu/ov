package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func Test_initConfig(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		xdg     string
		cfgFile string
		wantErr bool
	}{
		{
			name:    "test-ov.yaml",
			xdg:     "",
			cfgFile: "ov.yaml",
			wantErr: false,
		},
		{
			name:    "test-ov-less.yaml",
			xdg:     "",
			cfgFile: "ov-less.yaml",
			wantErr: false,
		},
		{
			name:    "no-file.yaml",
			xdg:     "",
			cfgFile: "no-file.yaml",
			wantErr: true,
		},
		{
			name:    "not found",
			xdg:     "dummy",
			cfgFile: "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.xdg != "" {
				os.Setenv("XDG_CONFIG_HOME", tt.xdg)
			}
			cfgFile = tt.cfgFile
			// Backup original stderr
			origStderr := os.Stderr

			// Create a buffer to capture stderr output
			r, w, _ := os.Pipe()
			os.Stderr = w

			initConfig()
			w.Close()
			// Restore original stderr
			os.Stderr = origStderr

			// Read captured stderr output
			var buf bytes.Buffer
			if _, err := io.Copy(&buf, r); err != nil {
				t.Fatal(err)
			}
			capturedStderr := buf.String()

			// Now you can assert capturedStderr
			// For example, check if it contains a specific error message
			got := len(capturedStderr) > 0
			if got != tt.wantErr {
				t.Errorf("initConfig() error = %v, wantErr %v", capturedStderr, tt.wantErr)
			}
		})
	}
}
