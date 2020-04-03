package types

// Image represents an image media object.
type Image struct {
	Key      string         `json:"key"`
	URL      string         `json:"url"`
	Variants []ImageVariant `json:"variants"`
}
