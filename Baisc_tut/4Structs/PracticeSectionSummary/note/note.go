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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v", n.Title, n.Content)
	fmt.Printf("Created At: %s\n", n.CreatedAt.Format(time.RFC1123))
}

//save

func (note *Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	jsonData, err := json.Marshal(note)
	if err != nil {
		fmt.Println("Error marshaling note:", err)
		return err
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

func New(title, content string) (*Note, error) {

	if title == "" || content == "" {
		return &Note{}, errors.New("Invalid Input")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
