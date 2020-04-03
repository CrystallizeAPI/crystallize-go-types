package types

// ImageVariant represents a variant for an image.
type ImageVariant struct {
	Key   string `json:"key"`
	URL   string `json:"url"`
	Width int    `json:"width"`
}
