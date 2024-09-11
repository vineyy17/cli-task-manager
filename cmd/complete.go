package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vineyy17/cli-task-manager/db"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
				continue
			}
			ids = append(ids, id)
		}

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			if task.Done {
				fmt.Printf("Task \"%d\" - \"%s\" is already completed.\n", id, task.Value)
				continue
			}

			err := db.MarkTaskDone(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" - \"%s\" as completed.\n", id, task.Value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
