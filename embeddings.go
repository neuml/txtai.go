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

// Search result
type SearchResult struct {
    Id    string `json:"id"`
    Score float64 `json:"score"`
}

// Creates an Embeddings instance.
func Embeddings(url string) EmbeddingsAPI {
    return EmbeddingsAPI{url, API{url}}
}

// Finds documents in the embeddings model most similar to the input query.
func (embeddings *EmbeddingsAPI) Search(query string, limit int) []SearchResult {
    var results []SearchResult

    embeddings.api.Get("search", map[string]string{
        "query": query,
        "limit": strconv.Itoa(limit),
    }, &results)

    return results;
}

// Finds documents in the embeddings model most similar to the input queries.
func (embeddings *EmbeddingsAPI) BatchSearch(query string, limit int) [][]SearchResult {
    var results [][]SearchResult

    embeddings.api.Get("batchsearch", map[string]string{
        "queries": query,
        "limit": strconv.Itoa(limit),
    }, &results)

    return results;
}

// Adds a batch of documents for indexing.
func (embeddings *EmbeddingsAPI) Add(documents interface{}) {
    embeddings.api.Post("add", documents, nil)
}

// Builds an embeddings index for previously batched documents.
func (embeddings *EmbeddingsAPI) Index() {
    embeddings.api.Get("index", nil, nil)
}

// Runs an embeddings upsert operation for previously batched documents.
func (embeddings *EmbeddingsAPI) Upsert() {
    embeddings.api.Get("upsert", nil, nil)
}

// Deletes from an embeddings index. Returns list of ids deleted.
func (embeddings *EmbeddingsAPI) Delete(ids []string) []string {
    var dids []string

    embeddings.api.Post("delete", ids, &dids)

    return dids
}

// Total number of elements in this embeddings index.
func (embeddings *EmbeddingsAPI) Count() int {
    var count int

    embeddings.api.Get("count", nil, &count)

    return count
}

// Computes the similarity between query and list of text.
func (embeddings *EmbeddingsAPI) Similarity(query string, texts []string) []IndexResult {
    var results []IndexResult

    embeddings.api.Post("similarity", map[string]interface{}{
        "query": query,
        "texts": texts,
    }, &results)

    return results
}

// Computes the similarity between list of queries and list of text.
func (embeddings *EmbeddingsAPI) BatchSimilarity(queries []string, texts []string) [][]IndexResult {
    var results [][]IndexResult

    embeddings.api.Post("batchsimilarity", map[string]interface{}{
        "queries": queries,
        "texts": texts,
    }, &results)

    return results
}

// Transforms text into an embeddings array.
func (embeddings *EmbeddingsAPI) Transform(text string) []float64 {
    var scores []float64

    embeddings.api.Get("transform", map[string]string{
        "text": text,
    }, &scores)

    return scores
}

// Transforms list of text into embeddings array.
func (embeddings *EmbeddingsAPI) BatchTransform(texts []string) [][]float64 {
    var scores [][]float64

    embeddings.api.Post("batchtransform", texts, &scores)

    return scores
}