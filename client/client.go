package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	dto "notesClient/models/dto"
)

func AddNote(scanner *bufio.Scanner) {
	fmt.Println("Lets create a new Note!")

	note := getNoteFromTerminal(scanner)

	jsonRequestBytes, err := json.Marshal(note)
	if err != nil {
		fmt.Println("addNote()", err)
	}

	url := "http://localhost:8000/add"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("addNote()", err)
	}

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("NoteAdd()", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("addNote()", err)
		return
	}
	responseHandler(body)
}

func GetNote(scanner *bufio.Scanner) {
	note := getNoteFromTerminal(scanner)

	jsonRequestBytes, err := json.Marshal(note)
	if err != nil {
		fmt.Println("getNote()", err)
		return
	}

	url := "http://localhost:8000/get"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getNote()", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("getNote()", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseHandler(body)
}

func UpdateNotes(scanner *bufio.Scanner) {
	fmt.Println("Let s Update Note. Dont forget fill Phone field")
	note := getNoteFromTerminal(scanner)

	jsonRequestBytes, err := json.Marshal(note)
	if err != nil {
		fmt.Println("getNote()", err)
		return
	}

	url := "http://localhost:8000/update"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getNote()", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseHandler(body)
}

func DeleteNote(scanner *bufio.Scanner) {

	fmt.Println("To delete Note Enter the ID")
	fmt.Print("Enter ID: ")

	// Scanner
	scanner.Scan()
	ID := scanner.Text()

	// Marshaling
	jsonRequestBytes, err := json.Marshal(dto.Note{Id: ID})
	if err != nil {
		fmt.Println("getNote()", err)
	}

	// POST request
	url := "http://localhost:8000/delete"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getNote()", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Response handling
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseHandler(body)
}
