package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Record struct {
	ID         int64  `json:"-" sql.field:"id"`
	Name       string `json:"name" sql.field:"name"`
	LastName   string `json:"last_name" sql.field:"lastname"`
	MiddleName string `json:"middle_name" sql.field:"middlename"`
	Address    string `json:"address" sql.field:"address"`
	Phone      string `json:"phone" sql.field:"phone"`
}

type Records []Record

type Response struct {
	Records      []Record `json:"records"`
	ErrorMessage string   `json:"err"`
}

// func for scanning terminal input, reurns the gotten record and num of fields
func getRecordFromTerminal() Record {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Name: ")
	scanner.Scan()
	Name := scanner.Text()

	fmt.Print("Enter LastName: ")
	scanner.Scan()
	LastName := scanner.Text()

	fmt.Print("Enter MiddleName: ")
	scanner.Scan()
	MiddleName := scanner.Text()

	fmt.Print("Enter Address: ")
	scanner.Scan()
	Address := scanner.Text()

	fmt.Print("Enter Phone: ")
	scanner.Scan()
	Phone := scanner.Text()

	return Record{Name: Name, LastName: LastName, MiddleName: MiddleName, Address: Address, Phone: Phone}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Create a Record")
		fmt.Println("2. Get a Record")
		fmt.Println("3. Update record by phone. Phone is NEEDABLE!")
		fmt.Println("4. Delete record by Phone. Phone is NEEDABLE!")
		fmt.Println("5. Exit")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			addRecord()
		case "2":
			getRecord()
		case "3":
			updateNotes()
		case "4":
			deleteRecord()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}

		fmt.Println()
	}
}

func addRecord() {
	fmt.Println("Lets create a new Record!")

	rec := getRecordFromTerminal()

	jsonRequestBytes, err := json.Marshal(rec)
	if err != nil {
		fmt.Println("addRecord()", err)
	}

	url := "http://localhost:8000/add"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("addRecord()", err)
	}

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("recordAdd()", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("addRecord()", err)
		return
	}
	fmt.Println("response Body:", string(body))
}

func getRecord() {
	rec := getRecordFromTerminal()

	jsonRequestBytes, err := json.Marshal(rec)
	if err != nil {
		fmt.Println("getRecord()", err)
		return
	}

	url := "http://localhost:8000/get"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getRecord()", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("getRecord()", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.ErrorMessage != "" {
		fmt.Println("error from server: ", response.ErrorMessage)
		return
	}

	fmt.Println("--GET NOTES--")
	for i, record := range response.Records {
		fmt.Println("\tRecord ", i)
		fmt.Println("\t\tName: ", record.Name)
		fmt.Println("\t\tLastName: ", record.LastName)
		fmt.Println("\t\tLastName: ", record.MiddleName)
		fmt.Println("\t\tAddress: ", record.Address)
		fmt.Println("\t\tPhone: ", record.Phone)
	}
	fmt.Println("--END--")

}

func updateNotes() {
	fmt.Println("Let s Update records. Dont forget fill Phone field")
	rec := getRecordFromTerminal()

	jsonRequestBytes, err := json.Marshal(rec)
	if err != nil {
		fmt.Println("getRecord()", err)
		return
	}

	url := "http://localhost:8000/update"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getRecord()", err)
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
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.ErrorMessage != "" {
		fmt.Println("error from server: ", response.ErrorMessage)
		return
	}

}

func deleteRecord() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("To delete record Enter the Phone number")
	fmt.Print("Enter Phone: ")
	scanner.Scan()
	Phone := scanner.Text()

	jsonRequestBytes, err := json.Marshal(Record{Phone: Phone})
	if err != nil {
		fmt.Println("getRecord()", err)
	}

	url := "http://localhost:8000/delete"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getRecord()", err)
	}
	req.Header.Set("X-Custom-Header", "myvalue")
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
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.ErrorMessage != "" {
		fmt.Println("error from server: ", response.ErrorMessage)
		return
	}
}
