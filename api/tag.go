package api

import (
	"github.com/Demonx24/Vblog-backend/model/database"
	"github.com/Demonx24/Vblog-backend/model/request"
	"github.com/Demonx24/Vblog-backend/model/response"
	"github.com/gin-gonic/gin"
)

type TagApi struct {
}

func (TagApi *TagApi) GetTagId(c *gin.Context) {
	var req request.TagId
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tag := database.Tag{TagId: req.TagId}
	tag, err := tagService.TagId(tag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(tag, c)
}
func (TagApi *TagApi) AddTag(c *gin.Context) {
	var req request.AddTag
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	tag := database.Tag{TagName: req.TagName, TagId: req.TagId}
	tag, err := tagService.AddTag(tag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithData(tag, c)
}
func (TagApi *TagApi) DeleteTag(c *gin.Context) {
	var req request.TagId
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	tag := database.Tag{TagId: req.TagId}
	if err := tagService.DeleteTag(tag); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithMessage("删除成功", c)
}
func (TagApi *TagApi) EditTag(c *gin.Context) {
	var req request.AddTag
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	tag := database.Tag{TagId: req.TagId, TagName: req.TagName}
	tag, err := tagService.UpdateTag(tag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithData(tag, c)
}
func (TagApi *TagApi) GetAllTags(c *gin.Context) {
	tags, err := tagService.GetAllTags()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithData(tags, c)
}
