/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Eklund2012/recall/internal/cards"
	"github.com/spf13/cobra"
)

// listDecksCmd represents the listDecks command
var listDecksCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all decks",
	Run: func(cmd *cobra.Command, args []string) {
		store, err := cards.GetStore(dataFile)
		if err != nil {
			log.Fatal(err)
		}

		if len(store.Decks) == 0 {
			fmt.Println("No decks found.")
			return
		}
		fmt.Println("Decks:")

		for _, d := range store.Decks {
			active := ""
			if d.Name == store.ActiveDeck {
				active = "(active)"
			}
			fmt.Printf("- %s %s\n", d.Name, active)
		}
	},
}

func init() {
	deckCmd.AddCommand(listDecksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listDecksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listDecksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
