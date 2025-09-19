package models

import "fmt"

type ScryfallError struct {
	Status   int      `json:"status"`
	Code     string   `json:"code"`
	Details  string   `json:"details"`
	Type     string   `json:"type"`
	Warnings []string `json:"warnings"`
}

func (e ScryfallError) Error() string {
	return fmt.Sprintf("Failed with details: %s", e.Details)
}
