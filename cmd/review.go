package cmd

import (
	"log"

	"github.com/Eklund2012/recall/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// reviewCmd represents the review command
var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Review your flashcards",
	Long:  `Review your flashcards by going through the stored cards one by one.`,
	Run:   runReview,
}

func runReview(cmd *cobra.Command, args []string) {
	p := tea.NewProgram(tui.NewModel())
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(reviewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reviewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reviewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
