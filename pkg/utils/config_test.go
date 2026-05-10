package utils

import (
	"os"
	"path/filepath"
	"testing"

	"slices"
)

func TestConfig(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	// Test GetConfigPath
	path, err := GetConfigPath()
	if err != nil {
		t.Fatalf("GetConfigPath failed: %v", err)
	}
	expectedPath := filepath.Join(tmpDir, ".tmz.yaml")
	if path != expectedPath {
		t.Errorf("expected path %s, got %s", expectedPath, path)
	}

	// Test LoadConfig (missing file)
	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}
	if len(config.Timezones) != 0 {
		t.Errorf("expected empty timezones, got %v", config.Timezones)
	}

	// Test SaveConfig
	config.Timezones = []string{"Asia/Kolkata", "America/New_York"}
	err = SaveConfig(config)
	if err != nil {
		t.Fatalf("SaveConfig failed: %v", err)
	}

	// Test LoadConfig (existing file)
	config2, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}
	if !slices.Equal(config.Timezones, config2.Timezones) {
		t.Errorf("expected %v, got %v", config.Timezones, config2.Timezones)
	}

	// Test manual edit and reload
	err = os.WriteFile(path, []byte("timezones:\n  - Europe/London\n"), 0644)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
	config3, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}
	if !slices.Equal(config3.Timezones, []string{"Europe/London"}) {
		t.Errorf("expected [Europe/London], got %v", config3.Timezones)
	}
}
