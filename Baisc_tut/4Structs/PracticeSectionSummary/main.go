package main

import (
	"bufio"
	"fmt"
	"os"
	"practicesection/note"
	"strings"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Println("Note saved successfully!")

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
