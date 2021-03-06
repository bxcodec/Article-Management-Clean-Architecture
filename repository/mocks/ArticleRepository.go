package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/bxcodec/Article-Management-Clean-Architecture/models"

// ArticleRepository is an autogenerated mock type for the ArticleRepository type
type ArticleRepository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: cursor, num
func (_m *ArticleRepository) Fetch(cursor string, num int64) ([]*models.Article, error) {
	ret := _m.Called(cursor, num)

	var r0 []*models.Article
	if rf, ok := ret.Get(0).(func(string, int64) []*models.Article); ok {
		r0 = rf(cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(cursor, num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
