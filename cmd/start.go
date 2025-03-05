/*
Copyright © 2025 CARLOS MIRANDA carlosermiranda.dev@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/CarlosERM/pomo-cli/pomo"
	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
	"github.com/mbndr/figlet4go"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var taskDescription string

func drawPomo(app *tview.Application, textView *tview.TextView, pomoCountView *tview.TextView, pomodoroQtd int, pomoTask pomo.TaskPomo, isFound bool, idFound int, pomoTasks []pomo.TaskPomo) {
	ticker := time.Tick(1 * time.Millisecond)

	durationPomo := 5 * time.Second

	for i := 0; i < pomodoroQtd; i++ {
		target := time.Now()

		pomoEmoji := strings.Repeat("🍅 ", (i + 1))
		pomoCount := fmt.Sprintf("\n%d/%d  %s", i+1, pomodoroQtd, pomoEmoji)

		app.QueueUpdateDraw(func() {
			pomoCountView.SetText(pomoCount)
		})

		for range ticker {
			timeLeft := durationPomo - time.Since(target)

			if timeLeft < 0 {
				break
			}

			minutes := int(timeLeft.Minutes())
			seconds := int(timeLeft.Seconds()) % 60

			message := fmt.Sprintf("[green]Work: %2dm %2ds", minutes, seconds)

			app.QueueUpdateDraw(func() {
				textView.SetText(message)
			})
		}

		durationRest := 5 * time.Second
		target = time.Now()

		for range ticker {
			timeLeft := durationRest - time.Since(target)

			if timeLeft < 0 {
				break
			}

			minutes := int(timeLeft.Minutes())
			seconds := int(timeLeft.Seconds()) % 60

			message := fmt.Sprintf("[blue]Rest: %2dm %2ds", minutes, seconds)

			app.QueueUpdateDraw(func() {
				textView.SetText(message)
			})
		}

		pomoTask.TimeSpent += int(durationPomo.Seconds())
		pomoTask.Pomodoro += 1
	}

	app.QueueUpdateDraw(func() {
		textView.SetText(`[green]Pomodoro Finished! Closing...`)
	})

	if isFound {
		pomoTasks[idFound] = pomoTask
	} else {
		pomoTasks = append(pomoTasks, pomoTask)
	}

	if err := pomo.SaveTasks(dataFile, pomoTasks); err != nil {
		fmt.Println(err)
	}
	time.Sleep(2 * time.Second)
	app.Stop()
	os.Exit(0)
}

func startPomo(cmd *cobra.Command, args []string) {
	pomoTasks, err := pomo.ReadTasks(dataFile)
	var idFound int
	isFound := false
	id := uuid.New()
	var pomoTask pomo.TaskPomo
	pomodoroQtd := 1
	description := ""
	ascii := figlet4go.NewAsciiRender()

	if err != nil {
		fmt.Println("Failed to read datafile!")
		return
	}

	if len(args) > 0 {
		value, err := strconv.Atoi(args[0])

		if err != nil {
			panic(err)
		}

		pomodoroQtd = value
	}

	if taskDescription != "" {
		description = taskDescription
	} else {
		// if there is no name generate one.
		seed := time.Now().UTC().UnixNano()
		nameGenerator := namegenerator.NewNameGenerator(seed)
		description = nameGenerator.Generate()
	}

	for i, task := range pomoTasks {
		if description == task.Description || description == task.Id {
			isFound = true
			pomoTask = task
			idFound = i
		}
	}

	if !isFound {
		pomoTask = pomo.TaskPomo{Id: id.String(), Description: description, Pomodoro: 0, TimeSpent: 0, Done: false, Priority: priority}
	}

	// DRAW SCREEN
	titleRender, _ := ascii.Render("PomoTask")

	app := tview.NewApplication()

	titleView := tview.NewTextView().
		SetText(titleRender).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetTextColor(tcell.ColorIndianRed)
	taskDescriptionView := tview.NewTextView().
		SetText("\n" + pomoTask.Description).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetTextColor(tcell.ColorIndianRed)
	pomoCountView := tview.NewTextView().
		SetText("").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	textView := tview.NewTextView().
		SetText("").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	go drawPomo(app, textView, pomoCountView, pomodoroQtd, pomoTask, isFound, idFound, pomoTasks)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(titleView, 0, 1, true).
		AddItem(taskDescriptionView, 0, 1, true).
		AddItem(pomoCountView, 0, 1, true).
		AddItem(textView, 0, 5, true)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
	// DRAW SCREEN
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start <number_of_pomodoros>",
	Short: "Starts a Pomodoro session for a specified task or creates a new one",
	Long: `The 'start' command begins a Pomodoro session (default: 25 minutes each) for a specified task. 
You can either define an existing task or create a new one using -d or --description.
If no description is provided, a random task name is generated.

Usage:
  pomo-cli start <number_of_pomodoros> [flags]

Examples:
  # Start a Pomodoro session with 2 Pomodoros
  pomo-cli start 2

  # Start a Pomodoro session for a specific task (or create it if it doesn't exist).
  pomo-cli start 1 -d="Pipoca Maluca"
  # Start a Pomodoro session for a specific task using it ID.
  pomo-cli start 1 -d="1"`,
	Run: startPomo,
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")1

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().StringVarP(&taskDescription, "description", "d", "", "O nome/id da tarefa que era feita com o pomodoro, caso não exista, será criada.")
}
