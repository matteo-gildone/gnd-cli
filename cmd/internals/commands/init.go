package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	appDir        = ".gnd"
	charactersDir = "characters"
)

func EnsureAppDir(homeDir string) (error, string) {
	configDir := filepath.Join(homeDir, appDir)

	dirs := []string{
		configDir,
		filepath.Join(configDir, charactersDir),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed create directory %s: %w", dir, err), ""
		}
	}

	return nil, configDir
}
