package cmd

import (
	"fmt"
	"strconv"

	"github.com/aluyapeter/williamsgov/models"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [task_id]",
	Short: "Mark a task as completed",
	Long: `Mark a task as completed by providing its ID number. You can find task ID by running the 'list' command.
	
	Example:
		williamsgov done 3 # Marks task with ID 3 as completed.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Invalid task ID: '%s'. Use the 'list' command to see available task IDs \n", args[0])
			return
		}

		taskList, err := models.LoadTasks()
		if err != nil {
			fmt.Printf("❌ Error loading tasks: %v\n", err)
			return
		}

		err = taskList.CompleteTask(taskID)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		err = models.SaveTasks(taskList)
		if err != nil {
			fmt.Printf("❌ Error ssaving tasks: %v\n", err)
			return
		}

		fmt.Printf("✅ Task %d completed!\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
