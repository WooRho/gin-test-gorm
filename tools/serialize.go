package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-test-gorm/structure"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func ShouldBind(c *gin.Context, p structure.IParam) (err error) {
	err = c.ShouldBind(p)
	if err != nil {
		switch vErr := err.(type) {
		case *json.UnmarshalTypeError:
			msg := fmt.Sprintf(`参数"%v"格式错误`, vErr.Field)
			err = errors.New(msg)
		default:
			err = err
		}
	}
	if err != nil {
		return
	}
	return
}

// 用于返回错误或详情数据
func BuildResponse(c *gin.Context, err error, data structure.IResponseData) {
	if err != nil {
		BuildCustomError(c, err)
	} else {
		BuildDataContext(c, data)
	}
}

// 用于返回错误
func BuildCustomError(c *gin.Context, err error) {
	if _, ok := getResult(c); ok {
		return
	}
	var (
		response structure.Response
	)
	response.Msg = err.Error()
	response.Code = -1
	//response.SetVersion(version)
	//switch vErr := err.(type) {
	//case *errors.Error:
	//	response.Code = vErr.GetCode()
	//	response.Msg = vErr.GetMsg()
	//default:
	//	response.Code = -1
	//	response.Msg = err.Error()
	//}
	fmt.Println("error = ", response.Msg)
	c.JSON(http.StatusInternalServerError, response)
	c.Abort()
	setResult(c, response)
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

// 用于返回详情数据
func BuildDataContext(c *gin.Context, data structure.IResponseData) {
	if _, ok := getResult(c); ok {
		return
	}

	if data != nil {
		//data.AdjustData()
	}
	response := structure.Response{}
	//response.SetVersion(version)
	response.Code = 0
	response.Msg = "success"
	response.Data = data
	c.JSON(http.StatusOK, response)
	c.Abort()
	setResult(c, response)
}

// 用于返回列表数据（不支持数据报表导出）  （新的API建议使用buildListResponseV2）
func BuildListResponse(c *gin.Context, err error, list structure.IResponseData, total int64) {
	if err != nil {
		BuildCustomError(c, err)
	} else {
		if _, ok := getResult(c); ok {
			return
		}

		response := structure.ListResponse{}
		//response.SetVersion(version)
		response.Code = 0
		response.Msg = "success"
		// 判断list值是否为空
		if !reflect.ValueOf(list).IsNil() {
			//list.AdjustData()
			response.Data.List = list
			response.Data.Total = total
		}
		c.JSON(http.StatusOK, response)
		setResult(c, response)
	}
}
