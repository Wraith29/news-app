package models

import (
	"crypto/sha256"
	"fmt"
	"time"
)

const dateFormat = "15:04:05"

type Article struct {
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	Link            string     `json:"link"`
	PublishedParsed *time.Time `json:"publishedParsed"`
	Author          string     `json:"author"`
}

func (a *Article) String() string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", a.Title, a.Description, a.Link, a.PublishedParsed.Format(dateFormat), a.Author)
}

func (a *Article) Bytes() []byte {
	return []byte(a.String())
}

type ArticleList []Article

func (f ArticleList) Len() int {
	return len(f)
}

func (f ArticleList) Less(i, j int) bool {
	return f[j].PublishedParsed.Before(*f[i].PublishedParsed)
}

func (f ArticleList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f ArticleList) Hash() []byte {
	hash := sha256.New()
	sum := make([]byte, 0)

	for _, article := range f {
		hash.Write(article.Bytes())

		for _, b := range hash.Sum(nil) {
			sum = append(sum, b)
		}
	}

	return sum
}
