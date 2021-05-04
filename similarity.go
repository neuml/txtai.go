// Similarity definition
package txtai

// Similarity definition
type SimilarityAPI struct {
    url string
    api API
}

// Creates a Similarity instance.
func Similarity(url string) SimilarityAPI {
    return SimilarityAPI{url, API{url}}
}

// Computes the similarity between query and list of text.
func (similarity *SimilarityAPI) Similarity(query string, texts []string) []IndexResult {
    var results []IndexResult

    similarity.api.Post("similarity", map[string]interface{}{
        "query": query,
        "texts": texts,
    }, &results)

    return results
}

// Computes the similarity between list of queries and list of text.
func (similarity *SimilarityAPI) BatchSimilarity(queries []string, texts []string) [][]IndexResult {
    var results [][]IndexResult

    similarity.api.Post("batchsimilarity", map[string]interface{}{
        "queries": queries,
        "texts": texts,
    }, &results)

    return results
}
