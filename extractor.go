// Extractor definition
package txtai

// Extractor definition
type ExtractorAPI struct {
    api API
}

// Question parameters
type Question struct {
    Name     string `json:"name"`
    Query    string `json:"query"`
    Question string `json:"question"`
    Snippet  bool   `json:"snippet"`
}

// Answer response
type Answer struct {
    Name string `json:"name"`
    Answer   string `json:"answer"`
}

// Creates an Extractor instance.
func Extractor(params ...string) ExtractorAPI {
    return ExtractorAPI{NewAPI(params...)}
}

// Extracts answers to input questions.
func (extractor *ExtractorAPI) Extract(queue []Question, texts []string) []Answer {
    var answers []Answer

    extractor.api.Post("extract", map[string]interface{}{
        "queue": queue,
        "texts": texts,
    }, &answers)

    return answers
}
