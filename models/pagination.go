package models

type Pagination struct {
	Total        int64             `json:"total"`
	PerPage      int               `json:"per_page"`
	CurrentPage  int               `json:"current_page"`
	LastPage     int               `json:"last_page"`
	FirstPageURL string            `json:"first_page_url"`
	PrevPageURL  string            `json:"prev_page_url"`
	NextPageURL  string            `json:"next_page_url"`
	LastPageURL  string            `json:"last_page_url"`
	From         int64             `json:"from"`
	To           int64             `json:"to"`
	Items        interface{}       `json:"items"`
	Filters      FiltersPagination `json:"filters"`
}

type FiltersPagination struct {
	// ONLY STRING TYPE FOR DYNAMIC SET VALUE
	Page    string `json:"page"`
	PerPage string `json:"per_page"`
	SortBy  string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
	Search  string `json:"search"`


	// Other filter may be available dynamic base on table field must be add manually.
	Source string `json:"source,omitempty"`

	// Used on filter registrations.
	Status     string `json:"status,omitempty"`
	CreateFrom string `json:"create_from,omitempty"`
	CreateTo   string `json:"create_to,omitempty"`

	// Used on filter dashboard user
	Role string `json:"role,omitempty"`
}
