package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/vineyy17/cli-task-manager/db"
)

var listAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks available")
			return
		}

		// Prepare to format output
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

		if listAll {
			fmt.Fprintln(w, "S/N\tTask\tCreated\tDone")
		} else {
			fmt.Fprintln(w, "S/N\tTask\tCreated")
		}

		// Filter and reorder tasks
		var displayedTasks []db.Task
		for _, task := range tasks {
			if listAll || !task.Done {
				displayedTasks = append(displayedTasks, task)
			}
		}

		// Print tasks with renumbered S/N
		for i, task := range displayedTasks {
			if listAll {
				fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", i+1, task.Value, db.FormatCreatedTime(task.Created), task.Done)
			} else {
				fmt.Fprintf(w, "%d\t%s\t%s\n", i+1, task.Value, db.FormatCreatedTime(task.Created))
			}
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&listAll, "all", "a", false, "Show all tasks, including completed ones")
}
