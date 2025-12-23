package cmd

import (
	"log"

	"github.com/Eklund2012/recall/internal/cards"
	review "github.com/Eklund2012/recall/internal/tui/review"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// reviewCmd represents the review command
var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Review your flashcards",
	Long:  `Review your flashcards by going through the stored cards one by one`,
	Run:   runReview,
}

func runReview(cmd *cobra.Command, args []string) {
	store, err := cards.GetStore(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(review.NewModel(store))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(reviewCmd)
}
