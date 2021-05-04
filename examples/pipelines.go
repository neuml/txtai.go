// Example pipeline and workflow functionality.
//
// Uses files from txtai unit tests: https://github.com/neuml/txtai/releases/download/v2.0.0/tests.tar.gz
package main

import (
    "fmt"
    "github.com/neuml/txtai.go"
)

func main() {
    service := "http://localhost:8000"

    segment := txtai.Segmentation(service)

    sentences := "This is a test. And another test."

    fmt.Println("---- Segmented Text ----")
    fmt.Println(segment.Segment(sentences))

    textractor := txtai.Textractor(service)
    text := textractor.Textract("/tmp/txtai/article.pdf");

    fmt.Println("\n---- Extracted Text ----")
    fmt.Println(text)

    summary := txtai.Summary(service)
    summarytext :=  summary.Summary(text.(string), -1, -1)

    fmt.Println("\n---- Summary Text ----")
    fmt.Println(summarytext)

    translate := txtai.Translation(service)
    translation := translate.Translate(summarytext, "es", "")

    fmt.Println("\n---- Summary Text in Spanish ----")
    fmt.Println(translation)

    workflow := txtai.Workflow(service)
    output := workflow.Workflow("sumspanish", []string {"file:///tmp/txtai/article.pdf"})

    fmt.Println("\n---- Workflow [Extract Text->Summarize->Translate] ----")
    fmt.Println(output)

    transcribe := txtai.Transcription(service)
    transcription := transcribe.Transcribe("/tmp/txtai/Make_huge_profits.wav")

    fmt.Println("\n---- Transcribed Text ----")
    fmt.Println(transcription)
}
