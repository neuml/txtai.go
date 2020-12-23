// Extractor definition
package txtai

// Extractor definition
type ExtractorAPI struct {
	url string
	api API
}

// Text section
type Section struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
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
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// Creates an Extractor instance.
func Extractor(url string) ExtractorAPI {
	return ExtractorAPI{url, API{url}}
}

// Extracts answers to input questions.
func (extractor *ExtractorAPI) Extract(documents []Section, queue []Question) []Answer {
	var data [][]interface{}

	extractor.api.Post("extract", map[string]interface{}{
		"documents": documents,
		"queue":     queue,
	}, &data)

	// Transform arrays into answers
	var answers []Answer
	for x := range data {
		answers = append(answers, Answer{data[x][0].(string), data[x][1].(string)})
	}

	return answers
}
