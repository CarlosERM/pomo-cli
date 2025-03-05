/*
Copyright © 2025 CARLOS MIRANDA carlosermiranda.dev@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dataFile string
var priority int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pomo-cli",
	Short: "The best pomodoro cli app ever!",
	Long:  "Pomodoro CLI with task management funcionalities.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get home directory: %v", err)
	}

	appDir := filepath.Join(homeDir, ".pomo")
	filePath := filepath.Join(appDir, "pomo-db.json")

	// Ensure the directory exists
	if err := os.MkdirAll(appDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}

	defer file.Close()

	fileInfo, err := os.Stat(filePath)

	if err != nil {
		log.Fatalf("Failed to get info from file: %v", err)
	}

	if fileInfo.Size() == 0 {
		data := "[]"
		if _, err := file.WriteString(data); err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
	}

	viper.AddConfigPath("$HOME/.pomo/")
	viper.SetConfigName("pomo-db")
	viper.SetConfigType("json")
	// viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found.")
		}
	}
}

func init() {
	home, _ := os.UserHomeDir()

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().IntVarP(&priority, "priority", "p", 2, "Priority: 1, 2, 3")
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+"/.pomo/pomo-db.json", "file where all tasks are stored")
	rootCmd.PersistentFlags().BoolVar(&done, "done", false, "List only done tasks.")
}
