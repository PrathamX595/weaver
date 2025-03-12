package cmd

import (
	"fmt"
	"strings"
	"time"
	"github.com/PrathamX595/weaver/cmd/flows"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "base command to start weaving your web app",
	Long:  `This is the entry command to creating your web app template`,
	Run: func(cmd *cobra.Command, args []string) {

		frameWork, _ := cmd.Flags().GetString("framework")
		frameWork = strings.ToLower(frameWork)

		name, _ := cmd.Flags().GetString("name")
		name = strings.TrimSpace(name)

		auth, _ := cmd.Flags().GetStringArray("auth")
		var hasAuth bool = checkValidAuth(auth)
        var authMethods []string

		name, frameWork, authMethods = setupProject(name, frameWork, auth, hasAuth)

		isConfirm := flows.SummarizeFlow(name, frameWork, authMethods)

		for !isConfirm {
            name, frameWork, authMethods = setupProject("", "", []string{}, false)
            isConfirm = flows.SummarizeFlow(name, frameWork, authMethods)
        }

		fmt.Println(name, frameWork, authMethods)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String("framework", "", "framework which you want to create")
	createCmd.Flags().String("name", "", "name of your project")
	createCmd.Flags().StringArray("auth", []string{}, "auth if any")
}

func isSelectedFramework(frameWork string) bool{
	if frameWork != "echo" && frameWork != "chi" && frameWork != "fiber" && frameWork != "http" {
		return false
	}
	return true
}

func setupProject( name string, frameWork string, auth []string, hasAuth bool) (string, string, []string) {
    if name == "" {
        name = flows.TextFlow()
        time.Sleep(5 * time.Millisecond)
        confirm := flows.ConformationFlow()
        for !confirm {
            name = flows.TextFlow()
            confirm = flows.ConformationFlow()
        }
    }

    selectedFramework := isSelectedFramework(frameWork)
    for !selectedFramework {
        time.Sleep(5 * time.Millisecond)
        frameWork = flows.ListFlow()
        frameWork = strings.ToLower(frameWork)
        selectedFramework = isSelectedFramework(frameWork)
    }

    var authMethods []string
    if !hasAuth {
        needAuth := flows.NeedAuthFlow()
        if needAuth {
            authMethods = flows.AuthListFlow()
            for i, val := range authMethods {
                authMethods[i] = strings.ToLower(val)
            }
        }
    } else {
        authMethods = auth
    }

    time.Sleep(5 * time.Millisecond)
    
    return name, frameWork, authMethods
}

func checkValidAuth(auth []string) bool {
    for _, val := range auth {
        val = strings.ToLower(val)
        if val != "google" && val != "discord" && val != "github" {
            return false
        }
    }
    return len(auth) > 0
}