package article

import (
	"strconv"

	"github.com/bxcodec/Article-Management-Clean-Architecture/models"
	"github.com/bxcodec/Article-Management-Clean-Architecture/repository"
)

type ArticleUsecase interface {
	Fetch(cursor string, num int64) ([]*models.Article, string, error)
}

type articleUsecase struct {
	repos repository.ArticleRepository
}

func (a *articleUsecase) Fetch(cursor string, num int64) ([]*models.Article, string, error) {
	if num == 0 {
		num = 10
	}

	listArticle, err := a.repos.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}
	nextCursor := ""

	if size := len(listArticle); size == int(num) {
		lastId := listArticle[num-1].ID
		nextCursor = strconv.Itoa(int(lastId))
	}
	return listArticle, nextCursor, nil
}

func NewArticleUsecase(repos repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{repos}
}
