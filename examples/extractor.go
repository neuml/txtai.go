// Example extractor functionality.
//
// Implements logic found in this notebook: https://github.com/neuml/txtai/blob/master/examples/05_Extractive_QA_with_txtai.ipynb
package main

import (
    "fmt"
    "github.com/neuml/txtai.go"
)

func main() {
    extractor := txtai.Extractor("http://localhost:8000")

    data := []string{"Giants hit 3 HRs to down Dodgers",
        "Giants 5 Dodgers 4 final",
        "Dodgers drop Game 2 against the Giants, 5-4",
        "Blue Jays 2 Red Sox 1 final",
        "Red Sox lost to the Blue Jays, 2-1",
        "Blue Jays at Red Sox is over. Score: 2-1",
        "Phillies win over the Braves, 5-0",
        "Phillies 5 Braves 0 final",
        "Final: Braves lose to the Phillies in the series opener, 5-0",
        "Final score: Flyers 4 Lightning 1",
        "Flyers 4 Lightning 1 final",
        "Flyers win 4-1"}

    // Run series of questions
    questions := []string{"What team won the game?", "What was score?"}
    for _, query := range []string{"Red Sox - Blue Jays", "Phillies - Braves", "Dodgers - Giants", "Flyers - Lightning"} {
        fmt.Println("----" + query + "----")

        var queue []txtai.Question
        for _, question := range questions {
            queue = append(queue, txtai.Question{
                Name:     question,
                Query:    query,
                Question: question,
                Snippet:  false,
            })
        }

        for _, answer := range extractor.Extract(queue, data) {
            fmt.Println(answer.Name, answer.Answer)
        }
        fmt.Println()
    }

    // Ad-hoc questions
    question := "What hockey team won?"

    fmt.Println("----" + question + "----")
    queue := []txtai.Question{txtai.Question{
        Name:     question,
        Query:    question,
        Question: question,
        Snippet:  false,
    }}

    for _, answer := range extractor.Extract(queue, data) {
        fmt.Println(answer.Name, answer.Answer)
    }
}
