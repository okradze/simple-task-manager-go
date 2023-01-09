package cmd

import (
	"fmt"
	"okradze/clitaskmanager/db"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		db.AddTask(task)
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}