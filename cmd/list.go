package cmd

import (
	"fmt"
	"okradze/clitaskmanager/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := db.GetTasks()

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Why not take a vacation? :)")
		}

		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}