package models

type ForeignNames struct {
	Name         string      `json:"name"`
	Text         string      `json:"text"`
	Type         string      `json:"type"`
	Flavor       string      `json:"flavor"`
	ImageUrl     string      `json:"imageUrl"`
	Language     string      `json:"language"`
	Identifiers  Identifiers `json:"identifiers"`
	Multiverseid int         `json:"multiverseid"`
}

type Identifiers struct {
	ScryfallId   string `json:"scryfallId"`
	Multiverseid int    `json:"multiverseid"`
}
