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

func getRecordFromTerminal() (Record, int) {
	fieldCounter := 0

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Name: ")
	scanner.Scan()
	Name := scanner.Text()
	if Name != "" {
		fieldCounter++
	}

	fmt.Print("Enter LastName: ")
	scanner.Scan()
	LastName := scanner.Text()
	if LastName != "" {
		fieldCounter++
	}

	fmt.Print("Enter MiddleName: ")
	scanner.Scan()
	MiddleName := scanner.Text()
	if MiddleName != "" {
		fieldCounter++
	}

	fmt.Print("Enter Address: ")
	scanner.Scan()
	Address := scanner.Text()
	if Address != "" {
		fieldCounter++
	}

	fmt.Print("Enter Phone: ")
	scanner.Scan()
	Phone := scanner.Text()
	if Phone != "" {
		fieldCounter++
	}

	return Record{Name: Name, LastName: LastName, MiddleName: MiddleName, Address: Address, Phone: Phone}, fieldCounter
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

	rec, fieldCounter := getRecordFromTerminal()

	if fieldCounter != 5 {
		fmt.Println("Fill all fields, Abort")
		return
	}

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

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("addRecord()", err)
		return
	}
	fmt.Println("response Body:", string(body))
}

func getRecord() {
	rec, _ := getRecordFromTerminal()

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
	//req.Header.Set("X-Custom-Header", "myvalue")
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
		fmt.Println("getRecord()", err)
	}

	var records Records
	err = json.Unmarshal(body, &records)
	if err != nil {
		fmt.Println("getRecord()", err)
		return
	}

	if len(records) == 0 {
		fmt.Println("--No RECORDS finded--")
		return
	}

	fmt.Println("--GET NOTES--")
	for i, record := range records {
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
	rec, fieldCounter := getRecordFromTerminal()

	if rec.Phone == "" {
		fmt.Println("You Forgot the Phone")
		return
	}

	if fieldCounter < 2 {
		fmt.Println("Typed only Phone, there will be no changes.")
		return
	}

	jsonRequestBytes, err := json.Marshal(rec)
	if err != nil {
		fmt.Println("getRecord()", err)
	}

	url := "http://localhost:8000/update"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getRecord()", err)
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func deleteRecord() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("To delete record Enter the Phone number")
	fmt.Print("Enter Phone: ")
	scanner.Scan()
	Phone := scanner.Text()

	if Phone == "" {
		fmt.Println("You Forgot the Phone")
		return
	}

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
		fmt.Println("getRecord()", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("getRecord()", err)
	}
	fmt.Println("response Body:", string(body))
}
