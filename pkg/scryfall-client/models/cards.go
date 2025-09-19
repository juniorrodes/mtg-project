package models

import (
	"github.com/google/uuid"
)

type Cards struct {
	Id              uuid.UUID         `json:"id"`
	ArenaId         *int              `json:"arena_id"`
	Language        string            `json:"lang"`
	MultiverseIds   *[]int            `json:"multiverse_ids"`
	Obeject         string            `json:"object"`
	PrintsSearchUri string            `json:"print_search_uri"`
	RulingsUri      string            `json:"rulings_uri"`
	Name            string            `json:"name"`
	CMC             float32           `json:"cmc"`
	Keywords        []string          `json:"keywords"`
	Legalities      map[string]string `json:"legalities"`
	TypeLine        string            `json:"type_line"`
	ImageUris       ImagesUris        `json:"image_uris"`
}

type CardSearchParams struct {
	Query               string
	Unique              *string
	Order               *string
	SortDirection       *string
	IncludeExtras       *bool
	IncludeMultilingual *bool
	IncludeVariations   *bool
	Page                *int
	Format              *string
	Pretty              *bool
}
