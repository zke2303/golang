package repository

import (
	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/transport/http/request"
	"github.com/zhang/bms/internal/transport/http/response"
	"gorm.io/gorm"
)

type IBookRepository interface {
	Create(b *model.Book) *gorm.DB
	Delete(id uint64) *gorm.DB
	FindById(id uint64) (*model.Book, *gorm.DB)
	Update(b *model.Book) *gorm.DB
	PageQuery(page *request.Page, query *request.BookQuery) (response.PageResult, *gorm.DB)
}

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &BookRepositoryImpl{db: db}
}

func (repo *BookRepositoryImpl) Create(b *model.Book) *gorm.DB {
	return repo.db.Create(b)
}

func (repo *BookRepositoryImpl) Delete(id uint64) *gorm.DB {
	return repo.db.Delete(&model.Book{}, id)
}

func (repo *BookRepositoryImpl) FindById(id uint64) (*model.Book, *gorm.DB) {
	var book model.Book
	result := repo.db.First(&book, id)
	return &book, result
}

func (repo *BookRepositoryImpl) Update(b *model.Book) *gorm.DB {
	return repo.db.Updates(b)
}

func (repo *BookRepositoryImpl) PageQuery(page *request.Page, query *request.BookQuery) (response.PageResult, *gorm.DB) {
	tx := buildDynamicWhere(repo.db, query)
}

func buildDynamicWhere(tx *gorm.DB, query *request.BookQuery) *gorm.DB {
	// 字符串匹配 (Title)
	if query.Title != "" {
		// 使用 LIKE 进行模糊匹配，或使用 = 进行精确匹配，这里用 LIKE
		tx = tx.Where("title LIKE ?", "%"+query.Title+"%")
	}

	// 精确匹配 (AuthorId, Category, Status, PublisherId) - 排除 0 值
	if query.AuthorId != 0 {
		tx = tx.Where("author_id = ?", query.AuthorId)
	}
	if query.Category != 0 {
		tx = tx.Where("category = ?", query.Category)
	}
	if query.Status != 0 {
		tx = tx.Where("status = ?", query.Status)
	}
	if query.PublisherId != 0 {
		tx = tx.Where("publisher_id = ?", query.PublisherId)
	}
	if query.WordCount != 0 {
		tx = tx.Where("word_count = ?", query.WordCount)
	}

	// ISBN (通常是精确匹配)
	if query.ISBN != "" {
		tx = tx.Where("ISBN = ?", query.ISBN)
	}

	// 价格 (例如，查询大于等于这个价格的)
	if query.Price != 0 {
		tx = tx.Where("price >= ?", query.Price) // 示例：查询最小价格
	}

	// 时间字段 (PublishDate)
	// 假设查询的是在该日期之后出版的
	if !query.PublishDate.IsZero() {
		tx = tx.Where("publish_date >= ?", query.PublishDate)
	}

	return tx
}
