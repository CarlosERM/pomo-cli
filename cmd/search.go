/*
Copyright © 2025 CARLOS MIRANDA carlosermiranda.dev@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/CarlosERM/pomo-cli/pomo"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pomoTasks, err := pomo.ReadTasks(dataFile)
		var searchedItems []pomo.TaskPomo

		if err != nil {
			fmt.Println("Failed to read datafile!")
			return
		}

		for _, a := range args {
			for i, task := range pomoTasks {
				if strings.Contains(task.Description, a) {
					searchedItems = append(searchedItems, pomoTasks[i])
				}
			}
		}

		w := tabwriter.NewWriter(os.Stdout, 4, 1, 5, ' ', tabwriter.StripEscape)
		fmt.Fprintf(w, "ID\tDescription\tPriority\tPomodoro\tTimeSpent\tDone\n")

		for _, task := range searchedItems {
			doneMessage := ""

			if task.Done {
				doneMessage = "X"
			}

			if task.Done == done {
				fmt.Fprintf(w, "%s\t%s\t%d\t%d\t%d\t%s\n", task.Id, task.Description, task.Priority, task.Pomodoro, task.TimeSpent, doneMessage)
			}
		}
		w.Flush()

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
