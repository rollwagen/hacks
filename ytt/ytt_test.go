package main

import (
	"testing"
)

func TestExtractVideoIDFromUrl(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
		wantErr  bool
	}{
		{
			name:     "Standard YouTube URL",
			url:      "https://www.youtube.com/watch?v=oD-d9B71yLo",
			expected: "oD-d9B71yLo",
			wantErr:  false,
		},
		{
			name:     "Short YouTube URL",
			url:      "https://youtu.be/oD-d9B71yLo",
			expected: "oD-d9B71yLo",
			wantErr:  false,
		},
		{
			name:     "YouTube URL with additional parameters",
			url:      "https://www.youtube.com/watch?v=oD-d9B71yLo&t=10s",
			expected: "oD-d9B71yLo",
			wantErr:  false,
		},
		{
			name:     "YouTube URL without protocol",
			url:      "www.youtube.com/watch?v=oD-d9B71yLo",
			expected: "oD-d9B71yLo",
			wantErr:  false,
		},
		{
			name:     "Invalid YouTube URL",
			url:      "https://www.example.com",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Empty URL",
			url:      "",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractVideoID(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractVideoIDFromUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("extractVideoIDFromUrl() = %v, want %v", got, tt.expected)
			}
		})
	}
}
