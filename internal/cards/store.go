package cards

import (
	"encoding/json"
	"fmt"
	"os"
)

type Deck struct {
	Name  string `json:"name"`
	Cards []Card `json:"cards"`
}

type Card struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Position int    `json:"position"`
}

type Store struct {
	Path       string `json:"-"`
	Decks      []Deck `json:"decks"`
	ActiveDeck string `json:"active_deck,omitempty"`
}

var s *Store

func (s *Store) DefaultDeck() *Deck {
	if len(s.Decks) == 0 {
		// create default deck if none exist
		s.Decks = append(s.Decks, Deck{Name: "Default"})
	}
	return &s.Decks[0]
}

func (s *Store) CreateDeck(name string) error {
	for _, d := range s.Decks {
		if d.Name == name {
			return fmt.Errorf("deck already exists")
		}
	}

	s.Decks = append(s.Decks, Deck{
		Name:  name,
		Cards: []Card{},
	})

	// If no active deck, set this one
	if s.ActiveDeck == "" {
		s.ActiveDeck = name
	}

	return s.Save()
}

func (s *Store) SelectDeck(name string) error {
	for _, d := range s.Decks {
		if d.Name == name {
			s.ActiveDeck = name
			return s.Save()
		}
	}
	return fmt.Errorf("deck not found")
}

func (s *Store) Active() (*Deck, error) {
	for i := range s.Decks {
		if s.Decks[i].Name == s.ActiveDeck {
			return &s.Decks[i], nil
		}
	}
	return nil, fmt.Errorf("no active deck selected")
}

func newStore(path string) (*Store, error) {
	s = &Store{Path: path}

	if err := s.load(); err != nil {
		return nil, err
	}

	// Re-number cards in all decks
	for di := range s.Decks {
		for i := range s.Decks[di].Cards {
			s.Decks[di].Cards[i].Position = i + 1
		}
	}

	return s, nil
}

func GetStore(path string) (*Store, error) {
	if s == nil {
		return newStore(path)
	}
	return s, nil
}

func (s *Store) DeleteCard(deckName string, position int) error {
	for i := range s.Decks {
		if s.Decks[i].Name == deckName {
			d := &s.Decks[i]
			if position < 1 || position > len(d.Cards) {
				return fmt.Errorf("invalid card position")
			}
			idx := position - 1
			d.Cards = append(d.Cards[:idx], d.Cards[idx+1:]...)
			// renumber
			for i := range d.Cards {
				d.Cards[i].Position = i + 1
			}
			return s.Save()
		}
	}
	return fmt.Errorf("deck not found")
}

func (s *Store) load() error {
	data, err := os.ReadFile(s.Path)
	if err != nil {
		if os.IsNotExist(err) {
			// No file yet → create default deck
			s.Decks = []Deck{{Name: "Default", Cards: []Card{}}}
			s.ActiveDeck = "Default"
			return nil
		}
		return err
	}

	if err := json.Unmarshal(data, s); err != nil {
		return err
	}

	// Migration: if old file had s.Cards field, move it to default deck
	if len(s.Decks) == 0 {
		s.Decks = []Deck{{Name: "Default", Cards: []Card{}}}
	}

	if s.ActiveDeck == "" && len(s.Decks) > 0 {
		s.ActiveDeck = s.Decks[0].Name
	}

	return nil
}

func (s *Store) Save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Path, data, 0644)
}

func (s *Store) Add(card Card) error {
	deck, err := s.Active()
	if err != nil {
		return err
	}

	deck.Cards = append(deck.Cards, card)

	// Re-number positions
	for i := range deck.Cards {
		deck.Cards[i].Position = i + 1
	}

	return s.Save()
}

func (s *Store) Delete(position int) error {
	deck, err := s.Active()
	if err != nil {
		return err
	}

	if position < 1 || position > len(deck.Cards) {
		return fmt.Errorf("invalid card position")
	}

	idx := position - 1
	deck.Cards = append(deck.Cards[:idx], deck.Cards[idx+1:]...)

	for i := range deck.Cards {
		deck.Cards[i].Position = i + 1
	}

	return s.Save()
}

func (s *Store) Update(position int, question, answer string) error {
	deck, err := s.Active()
	if err != nil {
		return err
	}

	if position < 1 || position > len(deck.Cards) {
		return fmt.Errorf("invalid card position")
	}

	idx := position - 1
	deck.Cards[idx].Question = question
	deck.Cards[idx].Answer = answer

	return s.Save()
}
