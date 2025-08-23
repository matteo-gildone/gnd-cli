package commands

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func Test_handleInit(t *testing.T) {
	configDir := t.TempDir()
	var stdout, stderr bytes.Buffer

	d := &Dispatcher{
		Stdout:  &stdout,
		Stderr:  &stderr,
		HomeDir: configDir,
	}

	err := handleInit(d, []string{})
	if err != nil {
		t.Fatalf("handleInit failed %v", err)
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
