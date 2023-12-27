// Textractor definition
package txtai

// Textractor definition
type TextractorAPI struct {
    api API
}

// Creates a Textractor instance.
func Textractor(params ...string) TextractorAPI {
    return TextractorAPI{NewAPI(params...)}
}

// Extracts text from a file at path.
func (textractor *TextractorAPI) Textract(file string) interface{} {
    var results interface{}

    textractor.api.Get("textract", map[string]string{
        "file": file,
    }, &results)

    return results
}

// Extracts text from a file at path.
func (textractor *TextractorAPI) BatchTextract(files []string) []interface{} {
    var results []interface{}

    textractor.api.Post("batchtextract", map[string]interface{}{
        "files": files,
    }, &results)

    return results
}
