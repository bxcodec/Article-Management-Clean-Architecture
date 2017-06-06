package repository

import "github.com/bxcodec/Article-Management-Clean-Architecture/models"

type ArticleRepository interface {
	Fetch() ([]*models.Article, error)
}
