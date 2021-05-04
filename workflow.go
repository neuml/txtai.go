// Workflow definition
package txtai

// Workflow definition
type WorkflowAPI struct {
    url string
    api API
}

// Creates a Workflow instance.
func Workflow(url string) WorkflowAPI {
    return WorkflowAPI{url, API{url}}
}

// Segments text into semantic units.
func (workflow *WorkflowAPI) Workflow(name string, elements []string) []interface{} {
    var results []interface{}

    workflow.api.Post("workflow", map[string]interface{}{
        "name": name,
        "elements": elements,
    }, &results)

    return results
}
