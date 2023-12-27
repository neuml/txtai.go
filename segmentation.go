// Segmentation definition
package txtai

// Segmentation definition
type SegmentationAPI struct {
    api API
}

// Creates a Segmentation instance.
func Segmentation(params ...string) SegmentationAPI {
    return SegmentationAPI{NewAPI(params...)}
}

// Segments text into semantic units.
func (segmentation *SegmentationAPI) Segment(text string) interface{} {
    var results interface{}

    segmentation.api.Get("segment", map[string]string{
        "text": text,
    }, &results)

    return results
}

// Segments text into semantic units.
func (segmentation *SegmentationAPI) BatchSegment(texts []string) []interface{} {
    var results []interface{}

    segmentation.api.Post("batchsegment", map[string]interface{}{
        "texts": texts,
    }, &results)

    return results
}
