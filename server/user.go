package server

import (
	"context"
	"gin-test-gorm/model"
	"gin-test-gorm/structure"
	"gin-test-gorm/tools"
	"gin-test-gorm/tools/metadata"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var (
		err  error
		db   = model.GetDb()
		ctx  = context.Background()
		user = model.User{}
		data = structure.Id{}
	)
	defer func() {
		if err != nil {
			db.Rollback()
		}
		db.Commit()
	}()

	db = db.Begin()
	defer func() {
		tools.BuildResponse(c, err, data)
	}()
	err = db.Debug().WithContext(ctx).Where("id = ?", 1).First(&user).Error
	if err != nil {
		return
	}
	data.Token = "token-123456"
	metadata.UserInfo.Store(user.ID, "token-123456")
	metadata.UserInfo.Store("token-123456", user.ID)
}

func Create(c *gin.Context) {
	var (
		err  error
		db   = model.GetDb()
		ctx  = context.TODO()
		p    = &model.UserParams{}
		user = model.User{}
		data = structure.Id{}
	)
	err = tools.ShouldBind(c, p)
	if err != nil {
		return
	}

	db = db.Begin()

	defer func() {
		if err != nil {
			db.Rollback()
		}
		db.Commit()
	}()

	defer func() {
		tools.BuildResponse(c, err, data)
	}()
	{
		user.Name = p.Name
		user.Password = p.Password
		user.Phone = p.Phone
	}
	value, _ := metadata.UserInfo.Load("token-123456")
	id := value.(uint)
	context.WithValue(ctx, "user_id", id)
	context.WithValue(ctx, "creator_id", id)
	context.WithValue(ctx, "updater_id", id)
	context.WithValue(ctx, "deleted_id", id)
	//err = db.Debug().WithContext(ctx).Create(&user).Error
	model.CreateModel(ctx, db, &user)
}

func AuthAdmin(c *gin.Context) {
	token := c.GetHeader("Authorization")
	value, ok := metadata.UserInfo.Load(token)
	if !ok {
		BuildUnauthorized(c)
	}
	c.Set("user", value)
}

func BuildUnauthorized(c *gin.Context) {
	if _, ok := getResult(c); ok {
		return
	}
	defer setResult(c, true)
	// _ = log.WarnLog(errors.NewError(errors.ErrCodeBusinessAuthorize))
	c.AbortWithStatus(http.StatusUnauthorized)
}
func getResult(c *gin.Context) (interface{}, bool) {
	v, ok := c.Get("result")
	if ok {
		return v, true
	}
	return nil, false
}
func setResult(c *gin.Context, resp interface{}) {
	c.Set("result", resp)
}
