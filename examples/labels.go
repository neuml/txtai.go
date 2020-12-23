// Example labels functionality.
//
// Implements logic found in this notebook: https://github.com/neuml/txtai/blob/master/examples/05_Labeling_with_zero_shot_classification.ipynb
package main

import (
	"fmt"
	"strings"
	"github.com/neuml/txtai.go"
)

func main() {
	labels := txtai.Labels("http://localhost:8000")

	sections := []string{"Dodgers lose again, give up 3 HRs in a loss to the Giants",
		"Giants 5 Cardinals 4 final in extra innings",
		"Dodgers drop Game 2 against the Giants, 5-4",
		"Flyers 4 Lightning 1 final. 45 saves for the Lightning.",
		"Slashing, penalty, 2 minute power play coming up",
		"What a stick save!",
		"Leads the NFL in sacks with 9.5",
		"UCF 38 Temple 13",
		"With the 30 yard completion, down to the 10 yard line",
		"Drains the 3pt shot!!, 0:15 remaining in the game",
		"Intercepted! Drives down the court and shoots for the win",
		"Massive dunk!!! they are now up by 15 with 2 minutes to go"}

	// List of labels
	tags := []string{"Baseball", "Football", "Hockey", "Basketball"}

	fmt.Printf("%-75s %s\n", "Text", "Label")
	fmt.Println(strings.Repeat("-", 100))

	for _, text := range sections {
		scores := labels.Label(text, tags)
		fmt.Printf("%-75s %s\n", text, scores[0].Id)
	}

	fmt.Println()
	fmt.Printf("%-75s %s\n", "Text", "Label")
	fmt.Println(strings.Repeat("-", 100))

	tags = []string{"😀", "😡"}

	for _, text := range sections {
		scores := labels.Label(text, tags)
		fmt.Printf("%-75s %s\n", text, scores[0].Id)
	}
}
