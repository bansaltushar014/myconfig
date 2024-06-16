package mathutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Add(a, b int) int {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %s\n", err)
		return 1
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return 1
	}

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return 1
	}

	// Unmarshal the JSON response into a Todo struct
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON response: %s\n", err)
		return 1
	}

	// Print the Todo struct
	fmt.Printf("Todo: %+v\n", todo)
	fmt.Println(todo.ID)
	return a + 2 + b + todo.ID
}
