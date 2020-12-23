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

// Applies a zero shot classifier to a text section using a list of labels.
func (label *LabelsAPI) Label(text string, labels []string) []Score {
	var data [][]interface{}

	label.api.Post("label", map[string]interface{}{
		"text":   text,
		"labels": labels,
	}, &data)

	// Transform arrays into scores
	var scores []Score
	for x := range data {
		scores = append(scores, Score{data[x][0].(string), data[x][1].(float64)})
	}

	return scores
}
