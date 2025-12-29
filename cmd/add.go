package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Eklund2012/recall/internal/cards"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new flashcard",
	Run:   runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	store, err := cards.GetStore(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Question: ")
	question, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Answer: ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	question = strings.TrimSpace(question)
	answer = strings.TrimSpace(answer)

	if question == "" || answer == "" {
		log.Fatal("Question and answer cannot be empty")
	}

	deck, err := store.Active()
	if err != nil {
		log.Fatal(err)
	}
	deck.Cards = append(deck.Cards, cards.Card{
		Question: question,
		Answer:   answer,
	})
	store.Save()
	fmt.Println("Flashcard added successfully!")
}

func init() {
	rootCmd.AddCommand(addCmd)
}
