// Translation definition
package txtai

// Translation definition
type TranslationAPI struct {
    api API
}

// Creates a Translation instance.
func Translation(params ...string) TranslationAPI {
    return TranslationAPI{NewAPI(params...)}
}

// Translates text from source language into target language.
func (translation *TranslationAPI) Translate(text string, target string, source string) string {
    var results string

    params := map[string]string {
        "text": text,
    }

    if target != "" {
        params["target"] = target
    }
    if source != "" {
        params["source"] = source
    }

    translation.api.Get("translate", params, &results)

    return results
}

// Translates text from source language into target language.
func (translation *TranslationAPI) BatchTranslate(texts []string, target string, source string) []string {
    var results []string

    params := map[string]interface{} {
        "texts": texts,
    }

    if target != "" {
        params["target"] = target
    }
    if source != "" {
        params["source"] = source
    }

    translation.api.Post("batchtranslate", params, &results)

    return results
}
