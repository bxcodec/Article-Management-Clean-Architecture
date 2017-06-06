package article

import "github.com/bxcodec/Article-Management-Clean-Architecture/models"
import "github.com/bxcodec/Article-Management-Clean-Architecture/repository"

type ArticleUsecase interface {
	Fetch() ([]*models.Article, error)
}

type articleUsecase struct {
	repos repository.ArticleRepository
}

func (a *articleUsecase) Fetch() ([]*models.Article, error) {
	return a.repos.Fetch()
}

func NewArticleUsecase(repos repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{repos}
}
