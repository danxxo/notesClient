package main

import (
	"bufio"
	"fmt"
	"os"

	client "notesClient/client"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Add Note")
		fmt.Println("2. Get Note")
		fmt.Println("3. Update Note")
		fmt.Println("4. Delete Note")
		fmt.Println("5. Exit")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			client.AddNote(scanner)
		case "2":
			client.GetNote(scanner)
		case "3":
			client.UpdateNotes(scanner)
		case "4":
			client.DeleteNote(scanner)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}

		fmt.Println()
	}
}
