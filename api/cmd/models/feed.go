package models

import "fmt"

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

func (f *Feed) String() string {
	return fmt.Sprintf("%d:%s:%s", f.Id, f.Author, f.FeedUrl)
}
