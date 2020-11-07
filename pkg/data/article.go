package data

import (
	"errors"
	"time"
)

type Article struct {
	ID        uint64
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var articles = map[uint64]Article{
	uint64(1): {
		ID:        uint64(1),
		Title:     "title1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	uint64(2): {
		ID:        uint64(2),
		Title:     "title2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	uint64(3): {
		ID:        uint64(3),
		Title:     "title3",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

type ArticleHandler struct{}

func NewArticleHandler() ArticleHandler {
	return ArticleHandler{}
}

func (h ArticleHandler) Find(id uint64) (a Article, err error) {
	a, ok := articles[id]
	if !ok {
		return a, errors.New("not found")
	}
	return a, nil
}

func (h ArticleHandler) FindAll() (as []Article, err error) {
	for _, a := range articles {
		as = append(as, a)
	}
	return as, nil
}

func (h ArticleHandler) Create(inp Article) (out Article, err error) {
	id := uint64(len(articles) + 1)
	out = Article{
		ID:        id,
		Title:     inp.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	articles[id] = out
	return out, nil
}
