package cmd

import (
	"fmt"
	"strconv"

	"github.com/aluyapeter/williamsgov/models"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task_id]",
	Short: "Delete a task by ID",
	Long: `Delete a task using its ID number. You can find task ID by running the 'list' command
	
	Example: 
		williamsgov delete 4 # Marks task with ID 4 as deleted`,
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

		err = taskList.DeleteTask(taskID)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		err = models.SaveTasks(taskList)
		if err != nil {
			fmt.Printf("❌ Error saving tasks: %v\n", err)
			return
		}

		fmt.Printf("✅ Task %d deleted! \n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
