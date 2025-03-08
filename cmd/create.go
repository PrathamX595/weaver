package cmd

import (
	"fmt"
	"strings"

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

		fmt.Println(name, frameWork, auth)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String("framework", "", "framework which you want to create")
	createCmd.Flags().String("name", "", "name of your project")
	createCmd.Flags().StringArray("auth", []string{}, "auth if any")
}
