package usecase

import "github.com/Le0tk0k/go-rest-api/domain"

type ArticleInteractor struct {
	ArticleRepository ArticleRepository
}

func (interactor *ArticleInteractor) ArticleById(id int) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.FindByID(id)
	return
}

func (interactor *ArticleInteractor) Articles() (articles domain.Articles, err error) {
	articles, err = interactor.ArticleRepository.FindAll()
	return
}

func (interactor *ArticleInteractor) Add(a domain.Article) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.Store(a)
	return
}

func (interactor *ArticleInteractor) Update(a domain.Article) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.Update(a)
	return
}

func (interactor *ArticleInteractor) DeleteById(article domain.Article) (err error) {
	err = interactor.ArticleRepository.Delete(article)
	return
}
