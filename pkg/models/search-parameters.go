package models

type SingleParam struct {
	Value string `json:"value"`
}

type MultiParams struct {
	Value   []string `json:"value"`
	Operand string   `json:"operand"`
}

type SearchParameters struct {
	Name     *SingleParam `json:"name"`
	Color    *MultiParams `json:"color"`
	PageSize SingleParam  `json:"pageSize"`
}
