package models

type Feed struct {
	Id      int    `json:"id"`
	Author  string `json:"author"`
	FeedUrl string `json:"feedUrl"`
}

func NewFeed(id int, author, feedUrl string) Feed {
	return Feed{
		id, author, feedUrl,
	}
}
