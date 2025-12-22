package cmd

import (
	"log"

	"github.com/Eklund2012/recall/internal/cards"
	tuilist "github.com/Eklund2012/recall/internal/tui/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all flashcards",
	Run:   runList,
}

func runList(cmd *cobra.Command, args []string) {
	store, err := cards.GetStore(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(tuilist.NewModel(store))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
