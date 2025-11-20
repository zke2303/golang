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
	// 1. 构建基础查询 (只包含 Where 条件)
	// 这个 tx 是干净的，包含所有的过滤条件
	tx := buildDynamicWhere(repo.db, query)

	// 2. 处理分页参数默认值 (防止 Size=0 导致查全表)
	if page.Page <= 0 {
		page.Page = 1
	}
	if page.PageSize <= 0 {
		page.PageSize = 10 // 给一个默认值
	}

	// 3. 计算总数 (Count)
	var total int64
	// 【关键修复】：
	// 使用 tx.Model(...) 而不是 tx.Find(...)。
	// 注意：这里直接链式调用，不要把结果赋值回 tx，或者依赖返回值，这样不会污染原始 tx
	if err := tx.Model(&model.Book{}).Count(&total).Error; err != nil {
		return response.PageResult{}, repo.db // 返回错误
	}

	// 4. 查询数据列表 (List)
	var books []model.Book

	// 【关键修复】：
	// 基于原始的 tx 构建一个新的查询对象 listTx
	// 这样 buildPagination 里的 Offset/Limit 也是加在一个干净的基础上
	listTx := buildPagination(tx, page)

	if err := listTx.Find(&books).Error; err != nil {
		return response.PageResult{}, repo.db
	}

	// 5. 组装返回结果
	var pageResult response.PageResult
	pageResult.Record = books
	pageResult.Current = page.Page
	pageResult.Size = page.PageSize
	pageResult.Total = total

	return pageResult, nil
}

// buildPagination 不需要改动，但在调用时要注意不要覆盖原始 tx
func buildPagination(tx *gorm.DB, page *request.Page) *gorm.DB {
	offset := (page.Page - 1) * page.PageSize
	// 这里的 Session(&gorm.Session{}) 是可选的高级技巧，用于确保完全隔离，
	// 但通常只要上面的 PageQuery 逻辑写对，这里直接返回 tx.Offset... 也是可以的。
	return tx.Offset(offset).Limit(page.PageSize)
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
