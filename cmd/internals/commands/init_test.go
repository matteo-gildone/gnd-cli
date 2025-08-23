package commands

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
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

func Test_handleInit_InvalidHomeDir(t *testing.T) {
	var stdout, stderr bytes.Buffer

	d := &Dispatcher{
		Stdout:  &stdout,
		Stderr:  &stderr,
		HomeDir: "/invalid/nonexistent/path",
	}

	err := handleInit(d, []string{})
	if err == nil {
		t.Fatal("Expected error using invalid home directory, but got none")
	}

	if !strings.Contains(err.Error(), "failed create directory") {
		t.Errorf("Expected error message about directory creation, got: %v", err)
	}
}
