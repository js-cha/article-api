package model

type Article struct {
	ID    int64    `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

type Tag struct {
	Tag          string   `json:"tag"`
	Count        int      `json:"count"`
	Articles     []string `json:"articles"`
	Related_Tags []string `json:"related_tags"`
}
