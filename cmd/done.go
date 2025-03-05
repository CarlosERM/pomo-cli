/*
Copyright © 2025 CARLOS EDUARDO carlosermiranda.dev@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/CarlosERM/pomo-cli/pomo"
	"github.com/spf13/cobra"
)

func doneTask(cmd *cobra.Command, args []string) {
	pomoTasks, err := pomo.ReadTasks(dataFile)

	for _, description := range args {
		for i, task := range pomoTasks {
			if description == task.Description || description == task.Id {
				pomoTasks[i].Done = true
			}
		}
	}

	if err != nil {
		fmt.Println("Failed to read datafile!")
		return
	}

	if err := pomo.SaveTasks(dataFile, pomoTasks); err != nil {
		fmt.Println(err)
	}
}

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done <task_id|task_description>",
	Short: "Marks a task as completed based on ID or description.",
	Long: `The 'done' command allows you to mark a task as completed by specifying either its ID or description.

Examples:
  # Mark a task as completed using an ID
  pomo-cli done 1

  # Mark a task as completed using a description
  pomo-cli done "Pipoca Maluca"`,
	Run: doneTask,
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
