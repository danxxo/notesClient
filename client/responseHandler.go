package client

import (
	"encoding/json"
	"fmt"

	dto "notesClient/models/dto"
)

func responseHandler(body []byte) {
	fmt.Println("--Response--")

	var response dto.Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\tResult: ", response.Result)
	if response.Result == "error" {
		fmt.Println("\tserver returned error: ", response.Error)
		fmt.Println("--End Response--")
		return
	}
	fmt.Println("--End Response--")

	if response.Data == nil {
		return
	}

	var notes dto.Notes
	err = json.Unmarshal(response.Data, &notes)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(notes) == 0 {
		return
	}

	fmt.Println("--Get Notes--")
	for i, record := range notes {
		fmt.Println("\tRecord ", i)
		fmt.Println("\t\tFirtsName: ", record.AuthorFirstName)
		fmt.Println("\t\tLastName: ", record.AuthorLastName)
		fmt.Println("\t\tNote: ", record.Note)
	}
	fmt.Println("--END--")
}
