// Example embeddings functionality.
//
// Implements functionality found in this notebook: https://github.com/neuml/txtai/blob/master/examples/01_Introducing_txtai.ipynb
package main

import (
    "fmt"
    "strconv"
    "strings"
    "github.com/neuml/txtai.go"
)

func main() {
    embeddings := txtai.Embeddings("http://localhost:8000")

    data := []string{"US tops 5 million confirmed virus cases",
        "Canada's last fully intact ice shelf has suddenly collapsed, forming a Manhattan-sized iceberg",
        "Beijing mobilises invasion craft along coast as Taiwan tensions escalate",
        "The National Park Service warns against sacrificing slower friends in a bear attack",
        "Maine man wins $1M from $25 lottery ticket",
        "Make huge profits without work, earn up to $100,000 a day"}

    var documents []txtai.Document
    for x := range data {
        documents = append(documents, txtai.Document{
            Id:   strconv.Itoa(x),
            Text: data[x],
        })
    }

    fmt.Println("Running similarity queries")
    fmt.Printf("%-20s %s\n", "Query", "Best Match")
    fmt.Println(strings.Repeat("-", 50))

    for _, query := range []string{"feel good story", "climate change", "health", "war", "wildlife", "asia", "north america", "dishonest junk"} {
        results := embeddings.Similarity(query, data)
        fmt.Printf("%-20s %s\n", query, data[results[0].Id])
    }

    embeddings.Add(documents)
    embeddings.Index()

    fmt.Println("\nBuilding an Embeddings index")
    fmt.Printf("%-20s %s\n", "Query", "Best Match")
    fmt.Println(strings.Repeat("-", 50))

    for _, query := range []string{"feel good story", "climate change", "health", "war", "wildlife", "asia", "north america", "dishonest junk"} {
        results := embeddings.Search(query, 1)
        uid, _ := strconv.Atoi(results[0].Id)
        fmt.Printf("%-20s %s\n", query, data[uid])
    }
}
