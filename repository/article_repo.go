package repository

import "github.com/bxcodec/Article-Management-Clean-Architecture/models"

type ArticleRepository interface {
	Fetch(cursor string, num int64) ([]*models.Article, error)
}
