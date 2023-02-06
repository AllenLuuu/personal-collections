package handler

import (
	"personal-collections/model"
	"personal-collections/resp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InsertCollection(c *gin.Context) {
	var req struct {
		Content string `json:"content"`
		Author  string `json:"author"`
		Book	string `json:"book"`
		Tags	[]string `json:"tags"`
	}
	var collection model.Collection
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	collection.Content = req.Content
	collection.Author = req.Author
	collection.Book = req.Book
	collection.Tags = req.Tags
	err = collection.Create()
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_INSERT_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, collection)
}

func GetCollectionByID(c *gin.Context) {
	var req struct {
		ID string `json:"id"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	var collection model.Collection
	err = collection.GetByID(req.ID)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_SEARCH_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, collection)
}

func GetCollectionsByIDs(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	collections, err := model.GetCollectionsByIDs(req.IDs)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_SEARCH_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, collections)
}

func ListStarredCollections(c *gin.Context) {
	var req model.Filter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	collections, err := model.ListStarredCollections(req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_SEARCH_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, collections)
}

func ListCollections(c *gin.Context) {
	var req model.Filter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	collections, err := model.ListCollections(req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_SEARCH_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, collections)
}

func UpdateCollection(c *gin.Context) {
	var req model.Collection
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	err = req.Update()
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_UPDATE_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, interface{}(nil))
}

func DeleteCollection(c *gin.Context) {
	var req struct {
		ID string `json:"id"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	err = model.DeleteCollection(req.ID)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_DELETE_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, interface{}(nil))
}