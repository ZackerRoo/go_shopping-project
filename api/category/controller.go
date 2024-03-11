package category

import (
	"fmt"
	"net/http"
	"pro05shopping/domain/category"
	"pro05shopping/utils/api_helper"
	pagination "pro05shopping/utils/pageination"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	categoryService *category.Service
}

// 实例化控制器
func NewCategoryController(service *category.Service) *Controller {
	return &Controller{
		categoryService: service,
	}
}

// CreateCategory godoc
// @Summary 根据给定的参数创建分类
// @Tags Category
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param CreateCategoryRequest body CreateCategoryRequest true "category information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /category [post]
func (c *Controller) CreateCategory(g *gin.Context) {
	var req CreateCategoryRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	newCategory := category.NewCategory(req.Name, req.Desc)
	err := c.categoryService.Create(newCategory)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, api_helper.Response{Message: "Category created"})
}

// BulkCreateCategory godoc
// @Summary 根据给定的csv文件，批量创建分类
// @Tags Category
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param   file formData file true  "file contains category information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /category/upload [post]
func (c *Controller) BulkCreateCategory(g *gin.Context) {
	fileHeader, err := g.FormFile("file")
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	count, err := c.categoryService.BulkCreate(fileHeader)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(
		http.StatusOK, api_helper.Response{
			Message: fmt.Sprintf(
				"'%s' uploaded! '%d' new categories created", fileHeader.Filename, count)})
}

// GetCategories godoc
// @Summary 获得分类列表
// @Tags Category
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} pagination.Pages
// @Router /category [get]
func (c *Controller) GetCategories(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	page = c.categoryService.GetAll(page)
	g.JSON(http.StatusOK, page)

}
