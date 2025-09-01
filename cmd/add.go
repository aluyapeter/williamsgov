/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
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
	williamsgo add "Buy beans"
	williams go add "work on project" --description "remember to complete before deadline"
	williamsgo add "call colleague" -d "weekly meeting"`,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		description, _ := cmd.Flags().GetString("description")

		taskList, err := models.LoadTasks()
		if err != nil {
			fmt.Printf("Error loading tasks: %v\n", err)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
