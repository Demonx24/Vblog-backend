package api

import (
	"github.com/Demonx24/Vblog-backend/model/database"
	"github.com/Demonx24/Vblog-backend/model/request"
	"github.com/Demonx24/Vblog-backend/model/response"
	"github.com/gin-gonic/gin"
)

type BlogApi struct{}

func (BlogApi *BlogApi) BloggerID(c *gin.Context) {

	var req request.BloggerId
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage("传参数错误", c)
		return
	}
	b := database.Blog{Id: req.BlogId}
	blog, err := blogService.BloggerId(b)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(blog, c)
}
func (BlogApi *BlogApi) CreateBlog(c *gin.Context) {
	var req request.CreateBlog
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("传参数错误", c)
		return
	}
	b := database.Blog{Title: req.Title, Content: req.Content, Class: req.Class}
	blog, err := blogService.CreateBlog(b)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(blog, c)
}
func (BlogApi *BlogApi) DeleteBlog(c *gin.Context) {
	var req request.DeleteBlog
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("传参数错误", c)
		return
	}
	b := database.Blog{Id: req.BlogId}
	err = blogService.DeleteBlog(b)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
func (BlogApi *BlogApi) UpdateBlog(c *gin.Context) {
	var req request.UpdateBlog
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	b := database.Blog{Id: req.BlogId, Title: req.Title, Content: req.Content, Class: req.Class}
	err = blogService.UpdateBlog(b)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
func (BlogApi *BlogApi) BlogByPage(c *gin.Context) {
	var pageInfo request.BlogList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage("<参数错误>", c)
	}
	list, total, err := blogService.BlogList(pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}
func (BlogApi *BlogApi) GetCard1(c *gin.Context) {
	var card1 database.Card1
	card1, err := blogService.GetCard1(card1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(card1, c)
}
func (BlogApi *BlogApi) GetCard4(c *gin.Context) {

	blogs, err := blogService.GetCard4()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(blogs, c)
}
func (BlogApi *BlogApi) GetArchive(c *gin.Context) {

	blogs, err := blogService.GetArchive()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(blogs, c)
}
func (BlogApi *BlogApi) GetIndex(c *gin.Context) {
	blogs, err := blogService.GetIndex()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(blogs, c)
}
