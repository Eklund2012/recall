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

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select an active deck",
	Long:  `Select a deck to make it the active deck for studying.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load the store
		store, err := cards.GetStore(dataFile)
		if err != nil {
			log.Fatal(err)
		}

		// Select the deck
		if err := store.SelectDeck(deckName); err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Printf("Deck '%s' selected!\n", deckName)
	},
}

func init() {
	deckCmd.AddCommand(selectCmd)

	// Flag to set deck name
	selectCmd.Flags().StringVarP(&deckName, "name", "n", "", "Name of the deck to select")
	selectCmd.MarkFlagRequired("name")
}
