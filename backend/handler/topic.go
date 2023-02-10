package handler

import (
	"personal-collections/model"
	"personal-collections/resp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InsertTopic(c *gin.Context) {
	var req struct {
		Title       string   `json:"title"`
		Detail      string   `json:"detail"`
		Collections []string `json:"collections"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	var topic model.Topic
	topic.Title = req.Title
	topic.Detail = req.Detail
	topic.Collctions = req.Collections
	err = topic.Create()
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_INSERT_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, topic)
}

func GetTopicByID(c *gin.Context) {
	var req struct {
		ID string `json:"id"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	var topic model.Topic
	err = topic.GetByID(req.ID)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_SEARCH_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, topic)
}

func ListTopics(c *gin.Context) {
	topics, err := model.ListTopics()
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_SEARCH_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, topics)
}

func UpdateTopic(c *gin.Context) {
	var req model.Topic
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

func DeleteTopic(c *gin.Context) {
	var req struct {
		ID string `json:"id"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	err = model.DeleteTopic(req.ID)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_DELETE_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, interface{}(nil))
}

func GetTopicCollections(c *gin.Context) {
	var req struct {
		Tid    string       `json:"tid"`
		Filter model.Filter `json:"filter"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	var topic model.Topic
	err = topic.GetByID(req.Tid)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_SEARCH_ERROR, "数据库错误")
		return
	}

	collections, err := topic.GetCollections(req.Filter)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_SEARCH_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, collections)
}

func AddTopicCollection(c *gin.Context) {
	var req struct {
		Tid string `json:"tid"`
		Cid string `json:"cid"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	var topic model.Topic
	err = topic.GetByID(req.Tid)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_SEARCH_ERROR, "数据库错误")
		return
	}

	err = topic.AddCollection(req.Cid)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_UPDATE_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, interface{}(nil))
}

func RemoveTopicCollection(c *gin.Context) {
	var req struct {
		Tid string `json:"tid"`
		Cid string `json:"cid"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	var topic model.Topic
	err = topic.GetByID(req.Tid)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_SEARCH_ERROR, "数据库错误")
		return
	}

	err = topic.RemoveCollection(req.Cid)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_DB_UPDATE_ERROR, "数据库错误")
		return
	}
	resp.JSON(c, interface{}(nil))
}
