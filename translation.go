// Translation definition
package txtai

// Translation definition
type TranslationAPI struct {
	url string
	api API
}

// Creates a Translation instance.
func Translation(url string) TranslationAPI {
	return TranslationAPI{url, API{url}}
}

// Translates text from source language into target language.
func (translation *TranslationAPI) Translate(text string, target string, source string) interface{} {
	var results interface{}

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
func (translation *TranslationAPI) BatchTranslate(texts []string, target string, source string) []interface{} {
	var results []interface{}

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
