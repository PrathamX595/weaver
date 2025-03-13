package runablescripts

import (
    _ "embed"
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "strings"
)

//go:embed scripts/test.sh
var scriptContent string

func Test(projName string, frameWork string, authMethods []string) error {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        tempFile, err := os.CreateTemp("", "test-*.ps1")
        if err != nil {
            return fmt.Errorf("failed to create temp PowerShell file: %v", err)
        }
        defer os.Remove(tempFile.Name())
        
        psScript := ConvertBashToPowerShell(scriptContent)
        
        if _, err := tempFile.WriteString(psScript); err != nil {
            return fmt.Errorf("failed to write PowerShell script: %v", err)
        }
        
        if err := tempFile.Close(); err != nil {
            return fmt.Errorf("failed to close temp PowerShell file: %v", err)
        }
        
        cmd = exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", tempFile.Name())
    } else {
        // For Linux/Mac, use the original bash script
        tmpFile, err := os.CreateTemp("", "test-*.sh")
        if err != nil {
            return fmt.Errorf("failed to create temp file: %v", err)
        }
        defer os.Remove(tmpFile.Name())
        
        if _, err := tmpFile.WriteString(scriptContent); err != nil {
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
    
    cmd.Env = append(os.Environ(),
        fmt.Sprintf("PROJ_NAME=%s", projName),
        fmt.Sprintf("FRAMEWORK=%s", frameWork),
        fmt.Sprintf("AUTH_METHODS=%s", strings.Join(authMethods, ",")),
    )

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("script execution error: %v", err)
    }
    
    return nil
}

func ConvertBashToPowerShell(bash string) string {
    lines := strings.Split(bash, "\n")
    var psLines []string
    
    // Add PowerShell interpreter line
    psLines = append(psLines, "# Converted from bash script")
    
    // Define PowerShell variables
    psLines = append(psLines, "$PROJ_NAME = $env:PROJ_NAME")
    psLines = append(psLines, "$FRAMEWORK = $env:FRAMEWORK")
    psLines = append(psLines, "$AUTH_METHODS = $env:AUTH_METHODS")
    psLines = append(psLines, "")
    
    inIfBlock := false
    
    for i := 0; i < len(lines); i++ {
        line := lines[i]
        trimmedLine := strings.TrimSpace(line)
        
        // Skip bash shebang and empty lines
        if strings.HasPrefix(trimmedLine, "#!/bin/bash") || trimmedLine == "" {
            continue
        }
        
        // Handle if statements
        if strings.HasPrefix(trimmedLine, "if") && strings.Contains(trimmedLine, "command -v") {
            // Convert bash command check to PowerShell Get-Command
            cmdToCheck := strings.TrimSpace(strings.Split(strings.Split(trimmedLine, "command -v")[1], "&")[0])
            psLines = append(psLines, "if (Get-Command "+cmdToCheck+" -ErrorAction SilentlyContinue) {")
            inIfBlock = true
            continue
        }
        
        // Handle else statements
        if trimmedLine == "else" {
            psLines = append(psLines, "} else {")
            continue
        }
        
        // Handle fi (end of if block)
        if trimmedLine == "fi" {
            psLines = append(psLines, "}")
            inIfBlock = false
            continue
        }
        
        // Convert git init
        if strings.Contains(trimmedLine, "git init") {
            psLines = append(psLines, "    git init")
            continue
        }
        
        // Convert echo statements
        if strings.HasPrefix(trimmedLine, "echo ") {
            content := trimmedLine[5:]
            if strings.HasPrefix(content, "\"") && strings.HasSuffix(content, "\"") {
                if inIfBlock {
                    psLines = append(psLines, "    Write-Host "+content)
                } else {
                    psLines = append(psLines, "Write-Host "+content)
                }
            } else {
                if inIfBlock {
                    psLines = append(psLines, "    Write-Host \""+content+"\"")
                } else {
                    psLines = append(psLines, "Write-Host \""+content+"\"")
                }
            }
            continue
        }
        
        // Convert variable assignment with command substitution
        if strings.Contains(trimmedLine, "SAFE_PROJ_NAME=$(echo") {
            psLines = append(psLines, "$SAFE_PROJ_NAME = ($PROJ_NAME -replace ' ', '-' -replace '[^a-zA-Z0-9\\-]', '').ToLower()")
            continue
        }
        
        // Convert go mod init command (with example.com prefix to avoid module issues)
        if strings.Contains(trimmedLine, "go mod init") {
            if inIfBlock {
                psLines = append(psLines, "    go mod init $PROJ_NAME")
            } else {
                psLines = append(psLines, "go mod init $PROJ_NAME")
            }
            continue
        }
        
        // Handle other environment variable references
        if strings.Contains(line, "$") {
            if !strings.Contains(line, "$env:") {
                for _, envVar := range []string{"PROJ_NAME", "FRAMEWORK", "AUTH_METHODS", "SAFE_PROJ_NAME"} {
                    line = strings.ReplaceAll(line, "$"+envVar, "$"+envVar)
                }
            }
        }
        
        // Add proper indentation if inside an if block
        if inIfBlock && trimmedLine != "" {
            psLines = append(psLines, "    "+line)
        } else {
            // Add the line (possibly modified)
            psLines = append(psLines, line)
        }
    }
    
    return strings.Join(psLines, "\n")
}