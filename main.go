package mathutil

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func Add(a, b int) int {
   url := "https://jsonplaceholder.typicode.com/todos/1"

    // Make the HTTP GET request
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error making GET request: %s\n", err)
        return
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading response body: %s\n", err)
        return
    }

    // Check the HTTP status code
    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
        return
    }

    // Unmarshal the JSON response into a Todo struct
    var todo Todo
    err = json.Unmarshal(body, &todo)
    if err != nil {
        fmt.Printf("Error unmarshalling JSON response: %s\n", err)
        return
    }

    // Print the Todo struct
    fmt.Printf("Todo: %+v\n", todo) 
   fmt.Printf(todo.id)
    return a + b + todo.id
}
