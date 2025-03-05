/*
Copyright © 2025 CARLOS MIRANDA carlosermiranda.dev@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/CarlosERM/pomo-cli/pomo"
	"github.com/spf13/cobra"
)

func removeTask(cmd *cobra.Command, args []string) {
	pomoTasks, err := pomo.ReadTasks(dataFile)

	if err != nil {
		fmt.Println("Failed to read datafile!")
		return
	}

	for _, a := range args {
		for i, task := range pomoTasks {
			if a == task.Description {
				pomoTasks = append(pomoTasks[:i], pomoTasks[i+1:]...)
				break
			} else if a == task.Id {
				pomoTasks = append(pomoTasks[:i], pomoTasks[i+1:]...)
				break
			}
		}
	}

	if err := pomo.SaveTasks(dataFile, pomoTasks); err != nil {
		fmt.Println(err)
	}
}

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove <task_id|task_description> [additional_task_ids_or_descriptions...]",
	Short: "Removes a task or multiple tasks based on ID or description",
	Long: `The 'remove' command allows you to remove one or more tasks by specifying either 
their ID or description. You can remove tasks by their exact name or ID.

Examples:
  # Remove a task by ID
  pomo-cli remove 1

  # Remove tasks by description
  pomo-cli remove Pipoca "Pipoca Maluca"`,
	Run: removeTask,
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
