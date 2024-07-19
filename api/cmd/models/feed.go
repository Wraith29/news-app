package models

type Feed struct {
	id              int
	Author, FeedUrl string
}

func NewFeed(id int, author, feedUrl string) Feed {
	return Feed{
		id, author, feedUrl,
	}
}
