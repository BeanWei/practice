package entrest

import "strings"

// api errors
const (
	ErrorParameterBind int = 1 << iota
	ErrorDuplicate
	ErrorCheckDuplicate
	ErrorNotFound
	ErrorList
	ErrorGet
	ErrorCreate
	ErrorUpdate
	ErrorDelete
)

// api request types
const (
	AllRequest    = ""
	ListRequest   = "list"
	GetRequest    = "get"
	CreateRequest = "create"
	UpdateRequest = "update"
	DeleteRequest = "delete"
)

// Result .
type Result struct {
	Data      interface{}
	IsList    bool
	Total     int
	ErrorType int
	Error     error
}

// MiddlewareConfig 中间件定义
type MiddlewareConfig struct {
	PkgPath string // 中间件包路径
	ReqType string // 接口请求类型
	Code    string // 中间件代码
}

const (
	ASC  = "asc"
	DESC = "desc"
)

// GetSortArg .
func GetSortArg(sort string) (sortName, sortOp string) {
	if strings.HasPrefix(sort, "-") {
		sortName = strings.TrimLeft(sort, "-")
		sortOp = DESC
	} else if strings.HasPrefix(sort, "+") {
		sortName = strings.TrimLeft(sort, "+")
		sortOp = ASC
	} else {
		sortName = sort
		sortOp = ASC
	}
	return
}
