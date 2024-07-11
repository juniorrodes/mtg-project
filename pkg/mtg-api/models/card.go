package models

type Cards struct {
	Cards []Card `json:"cards"`
}

type Card struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	ManaCost      string         `json:"manaCost"`
	Cmc           float64        `json:"cmc"`
	Colors        []string       `json:"colors"`
	ColorIdentity []string       `json:"colorIdentity"`
	Type          string         `json:"type"`
	SubTypes      []string       `json:"subtypes"`
	Types         []string       `json:"types"`
	Rarity        string         `json:"rarity"`
	Set           string         `json:"set"`
	SetName       string         `json:"setName"`
	Text          string         `json:"text"`
	Aritst        string         `json:"artist"`
	Number        string         `json:"number"`
	Power         string         `json:"power"`
	Toughness     string         `json:"toughness"`
	Layout        string         `json:"layout"`
	Multiverseid  string         `json:"multiverseid"`
	ImageUrl      string         `json:"imageUrl"`
	Variations    []string       `json:"variations"`
	Printings     []string       `json:"printings"`
	OriginalText  string         `json:"originalText"`
	OriginalType  string         `json:"originalType"`
	ForeignNames  []ForeignNames `json:"foreignNames"`
	Legalities    []Legalities   `json:"legalities"`
}

type Legalities struct {
	Format   string `json:"format"`
	Legality string `json:"legality"`
}
