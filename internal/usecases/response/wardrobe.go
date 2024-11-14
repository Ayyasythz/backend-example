package response

type WardrobeResponse struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Color string  `json:"color,omitempty"`
	Size  string  `json:"size,omitempty"`
	Price float32 `json:"price,omitempty"`
	Stock int     `json:"stock,omitempty"`
}
