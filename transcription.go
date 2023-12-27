// Transcription definition
package txtai

// Transcription definition
type TranscriptionAPI struct {
    api API
}

// Creates a Transcription instance.
func Transcription(params ...string) TranscriptionAPI {
    return TranscriptionAPI{NewAPI(params...)}
}

// Transcribes audio files to text.
func (transcription *TranscriptionAPI) Transcribe(file string) string {
    var results string

    transcription.api.Get("transcribe", map[string]string{
        "file": file,
    }, &results)

    return results
}

// Transcribes audio files to text.
func (transcription *TranscriptionAPI) BatchTranscribe(files []string) []string {
    var results []string

    transcription.api.Post("batchtranscribe", map[string]interface{}{
        "files": files,
    }, &results)

    return results
}
