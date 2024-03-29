// Summary definition
package txtai

import (
    "strconv"
)

// Summary definition
type SummaryAPI struct {
    api API
}

// Creates a Summary instance.
func Summary(params ...string) SummaryAPI {
    return SummaryAPI{NewAPI(params...)}
}

// Runs a summarization model against a block of text.
func (summary *SummaryAPI) Summary(text string, minlength int, maxlength int) string {
    var results string

    params := map[string]string {
        "text": text,
    }

    if minlength != -1 {
        params["minlength"] = strconv.Itoa(minlength)
    }
    if maxlength != -1 {
        params["maxlength"] = strconv.Itoa(maxlength)
    }

    summary.api.Get("summary", params, &results)

    return results
}

// Runs a summarization model against a block of text.
func (summary *SummaryAPI) BatchSummary(texts []string, minlength int, maxlength int) []string {
    var results []string

    params := map[string]interface{} {
        "texts": texts,
    }

    if minlength != -1 {
        params["minlength"] = minlength
    }
    if maxlength != -1 {
        params["maxlength"] = maxlength
    }	

    summary.api.Post("batchsummary", params, &results)

    return results
}
