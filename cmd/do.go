package cmd

import (
	"fmt"
	"okradze/clitaskmanager/db"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task on your list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		ids := make([]int, len(args))

		for i, arg := range args {
			fmt.Println(arg)

			id, err := strconv.Atoi(arg)

			if err != nil {
				fmt.Printf("Failed to parse the id: %s\n", arg)
			}

			ids[i] = id
		}

		db.DeleteTasks(ids)
		fmt.Printf("You have completed the tasks")
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}