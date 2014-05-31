package main

import "fmt"
import "os"
import "github.com/stefankopieczek/godocfind"

func main() {
    searchTerm := os.Args[1]
    fmt.Printf("Searching for keyword '%s'...  ", searchTerm)
    results, err := godocfind.ByKeyword(searchTerm)

    if (err != nil) {
        fmt.Printf("An error occurred: %s", err.Error())
    } else if len(results) == 0 {
        fmt.Printf("no results\n\n")
    } else {
        fmt.Printf("%d result(s) found\n\n", len(results))
        showResults(results)
    }
}

func showResults(results []string) {
    for _, result := range(results) {
        fmt.Printf("%s\n\n", result)
    }
}
