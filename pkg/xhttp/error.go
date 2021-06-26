package xhttp

import (
	"net/http"
	"haiyon/go-starter/pkg/ecode"
	"haiyon/go-starter/pkg/types"
)

// AlreadyExists 已存在
func AlreadyExists(message string, data ...types.JSON) *ResponseException {
	return &ResponseException{
		Status:  http.StatusOK,
		Code:    2003,
		Message: message,
		Data:    data,
	}

}

// NotExists 不存在
func NotExists(message string, data ...types.JSON) *ResponseException {
	return &ResponseException{
		Status:  http.StatusOK,
		Code:    2002,
		Message: message,
		Data:    data,
	}

}

// DBQuery 数据库查询错误
func DBQuery(message string, data ...types.JSON) *ResponseException {
	return &ResponseException{
		Status:  http.StatusInternalServerError,
		Code:    ecode.ServerErr,
		Message: message,
		Data:    data,
	}

}

// Transactions 交易错误
func Transactions(message string, data ...types.JSON) *ResponseException {
	return &ResponseException{
		Status:  http.StatusInternalServerError,
		Code:    ecode.ServerErr,
		Message: message,
		Data:    data,
	}

}

// UnAuthorized 未认证
func UnAuthorized(message string, data ...types.JSON) *ResponseException {
	return &ResponseException{
		Status:  http.StatusUnauthorized,
		Code:    ecode.Unauthorized,
		Message: message,
		Data:    data,
	}

}

// BadRequest 请求错误
func BadRequest(message string, data ...types.JSON) *ResponseException {
	return &ResponseException{
		Status:  http.StatusBadRequest,
		Code:    ecode.RequestErr,
		Message: message,
		Data:    data,
	}

}

// NotFound 找不到资源
func NotFound(message string, data ...types.JSON) *ResponseException {
	return &ResponseException{
		Status:  http.StatusNotFound,
		Code:    ecode.NothingFound,
		Message: message,
		Data:    data,
	}

}

// Forbidden 拒绝访问
func Forbidden(message string, data ...types.JSON) *ResponseException {
	return &ResponseException{
		Status:  http.StatusForbidden,
		Code:    ecode.AccessDenied,
		Message: message,
		Data:    data,
	}

}

// InternalServer 服务器错误
func InternalServer(message string, data ...types.JSON) *ResponseException {
	return &ResponseException{
		Status:  http.StatusInternalServerError,
		Code:    ecode.ServerErr,
		Message: message,
		Data:    data,
	}

}
