// Embeddings definition
package txtai

import "strconv"

// Embeddings definition
type EmbeddingsAPI struct {
	url string
	api API
}

// Base input Document
type Document struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

// API score result
type Score struct {
	Id    string  `json:"id"`
	Score float64 `json:"score"`
}

// Creates an Embeddings instance.
func Embeddings(url string) EmbeddingsAPI {
	return EmbeddingsAPI{url, API{url}}
}

// Runs an Embeddings search. Returns []Score.
func (embeddings *EmbeddingsAPI) Search(q string, n int) []Score {
	var data [][]interface{}

	embeddings.api.Get("search", map[string]string{
		"q": q,
		"n": strconv.Itoa(n),
	}, &data)

	// Transform arrays into scores
	var scores []Score
	for x := range data {
		scores = append(scores, Score{data[x][0].(string), data[x][1].(float64)})
	}

	return scores
}

// Adds a batch of documents for indexing.
func (embeddings *EmbeddingsAPI) Add(documents interface{}) {
	embeddings.api.Post("add", documents, nil)
}

// Builds an embeddings index. No further documents can be added after this call.
func (embeddings *EmbeddingsAPI) Index() {
	embeddings.api.Get("index", nil, nil)
}

// Calculates the similarity between search and list of elements in data.
func (embeddings *EmbeddingsAPI) Similarity(search string, data []string) []float64 {
	var scores []float64

	embeddings.api.Post("similarity", map[string]interface{}{
		"search": search,
		"data":   data,
	}, &scores)

	return scores
}

// Transforms text into an embeddings array.
func (embeddings *EmbeddingsAPI) Embeddings(t string) []float64 {
	var scores []float64

	embeddings.api.Get("embeddings", map[string]string{
		"t": t,
	}, &scores)

	return scores
}
