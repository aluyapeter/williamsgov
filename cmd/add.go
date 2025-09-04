package cmd

import (
	"fmt"
	"strings"

	"github.com/aluyapeter/williamsgov/models"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task title]",
	Short: "Add a new task",
	Long: `Add a new task to your task list. Also, you can optionally add a description of your task using the --description flag
	
	Examples:
	williamsgov add "Buy beans"
	williamsgov add "work on project" --description "remember to complete before deadline"
	williamsgov add "call colleague" -d "weekly meeting"`,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		description, _ := cmd.Flags().GetString("description")

		taskList, err := models.LoadTasks()
		if err != nil {
			fmt.Printf("❌ Error loading tasks: %v\n", err)
			return
		}

		taskList.AddTask(title, description)

		err = models.SaveTasks(taskList)
		if err != nil {
			fmt.Printf("❌ Error savig tasks: %v\n", err)
		}

		fmt.Printf("✅ Task added successfully: '%s'\n", title)
		if description != "" {
			fmt.Printf("Description: %s\n", description)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("description", "d", "", "optional description")
}
