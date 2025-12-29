package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "recall",
	Short: "recall is a CLI study helper",
	Long:  "recall is a simple command-line tool to help you study using flashcards.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.my_project.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	/*
		home, err := homedir.Dir()
		if err != nil {
			log.Println("Unable to detect home directory. Please set data file using --datafile.")
		}
	*/
	//rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".tridos.json", "data file to store todos")

	// For simplicity, use local directory for now
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", ".tridos.json", "data file to store todos")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
