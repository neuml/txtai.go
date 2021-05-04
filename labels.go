// Labels definition
package txtai

// Labels definition
type LabelsAPI struct {
    url string
    api API
}

// Creates a Labels instance.
func Labels(url string) LabelsAPI {
    return LabelsAPI{url, API{url}}
}

// Applies a zero shot classifier to a text using a list of labels.
func (label *LabelsAPI) Label(text string, labels []string) []IndexResult {
    var results []IndexResult

    label.api.Post("label", map[string]interface{}{
        "text":   text,
        "labels": labels,
    }, &results)

    return results
}

// Applies a zero shot classifier to list of text using a list of labels.
func (label *LabelsAPI) BatchLabel(texts []string, labels []string) []IndexResult {
    var results []IndexResult

    label.api.Post("batchlabel", map[string]interface{}{
        "texts": texts,
        "labels": labels,
    }, &results)

    return results
}
