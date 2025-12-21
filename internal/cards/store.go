package cards

import (
	"encoding/json"
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

func NewStore(path string) (*Store, error) {
	s := &Store{Path: path}

	if err := s.load(); err != nil {
		return nil, err
	}
	for i := range s.Cards {
		s.Cards[i].Position = i + 1
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
