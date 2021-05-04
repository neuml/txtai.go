// Transcription definition
package txtai

// Transcription definition
type TranscriptionAPI struct {
	url string
	api API
}

// Creates a Transcription instance.
func Transcription(url string) TranscriptionAPI {
	return TranscriptionAPI{url, API{url}}
}

// Transcribes audio files to text.
func (transcription *TranscriptionAPI) Transcribe(file string) interface{} {
	var results interface{}

	transcription.api.Get("transcribe", map[string]string{
		"file": file,
	}, &results)

	return results
}

// Transcribes audio files to text.
func (transcription *TranscriptionAPI) BatchTranscribe(files []string) []interface{} {
	var results []interface{}

	transcription.api.Post("batchtranscribe", map[string]interface{}{
		"files": files,
	}, &results)

	return results
}
