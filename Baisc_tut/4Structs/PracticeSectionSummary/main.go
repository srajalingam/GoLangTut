package main

import (
	todo "Interface"
	"bufio"
	"fmt"
	"os"
	"practicesection/note"
	"strings"
)

type saver interface {
	Save() error
}

func main() {

	title, content := getNoteData()

	todoText := getTodoData()

	todo, err := todo.New(todoText)

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	todo.Display()
	err = saveData(todo)
	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {

		return
	}

}

func getTodoData() string {
	content := getUserInput("Enter the content of the todo: ")
	return content
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data:", err)
		return err
	}
	fmt.Println("Data saved successfully!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the title of the book: ")

	content := getUserInput("Enter the content of the book: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	var value string
	// fmt.Scanln(&value) // use buffered reader for more complex input

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r") // for Windows

	return value
}
