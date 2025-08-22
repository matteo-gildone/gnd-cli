package commands

//
//import (
//	"bytes"
//	"strings"
//	"testing"
//)
//
//func TestNew(t *testing.T) {
//	var stdout, stderr bytes.Buffer
//
//	dispatcher, err := New(&stdout, &stderr)
//
//	if err != nil {
//		t.Fatalf("expected no error, got %v", err)
//	}
//
//	if dispatcher == nil {
//		t.Fatal("New() returned nil")
//	}
//	if dispatcher.Stdout == nil {
//		t.Error("Stdout should not be nil")
//	}
//	if dispatcher.Stderr == nil {
//		t.Error("Stderr should not be nil")
//	}
//
//	if dispatcher.HomeDir == "" {
//		t.Error("app dir should be set")
//	}
//}
//
//func TestDispatcher_Dispatch(t *testing.T) {
//	tests := []struct {
//		name       string
//		args       []string
//		wantErr    bool
//		wantErrMsg string
//		wantOutput string
//	}{
//		{
//			name:       "init command success",
//			args:       []string{"gnd", "init"},
//			wantErr:    false,
//			wantOutput: "Initialise app",
//		},
//		{
//			name:       "help command success",
//			args:       []string{"gnd", "help"},
//			wantErr:    false,
//			wantOutput: "Usage:",
//		},
//		{
//			name:       "unknown command",
//			args:       []string{"gnd", "unknown"},
//			wantErr:    true,
//			wantErrMsg: "unknown command: unknown",
//		},
//		{
//			name:       "no arguments shows usage",
//			args:       []string{"gnd"},
//			wantErr:    false,
//			wantOutput: "Gophers and Dragons Character Creator",
//		},
//		{
//			name:       "empty args shows usage",
//			args:       []string{},
//			wantErr:    false,
//			wantOutput: "Usage:",
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			var stdout, stderr bytes.Buffer
//
//			dispatcher, _ := New(&stdout, &stderr)
//
//			err := dispatcher.Dispatch(tt.args)
//
//			// Check error expectations
//			if tt.wantErr {
//				if err == nil {
//					t.Fatal("expected error but got none")
//				}
//				if tt.wantErrMsg != "" && !strings.Contains(err.Error(), tt.wantErrMsg) {
//					t.Errorf("error = %q, want to contain %q", err.Error(), tt.wantErrMsg)
//				}
//			} else {
//				if err != nil {
//					t.Fatalf("unexpected error: %v", err)
//				}
//			}
//
//			// Check output expectations
//			if tt.wantOutput != "" {
//				output := stdout.String()
//				if !strings.Contains(output, tt.wantOutput) {
//					t.Errorf("output = %q, want to contain %q", output, tt.wantOutput)
//				}
//			}
//		})
//	}
//}
//
//func TestDispatcher_printUsage(t *testing.T) {
//	var stdout, stderr bytes.Buffer
//
//	dispatcher, _ := New(&stdout, &stderr)
//
//	err := dispatcher.printUsage()
//
//	if err != nil {
//		t.Fatalf("printUsage() returned unexpected error: %v", err)
//	}
//
//	output := stdout.String()
//	requiredStrings := []string{
//		"Gophers and Dragons Character Creator",
//		"Usage:",
//		"gnd init",
//		"Examples:",
//	}
//
//	for _, required := range requiredStrings {
//		if !strings.Contains(output, required) {
//			t.Errorf("usage output missing %q, got: %q", required, output)
//		}
//	}
//}
