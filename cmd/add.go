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
	store, err := cards.NewStore(dataFile)
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

	err = store.Add(cards.Card{
		Question: question,
		Answer:   answer,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Card added successfully!")
}

func init() {
	rootCmd.AddCommand(addCmd)
}
