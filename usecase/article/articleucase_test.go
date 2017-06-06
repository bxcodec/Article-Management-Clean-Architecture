package article_test

import (
	"testing"

	"github.com/bxcodec/Article-Management-Clean-Architecture/models"
	"github.com/bxcodec/Article-Management-Clean-Architecture/repository/mocks"
	ucase "github.com/bxcodec/Article-Management-Clean-Architecture/usecase/article"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	mockRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListArtilce, nil)
	u := ucase.NewArticleUsecase(mockRepo)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(cursor, num)
	assert.Equal(t, "10", nextCursor)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArtilce))
	mockRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))
}
