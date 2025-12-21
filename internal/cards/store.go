package cards

type Card struct {
	Question string
	Answer   string
}

var Cards []Card

func SaveCard(card Card) error {
	Cards = append(Cards, card)
	// optionally write Cards to JSON file
	return nil
}
