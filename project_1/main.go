package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shivamchhuneja/learning-go/project_1/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {

	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content

}
