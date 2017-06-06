package article

import (
	"net/http"

	articleUcase "github.com/bxcodec/Article-Management-Clean-Architecture/usecase/article"
	"github.com/labstack/echo"
)

type ArticleHandler struct {
	AUsecase articleUcase.ArticleUsecase
}

func (a *ArticleHandler) FetchArticle(c echo.Context) error {

	listAr, err := a.AUsecase.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something Error On Our Services")
	}

	return c.JSON(http.StatusOK, listAr)
}
