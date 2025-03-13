package runablescripts

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

//go:embed scripts/tidy.sh
var tidyScript string

func RunTidyScript() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		tempFile, err := os.CreateTemp("", "tidy-*.ps1")
		if err != nil {
			return fmt.Errorf("failed to create temp PowerShell file: %v", err)
		}
		defer os.Remove(tempFile.Name())

		psScript := ConvertBashToPowerShell(tidyScript)

		if _, err := tempFile.WriteString(psScript); err != nil {
			return fmt.Errorf("failed to write PowerShell script: %v", err)
		}

		if err := tempFile.Close(); err != nil {
			return fmt.Errorf("failed to close temp PowerShell file: %v", err)
		}

		cmd = exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", tempFile.Name())
	} else {
		tmpFile, err := os.CreateTemp("", "tidy-*.sh")
		if err != nil {
			return fmt.Errorf("failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(tidyScript); err != nil {
			return fmt.Errorf("failed to write script content: %v", err)
		}

		if err := tmpFile.Close(); err != nil {
			return fmt.Errorf("failed to close temp file: %v", err)
		}

		if err := os.Chmod(tmpFile.Name(), 0755); err != nil {
			return fmt.Errorf("failed to make script executable: %v", err)
		}

		cmd = exec.Command("/bin/bash", tmpFile.Name())
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("script execution error: %v", err)
	}

	return nil
}
