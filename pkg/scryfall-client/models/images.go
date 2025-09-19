package models

type ImagesUris struct {
	PNG        string `json:"png"`
	BorderCrop string `json:"border_crop"`
	ArtCrop    string `json:"art_crop"`
	Large      string `json:"large"`
	Normal     string `json:"normal"`
	Small      string `json:"small"`
}
