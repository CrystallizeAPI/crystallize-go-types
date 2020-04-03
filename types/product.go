package types

// Product represents a catalogue item of type "product".
type Product struct {
	CatalogueItem
	Variants []ProductVariant `json:"variants"`
}
