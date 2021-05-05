package entrest

import (
	"entgo.io/ent/schema"
)

const (
	EntRESTAPI   = "EntRESTAPI"
	EntRESTField = "EntRESTField"
	EntRESTEdge  = "EntRESTEdge"
)

// API 接口配置注解
type API struct {
	// 忽略此 schema 接口自动生成
	SkipGen bool
	// 接口 uri 前缀
	Prefix string
	// 中间件
	Middlewares []MiddlewareConfig
}

// Name describes the API annotation name.
func (API) Name() string {
	return EntRESTAPI
}

// Field 接口字段注解
type Field struct {
	// 表单验证规则: 遵循 gf 的验证规则
	Validate string
	// 可新建
	Creatable bool
	// 可更新
	Updatable bool
	// 可排序
	Sortable bool
	// 可查询/筛选
	Queryable bool
	// 查询操作符
	QueryOp string
}

// Name describes the Field annotation name.
func (Field) Name() string {
	return EntRESTField
}

// Edge 关联字段注解
type Edge struct {
	// 表单验证规则: 遵循 gf 的验证规则
	Validate string
	// 详情查询时进行关联查询
	GetWith bool
	// 列表查询时进行关联查询
	QueryWith bool
}

// Name describes the Edge annotation name.
func (Edge) Name() string {
	return EntRESTEdge
}

var (
	_ schema.Annotation = (*Field)(nil)
	_ schema.Annotation = (*API)(nil)
	_ schema.Annotation = (*Edge)(nil)
)
