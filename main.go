// recall CLI
// MIT License © 2025 David Eklund

package main

import (
	"fmt"
	"os"

	"github.com/Eklund2012/recall/cmd"
	"github.com/Eklund2012/recall/internal/client"
)

func main() {
	fmt.Println("Starting Recall Search Test...")

	// Check if the user provided a search term
	if len(os.Args) < 2 {
		fmt.Println("Usage: recall search <query>")
		return
	}

	query := os.Args[1]

	// Call the function from your internal/client package
	client.CallSearch(query)

	cmd.Execute()
}
