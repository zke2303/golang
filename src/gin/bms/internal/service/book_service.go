package service

import (
	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/repository"
	"github.com/zhang/bms/internal/transport/http/request"
	"github.com/zhang/bms/internal/transport/http/response"
	"gorm.io/gorm"
)

type IBookService interface {
	Create(book *model.Book) *gorm.DB
	Delete(id uint64) *gorm.DB
	FindById(id uint64) (*model.Book, *gorm.DB)
	Update(r *model.Book) *gorm.DB
	PageQuery(page *request.Page, query *request.BookQuery) (response.PageResult, *gorm.DB)
}

type BookServiceImpl struct {
	BookRepository repository.IBookRepository
}

func NewBookService(repository repository.IBookRepository) IBookService {
	return &BookServiceImpl{BookRepository: repository}
}

func (s *BookServiceImpl) Create(b *model.Book) *gorm.DB {
	return s.BookRepository.Create(b)
}

func (s *BookServiceImpl) Delete(id uint64) *gorm.DB {
	return s.BookRepository.Delete(id)
}
func (s *BookServiceImpl) FindById(id uint64) (*model.Book, *gorm.DB) {
	return s.BookRepository.FindById(id)
}

func (s *BookServiceImpl) Update(b *model.Book) *gorm.DB {
	return s.BookRepository.Update(b)
}

func (s *BookServiceImpl) PageQuery(page *request.Page, query *request.BookQuery) (response.PageResult, *gorm.DB) {
	return s.BookRepository.PageQuery(page, query)
}
