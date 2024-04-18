package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zhangjiacheng-iHealth/IHCommunity/package/dao/impl"
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/middleware"
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/model"
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/utils"
	"github.com/gin-gonic/gin"
)

type IHCommunityControllerImpl struct {
	dao *impl.IHCommunityDaoImpl
}

func NewIHCommunityControllerImpl() *IHCommunityControllerImpl {
	return &IHCommunityControllerImpl{dao: impl.NewIHCommunityDaoImpl()}
}


func (impl IHCommunityControllerImpl) AuthAdmin(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	claims, error := middleware.ParseToken(auth)
	type authResponse struct {
		LoginName string `json:"loginName"`
		Password  string `json:"password"`
	}
	res := authResponse{LoginName: claims.LoginName, Password: claims.Password}

	if error != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"code": -1, "msg": error.Error(), "count": 0, "data": nil})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
	}
}

func (impl IHCommunityControllerImpl) Auth(c *gin.Context) {
	type SearchUserRequest struct {
		UserId string `json:"userId"`
	}

	var req SearchUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	user := impl.dao.SearchUser(c, req.UserId)

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "count": 0, "data": user})
}

func (impl IHCommunityControllerImpl) AddUser(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}
	
	user := model.User{}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	res, err := impl.dao.AddUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl IHCommunityControllerImpl) DeleteComment(c *gin.Context) {
	type DeleteCommentRequest struct {
		Id int64 `json:"id"`
	}

	var req DeleteCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	res := impl.dao.DeleteComment(c, req.Id)

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl IHCommunityControllerImpl) UpdateComment(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	comment := model.Comment{}
	err = json.Unmarshal(bytes, &comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	comment = impl.dao.UpdateComment(c, comment)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "count": 0, "data": comment})
}

func (impl IHCommunityControllerImpl) DeletePost(c *gin.Context) {
	type DeletePostRequest struct {
		Id int64 `json:"id"`
	}

	var req DeletePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	res := impl.dao.DeletePost(c, req.Id)

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "count": 0, "data": res})
}


func (impl IHCommunityControllerImpl) AddComment(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	comment := model.Comment{}
	err = json.Unmarshal(bytes, &comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	res, err := impl.dao.AddComment(c, comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl IHCommunityControllerImpl) AddPost(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	post := model.Post{}
	err = json.Unmarshal(bytes, &post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	res, err := impl.dao.AddPost(c, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl IHCommunityControllerImpl) UpdatePost(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": "读取请求体失败"})
			return
	}
	
	post := model.Post{}
	err = json.Unmarshal(bytes, &post)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": "解析请求体失败"})
			return
	}

	updatedPost, err := impl.dao.UpdatePost(post)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": "更新帖子失败"})
			return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "更新帖子成功", "data": updatedPost})
}


func (impl IHCommunityControllerImpl) GetAllPosts(c *gin.Context) {
	type GetAllPostsRequest struct {
		Id int64 `json:"id"`
	}

	var req GetAllPostsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
		return
	}

	res := impl.dao.GetAllPosts(c, req.Id)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl IHCommunityControllerImpl) GetCommentsForUser(c *gin.Context) {
    type GetCommentListRequest struct {
        UserID int64 `json:"userId"`
        PostID int64 `json:"postId"`
    }

    var req GetCommentListRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
        return
    }

    comments, err := impl.dao.GetCommentsForUser(c, req.UserID, req.PostID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": "Failed to get comment list", "count": 0, "data": nil})
        return
    }

    c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "count": len(comments), "data": comments})
}

func (impl IHCommunityControllerImpl) GetProposalInfo(c *gin.Context) {
	type GetProposalInfoRequest struct {
			Id int64 `json:"id"`
	}

	var req GetProposalInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error(), "count": 0, "data": nil})
			return
	}

	proposal := impl.dao.GetProposalInfo(c, req.Id)
	plans := impl.dao.GetPlansByProposalID(c, req.Id)
	privileges := impl.dao.GetPrivilegesByProposalID(c, req.Id)

	// 组装数据并返回
	data := gin.H{
			"proposal":     proposal,
			"plans":        plans,
			"privileges":   privileges,
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "count": 0, "data": data})
}

func (impl IHCommunityControllerImpl) UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": "获取上传文件失败", "data": nil})
		return
	}
	defer file.Close()

	filename := header.Filename
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": "上传文件名不能为空", "data": nil})
		return
	}

	err = impl.dao.UploadImage(c, file, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": "上传图片失败", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "上传图片成功", "data": nil})
}

func (impl IHCommunityControllerImpl) UploadVideo(c *gin.Context) {
	file, header, err := c.Request.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": "获取上传文件失败", "data": nil})
		return
	}
	defer file.Close()

	filename := header.Filename
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": "上传文件名不能为空", "data": nil})
		return
	}

	err = impl.dao.UploadVideo(c, file, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": "上传视频失败", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "上传视频成功", "data": nil})
}
