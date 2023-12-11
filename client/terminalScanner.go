package client

import (
	"bufio"
	"fmt"

	dto "notesClient/models/dto"
)

func getNoteFromTerminal(scanner *bufio.Scanner) dto.Note {

	fmt.Print("Enter Name: ")
	scanner.Scan()
	Name := scanner.Text()

	fmt.Print("Enter LastName: ")
	scanner.Scan()
	LastName := scanner.Text()

	fmt.Print("Enter Note: ")
	scanner.Scan()
	Note := scanner.Text()

	return dto.Note{AuthorFirstName: Name, AuthorLastName: LastName, Note: Note}
}
