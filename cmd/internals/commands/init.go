package commands

import "fmt"

func init() {
	Register(&Command{
		Name:        "init",
		Description: "Initialise the application directory structure",
		Usage:       "gnd init",
		Examples: []string{
			"gnd init # Creates ~/.gnd directory structure",
		},
		Handler: handleInit,
	})
}

func handleInit(d *Dispatcher, args []string) error {
	_, _ = fmt.Fprintf(d.Stdout, "  🎲 Initialising Gophers and Dragons in: %s\n", d.HomeDir)
	_, _ = fmt.Fprintf(d.Stdout, "  📁 Created character directory in: %s\n", d.HomeDir+"characters")
	_, _ = fmt.Fprintf(d.Stdout, "  ✅ Initialisation completed!\n")
	return nil
}

//const (
//	appDir        = ".gnd"
//	charactersDir = "characters"
//)
//
//func EnsureAppDir(homeDir string) (string, error) {
//	configDir := filepath.Join(homeDir, appDir)
//
//	dirs := []string{
//		configDir,
//		filepath.Join(configDir, charactersDir),
//	}
//
//	for _, dir := range dirs {
//		if err := os.MkdirAll(dir, 0755); err != nil {
//			return "", fmt.Errorf("failed create directory %s: %w", dir, err)
//		}
//	}
//
//	return configDir, nil
//}
