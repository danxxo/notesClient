package client

import (
	"encoding/json"
	"fmt"

	dto "notesClient/models/dto"
)

func responseHandler(body []byte) {

	var response dto.Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.Result == "error" {
		fmt.Print("\nserver returned error: \x1b[31m", response.Error, "\x1b[0m\n")
		return
	}
	if response.Result == "OK" {
		fmt.Print("\n\x1b[32m# Operation Succeded!\x1b[0m\n")
	}

	if response.Data == nil {
		return
	}

	var note dto.Note
	nilNote := dto.Note{}

	err = json.Unmarshal(response.Data, &note)
	if err != nil {
		fmt.Println(err)
		return
	}

	if note == nilNote {
		return
	}
	fmt.Println("\x1b[35mNote\x1b[0m")
	fmt.Println("\x1b[34mFirtsName:\x1b[0m ", note.AuthorFirstName)
	fmt.Println("\x1b[34mLastName:\x1b[0m ", note.AuthorLastName)
	fmt.Println("\x1b[34mNote:\x1b[0m ", note.Note)
}
