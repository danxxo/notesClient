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
	fmt.Println("NEW")

	note := getNoteFromTerminal(scanner)

	jsonRequestBytes, err := json.Marshal(note)
	if err != nil {
		fmt.Println("addNote(): json.Marshal(note)", err)
		return
	}

	url := "http://localhost:8000/add"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("addNote(): http.NewRequest(\"POST\", url, bytes.NewBuffer(jsonRequestBytes))", err)
	}

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("AddNote(): client.Do(req)", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("AddNote(): io.ReadAll(resp.Body)", err)
		return
	}
	responseHandler(body)
}

func GetNote(scanner *bufio.Scanner) {

	fmt.Println("GET")
	fmt.Print("[int] ID: ")

	// Scanner
	scanner.Scan()
	ID := scanner.Text()

	// Marshaling
	jsonRequestBytes, err := json.Marshal(dto.Note{Id: ID})
	if err != nil {
		fmt.Println("getNote(): json.Marshal(dto.Note{Id: ID})", err)
		return
	}

	url := "http://localhost:8000/get"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getNote():  http.NewRequest(\"POST\", url, bytes.NewBuffer(jsonRequestBytes))", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(resp)
		fmt.Println("getNote(): client.Do(req)", err)
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
	fmt.Println("UPDATE")
	fmt.Print("[int] ID: ")

	// Scanner
	scanner.Scan()
	ID := scanner.Text()

	fmt.Println("Fill fields you want to update")
	note := getNoteFromTerminal(scanner)

	note.Id = ID

	jsonRequestBytes, err := json.Marshal(note)
	if err != nil {
		fmt.Println("getNote(): json.Marshal(note)", err)
		return
	}

	url := "http://localhost:8000/update"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getNote(): http.NewRequest(\"POST\", url, bytes.NewBuffer(jsonRequestBytes))", err)
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

	fmt.Println("DELETE")
	fmt.Print("[int] ID: ")

	// Scanner
	scanner.Scan()
	ID := scanner.Text()

	// Marshaling
	jsonRequestBytes, err := json.Marshal(dto.Note{Id: ID})
	if err != nil {
		fmt.Println("getNote(): json.Marshal(dto.Note{Id: ID})", err)
		return
	}

	// POST request
	url := "http://localhost:8000/delete"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getNote(): http.NewRequest(\"POST\", url, bytes.NewBuffer(jsonRequestBytes))", err)
		return
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
