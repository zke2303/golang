package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/service"
	"github.com/zhang/bms/internal/transport/http/request"
	"github.com/zhang/bms/internal/transport/http/response"
)

type IBookHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	FindById(c *gin.Context)
	Update(c *gin.Context)
	PageQuery(c *gin.Context)
}
type BookHandlerImpl struct {
	bookService service.IBookService
}

func NewBookHandler(service service.IBookService) IBookHandler {
	return &BookHandlerImpl{bookService: service}
}

// Create 插入一条数据
func (h BookHandlerImpl) Create(c *gin.Context) {
	// 1.从请求中获取数据, 并转换成 model.Book 对象
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		fmt.Println("格式转换错误\n", err)
		response.FailWithMsg(c, http.StatusBadRequest, "数据格式错误")
		return
	}

	// 2.调用 service 层

	result := h.bookService.Create(&book)
	if result.Error != nil {
		fmt.Println("数据操作失败\n", result.Error)
		response.FailWithMsg(c, http.StatusInternalServerError, "数据操作失败")
		return
	}

	// 3.成功
	response.Success(c, nil)
}

func (h BookHandlerImpl) Delete(c *gin.Context) {
	// 1.从请求中获取参数Id
	strId := c.Param("id")
	// 2.转换成 uint64 类型
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		fmt.Println("输入的不是数字")
		response.FailWithMsg(c, http.StatusBadRequest, "请输入数字")
		return
	}

	// 3.调用 service 层
	result := h.bookService.Delete(id)
	if result.RowsAffected <= 0 {
		fmt.Println("无法删除未存在的数据", result.Error)
		response.FailWithMsg(c, http.StatusInternalServerError, "无法删除未存在的数据")
		return
	}

	if result.Error != nil {
		fmt.Println("数据库删除失败", result.Error)
		response.FailWithMsg(c, http.StatusInternalServerError, "数据库删除数据失败")
		return
	}
	// 4.返回成功信息
	response.Success(c, nil)
}

func (h BookHandlerImpl) FindById(c *gin.Context) {
	// 1.从请求中获取参数Id
	strId := c.Query("id")
	// 2.转换成 uint64 类型
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		fmt.Println("输入的不是数字")
		response.FailWithMsg(c, http.StatusBadRequest, "请输入数字")
		return
	}

	// 3.调用 service 层
	book, result := h.bookService.FindById(id)
	if result.RowsAffected <= 0 {
		fmt.Println("该数据不存在")
		response.FailWithMsg(c, http.StatusInternalServerError, "该数据不存在")
		return
	}

	if result.Error != nil {
		fmt.Println("数据库查询失败")
		response.FailWithMsg(c, http.StatusInternalServerError, "数据库查询失败")
		return
	}

	// 4.返回成功信息
	response.Success(c, book)
}

func (h BookHandlerImpl) Update(c *gin.Context) {
	var book model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		fmt.Println("格式转换错误\n", err)
		response.FailWithMsg(c, http.StatusBadRequest, "数据格式错误")
		return
	}

	// 2.调用 service 层
	result := h.bookService.Update(&book)
	if result.Error != nil {
		fmt.Println("数据操作失败\n", err)
		response.FailWithMsg(c, http.StatusInternalServerError, "数据操作失败")
		return
	}

	// 3.判断是否修改成功
	if result.RowsAffected <= 0 {
		fmt.Println("未找到相关记录, 无法进行修改")
		response.FailWithMsg(c, http.StatusBadRequest, "未找到相关记录, 无法进行修改")
		return
	}

	// 3.成功
	response.Success(c, nil)

}

func (h BookHandlerImpl) PageQuery(c *gin.Context) {
	var page request.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, "分页参数错误")
		return
	}

	fmt.Println(page)

	var query request.BookQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, "条件查询参数错误")
		return
	}

	fmt.Println(query)

	pageResult, dbResult := h.bookService.PageQuery(&page, &query)
	if dbResult != nil && dbResult.Error != nil {
		response.FailWithMsg(c, http.StatusInternalServerError, dbResult.Error.Error())
		return
	}

	response.Success(c, pageResult)
}
