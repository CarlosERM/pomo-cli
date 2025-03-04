/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"text/tabwriter"

	"github.com/CarlosERM/pomo-cli/pomo"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var done bool
var all bool

func listAllTasks(pomoTasks []pomo.TaskPomo) {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true)

	colNames := reflect.TypeOf(pomo.TaskPomo{})

	cols := reflect.TypeOf(pomo.TaskPomo{}).NumField()
	rows := len(pomoTasks)

	for i := 0; i < colNames.NumField(); i++ {
		color := tcell.ColorIndianRed
		field := colNames.Field(i)
		table.SetCell(0, i, tview.NewTableCell(field.Name).SetTextColor(color))
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			rowPomo := reflect.ValueOf(pomoTasks[r])

			value := rowPomo.Field(c).Interface()

			strValue := fmt.Sprintf("%v", value)

			color := tcell.ColorWhite

			table.SetCell(r+1, c,
				tview.NewTableCell(strValue).
					SetTextColor(color).
					SetAlign(tview.AlignCenter))

		}
	}

	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}

func listTasks(cmd *cobra.Command, args []string) {
	pomoTasks, err := pomo.ReadTasks(dataFile)

	if err != nil {
		fmt.Println("Failed to read datafile!")
		return
	}

	sort.Slice(pomoTasks, func(i, j int) bool {
		return pomoTasks[i].Priority > pomoTasks[j].Priority
	})

	w := tabwriter.NewWriter(os.Stdout, 4, 1, 5, ' ', tabwriter.StripEscape)

	fmt.Fprintf(w, "ID\tDescription\tPriority\tPomodoro\tTimeSpent\tDone\n")

	if all {
		listAllTasks(pomoTasks)
	}

	for _, task := range pomoTasks {
		doneMessage := ""

		if task.Done {
			doneMessage = "X"
		}

		if task.Done == done {
			fmt.Fprintf(w, "%s\t%s\t%d\t%d\t%d\t%s\n", task.Id, task.Description, task.Priority, task.Pomodoro, task.TimeSpent, doneMessage)
		}
	}

	w.Flush()
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: listTasks,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listCmd.Flags().BoolVarP(&done, "done", "d", false, "List only done tasks.")
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "List all tasks.")
}
