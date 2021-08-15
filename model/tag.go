package model

type Tag struct {
	Tag          string   `json:"tag"`
	Count        int      `json:"count"`
	Articles     []string `json:"articles"`
	Related_Tags []string `json:"related_tags"`
}
