package main

import (
	"fmt"

	"gopkg.in/resty.v1"
)

var url = "http://localhost:3000/members"

func main() {
	testPostMembers()
	fmt.Println("----------------------------")
	testGetMembers()
}

//using default methods in go
func testGetMembers() {
	fmt.Println("Consultando API Members /GET")
	resp, err := resty.R().Get(url)
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())
}

//using default methods in go
func testPostMembers() {
	fmt.Println("Consultando API Members /POST")
	jsonData := map[string]string{"id": "1", "name": "Yago Marques", "email": "yagomarquesja@gmail.com", "cellphone": "938298922"}

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonData).
		Post(url)

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Println("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())
}
