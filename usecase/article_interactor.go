package usecase

import "github.com/Le0tk0k/go-rest-api/domain"

type ArticleInteractor struct {
	ArticleRepository ArticleRepository
}

func (interactor *ArticleInteractor) ArticleById(id int) (article *domain.Article, err error) {
	article, err = interactor.ArticleRepository.FindByID(id)
	return
}

func (interactor *ArticleInteractor) Articles() (articles *domain.Article, err error) {
	articles, err = interactor.ArticleRepository.FindAll()
	return
}

func (interactor *ArticleInteractor) Add(article domain.Article) (article *domain.Article, err error) {
	article, err = interactor.ArticleRepository.Store(article)
	return
}

func (interactor *ArticleInteractor) Update(article domain.Article) (article *domain.Article, err error) {
	article, err = interactor.ArticleRepository.Update(article)
	return
}

func (interactor *ArticleInteractor) DeleteById(article domain.Article) (err error) {
	err = interactor.ArticleRepository.DeleteById(article)
	return
}
