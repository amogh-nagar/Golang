package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, os.ModeAppend)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Value must not be empty")
	}

	return Note{
		title,
		content,
		time.Now(),
	}, nil
}
