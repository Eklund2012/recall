package cards

import (
	"encoding/json"
	"fmt"
	"os"
)

type Card struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Position int    `json:"position"`
}

type Store struct {
	Path  string
	Cards []Card
}

var s *Store

func newStore(path string) (*Store, error) {
	s = &Store{Path: path}

	if err := s.load(); err != nil {
		return nil, err
	}
	for i := range s.Cards {
		s.Cards[i].Position = i + 1
	}
	return s, nil
}

func GetStore(path string) (*Store, error) {
	if s == nil {
		return newStore(path)
	}
	return s, nil
}

func (s *Store) load() error {
	data, err := os.ReadFile(s.Path)
	if err != nil {
		if os.IsNotExist(err) {
			s.Cards = []Card{}
			return nil
		}
		return err
	}

	return json.Unmarshal(data, &s.Cards)
}

func (s *Store) Save() error {
	data, err := json.MarshalIndent(s.Cards, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Path, data, 0644)
}

func (s *Store) Add(card Card) error {
	s.Cards = append(s.Cards, card)
	return s.Save()
}

func (s *Store) Delete(position int) error {
	if position < 1 || position > len(s.Cards) {
		return fmt.Errorf("invalid card position")
	}

	// Convert position (1-based) to index (0-based)
	idx := position - 1

	s.Cards = append(s.Cards[:idx], s.Cards[idx+1:]...)

	// Re-number positions
	for i := range s.Cards {
		s.Cards[i].Position = i + 1
	}

	return s.Save()
}
