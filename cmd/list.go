/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/aluyapeter/williamsgov/models"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `Used to display all your tasks that are still marked as incomplete. Use the --all flag to show both incomplete and completed task`,
	Run: func(cmd *cobra.Command, args []string) {
		taskList, err := models.LoadTasks()
		if err != nil {
			fmt.Printf("Error loading tasks: %v\n", err)
			return
		}

		// if len(taskList.Tasks) == 0 {
		// 	fmt.Println("No task found")
		// 	return
		// }

		// fmt.Println("you have the following tasks:")
		// for i, task := range taskList.Tasks {
		// 	status := " "
		// 	if task.Completed {
		// 		status = "done"
		// 	}
		// 	fmt.Printf("%s %d. %s\n", status, i+1, task.Title)
		// }
		if len(taskList.Tasks) == 0 {
			fmt.Println("No tasks found. Add some tasks to get started using 'williamsgo add \"your task here\"'")
			return
		}

		showAll, _ := cmd.Flags().GetBool("all")

		fmt.Println("Your tasks:")

		displayedCount := 0
		for _, task := range taskList.Tasks{
			if task.Completed && !showAll  {
				continue
			}

			status := "Pending"
			if task.Completed {
				status = "Completed"
			}

			fmt.Printf("\n [ID: %d] %s\n", task.ID, task.Title)
			fmt.Printf("Status: %s\n", status)

			if task.Description != "" {
				fmt.Printf("Description: %s\n", task.Description)
			}

			fmt.Printf("   📅 Created: %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
			
			if task.Completed && task.CompletedAt != nil {
				fmt.Printf("   ✅ Completed: %s\n", task.CompletedAt.Format("2006-01-02 15:04:05"))
			}

			displayedCount++
		}

		fmt.Printf("\n showing %d task(s)", displayedCount)
		if !showAll {
			fmt.Println("(use --all to show completed tasks)")
		} else {
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("all", "a", false, "Show all tasks including completed ones")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
