package Models

import (
	"errors"
	"time"
)

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"timestamp"`
}

func GetAllArticles() []Article {
	articles := []Article{}
	articles = append(articles, Article{ID: 1, Title: "This is the First Article", Description: "This might be the first description of the articles but not the last one just fyi.", CreatedAt: time.Now()})
	articles = append(articles, Article{ID: 2, Title: "This is the Second Article", Description: "This might be the Second description of the articles but not the last one just fyi.", CreatedAt: time.Now()})
	articles = append(articles, Article{ID: 3, Title: "This is the Third Article", Description: "This might be the Third description of the articles but not the last one just fyi.", CreatedAt: time.Now()})
	articles = append(articles, Article{ID: 4, Title: "This is the Fourth Article", Description: "This might be the Fourth description of the articles but not the last one just fyi.", CreatedAt: time.Now()})
	articles = append(articles, Article{ID: 5, Title: "This is the Fifth Article", Description: "This might be the Fifth description of the articles but not the last one just fyi.", CreatedAt: time.Now()})
	return articles
}

func GetArticleByID(id int) (*Article, error) {
	articles := GetAllArticles()
	for _, article := range articles {
		if article.ID == id {
			return &article, nil
		}
	}
	return nil, errors.New("Article Not Found")
}
