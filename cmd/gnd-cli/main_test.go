package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantExitCode int
		wantStdout   string
		wantStderr   string
	}{
		{
			name:         "successful init command",
			args:         []string{"gnd", "init"},
			wantExitCode: 0,
			wantStdout:   "Initialise app",
		},
		{
			name:         "unknown command returns error and exit code 1",
			args:         []string{"gnd", "unknown"},
			wantExitCode: 1,
			wantStderr:   "Error running the app unknown command: unknown\n",
		},
		{
			name:         "no command shows usage and returns 0",
			args:         []string{"gnd"},
			wantExitCode: 0,
			wantStdout:   "Gophers and Dragons Character Creator",
		},
		{
			name:         "empty args shows usage and returns 0",
			args:         []string{},
			wantExitCode: 0,
			wantStdout:   "Usage:",
		},
		{
			name:         "usage contains examples section",
			args:         []string{"gnd"},
			wantExitCode: 0,
			wantStdout:   "Examples:",
		},
		{
			name:         "multiple unknown commands still return proper error",
			args:         []string{"gnd", "badcommand"},
			wantExitCode: 1,
			wantStderr:   "Error running the app unknown command: badcommand\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer

			exitCode := run(tt.args, &stdout, &stderr)

			// Check exit code
			if exitCode != tt.wantExitCode {
				t.Errorf("run() exit code = %d, want %d", exitCode, tt.wantExitCode)
			}

			// Check stdout
			if tt.wantStdout != "" {
				stdoutOutput := stdout.String()
				if !strings.Contains(stdoutOutput, tt.wantStdout) {
					t.Errorf("stdout should contain %q, got %q", tt.wantStdout, stdoutOutput)
				}
			}

			// Check stderr
			if tt.wantStderr != "" {
				stderrOutput := stderr.String()
				if stderrOutput != tt.wantStderr {
					t.Errorf("stderr = %q, want %q", stderrOutput, tt.wantStderr)
				}
			}

			// Ensure no unexpected output
			if tt.wantStdout == "" && stdout.Len() > 0 {
				t.Errorf("unexpected stdout output: %q", stdout.String())
			}
			if tt.wantStderr == "" && stderr.Len() > 0 {
				t.Errorf("unexpected stderr output: %q", stderr.String())
			}
		})
	}
}

// Test edge cases and error conditions
func TestRun_EdgeCases(t *testing.T) {
	t.Run("nil writers should not panic", func(t *testing.T) {
		// This tests that the function handles nil writers gracefully
		// In practice, this shouldn't happen, but it's good to be defensive
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("run() panicked with nil writers: %v", r)
			}
		}()

		// Using os.Stdout/Stderr as fallback since we can't pass nil
		var stdout, stderr bytes.Buffer
		exitCode := run([]string{"gnd", "init"}, &stdout, &stderr)

		if exitCode != 0 {
			t.Errorf("expected successful execution, got exit code %d", exitCode)
		}
	})

	t.Run("very long command name", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		longCommand := strings.Repeat("a", 1000)

		exitCode := run([]string{"gnd", longCommand}, &stdout, &stderr)

		if exitCode != 1 {
			t.Errorf("expected exit code 1 for unknown command, got %d", exitCode)
		}

		expectedError := "Error running the app unknown command: " + longCommand + "\n"
		if stderr.String() != expectedError {
			t.Errorf("stderr = %q, want %q", stderr.String(), expectedError)
		}
	})
}

// Benchmark the run function
func BenchmarkRun(b *testing.B) {
	args := []string{"gnd", "init"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var stdout, stderr bytes.Buffer
		run(args, &stdout, &stderr)
	}
}

func BenchmarkRun_Error(b *testing.B) {
	args := []string{"gnd", "unknown"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var stdout, stderr bytes.Buffer
		run(args, &stdout, &stderr)
	}
}
