package article_test

import (
	"testing"

	"github.com/bxcodec/Article-Management-Clean-Architecture/models"
	"github.com/bxcodec/Article-Management-Clean-Architecture/repository/mocks"
	ucase "github.com/bxcodec/Article-Management-Clean-Architecture/usecase/article"
	"github.com/stretchr/testify/assert"
)

var (
	mockArticle = &models.Article{
		ID:      int64(10),
		Title:   "Cinta Buta",
		Content: "Jatuh Cinta Membunuhku",
	}
)

func TestFetch(t *testing.T) {
	mockRepo := new(mocks.ArticleRepository)
	mockListArtilce := make([]*models.Article, 0)
	mockListArtilce = append(mockListArtilce, mockArticle)
	mockRepo.On("Fetch").Return(mockListArtilce, nil)
	u := ucase.NewArticleUsecase(mockRepo)
	list, err := u.Fetch()

	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArtilce))
	mockRepo.AssertCalled(t, "Fetch")

}
