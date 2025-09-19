package models

type List[T any] struct {
	Object     string    `json:"object"`
	Data       []T       `json:"data"`
	HasMore    bool      `json:"has_more"`
	NextPage   *string   `json:"next_page"`
	TotalCards *int      `json:"total_cards"` // If this is a list of cards, this will be the total number of cards
	Warnings   *[]string `json:"warnings"`
}
