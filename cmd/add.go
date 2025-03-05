/*
Copyright © 2025 CARLOS MIRANDA carlosermiranda.dev@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/CarlosERM/pomo-cli/pomo"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func addTask(cmd *cobra.Command, args []string) {
	pomoTasks, err := pomo.ReadTasks(dataFile)
	id := uuid.New()

	if err != nil {
		fmt.Println("Failed to read datafile!")
		return
	}

	for _, a := range args {
		task := pomo.TaskPomo{Id: id.String(), Description: a, Done: false, Priority: priority}
		pomoTasks = append(pomoTasks, task)
	}

	if err := pomo.SaveTasks(dataFile, pomoTasks); err != nil {
		fmt.Println(err)
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <task_name> [additional_task_names...]",
	Short: "Adds a new task or multiple ones",
	Long: `The 'add' command allows you to add one or more tasks.
	
Examples:
  # Add a single task
  pomo-cli add Pipoca

  # Add multiple tasks
  pomo-cli add Pipoca Maluca

  # Add a task with multiple words
  pomo-cli add "Pipoca Maluca"

If a task contains spaces, wrap it in double quotes.`,
	Run: addTask,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
