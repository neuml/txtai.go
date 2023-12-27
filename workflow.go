// Workflow definition
package txtai

// Workflow definition
type WorkflowAPI struct {
    api API
}

// Creates a Workflow instance.
func Workflow(params ...string) WorkflowAPI {
    return WorkflowAPI{NewAPI(params...)}
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
