package cmd

import (
	"fmt"
	"log"

	"github.com/Eklund2012/recall/internal/cards"
	"github.com/spf13/cobra"
)

var deckName string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new deck",
	Long:  `Create a new deck to organize your flashcards.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load the store
		store, err := cards.GetStore(dataFile)
		if err != nil {
			log.Fatal(err)
		}

		// Create the deck
		if err := store.CreateDeck(deckName); err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Printf("Deck '%s' created!\n", deckName)
	},
}

func init() {
	deckCmd.AddCommand(createCmd)

	// Flag to set deck name
	createCmd.Flags().StringVarP(&deckName, "name", "n", "", "Name of the new deck")
	createCmd.MarkFlagRequired("name")
}
