package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vineyy17/cli-task-manager/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
