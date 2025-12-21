/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Eklund2012/recall/internal/cards"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new flashcard",
	Long:  `Add a new flashcard by providing a question and an answer.`,
	Run:   runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	var question, answer string
	fmt.Print("Enter the question: ")
	fmt.Scanln(&question)
	fmt.Print("Enter the answer: ")
	fmt.Scanln(&answer)

	card := cards.Card{
		Question: question,
		Answer:   answer,
	}

	err := cards.SaveCard(card)
	if err != nil {
		fmt.Println("Error saving card:", err)
		return
	}

	fmt.Println("Card added successfully!")
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
