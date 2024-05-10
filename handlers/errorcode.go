package handlers

import "errors"

var (
	ErrNotFound        = errors.New("resource not found")    // 资源未找到
	ErrInternalServer  = errors.New("internal server error") // 服务器内部错误
	ErrBadRequest      = errors.New("bad request")           // 错误的请求
	ErrUnauthorized    = errors.New("unauthorized")          // 未授权
	ErrForbidden       = errors.New("forbidden")             // 禁止访问
	ErrConflict        = errors.New("conflict")              // 冲突错误，如尝试创建已存在的资源
	ErrTooManyRequests = errors.New("too many requests")     // 请求过多
)
