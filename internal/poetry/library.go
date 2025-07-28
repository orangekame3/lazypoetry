package poetry

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed poems.json
var poemData []byte

type Poem struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Form    string `json:"form"`
	Content string `json:"content"`
}

var poems []Poem

func init() {
	if err := json.Unmarshal(poemData, &poems); err != nil {
		panic(fmt.Sprintf("failed to load poems: %v", err))
	}
}

func ListPoems() []Poem {
	return poems
}

func GetPoem(id int) (*Poem, error) {
	for _, poem := range poems {
		if poem.ID == id {
			return &poem, nil
		}
	}
	return nil, fmt.Errorf("poem with ID %d not found", id)
}