package commands

import (
	"os"
	"path/filepath"
	"testing"
)

func TestManager_EnsureConfigDir(t *testing.T) {
	configDir := t.TempDir()

	err, _ := EnsureAppDir(configDir)

	if err != nil {
		t.Fatalf("EnsureConfigDir failed %v", err)
	}

	expectedDirs := []string{
		filepath.Join(configDir, ".gnd"),
		filepath.Join(configDir, ".gnd", "characters"),
	}

	for _, dir := range expectedDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Directory %q was not created", dir)
		}
	}
}
