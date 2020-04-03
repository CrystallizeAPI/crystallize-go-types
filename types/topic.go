package types

// Topic represents a Crystallize topic.
type Topic struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parentId"`
}
