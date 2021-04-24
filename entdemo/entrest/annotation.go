package entrest

import (
	"entgo.io/ent/schema"
)

const (
	EntRESTAPI   = "EntRESTAPI"
	EntRESTField = "EntRESTField"
)

// Field 接口字段注解
type Field struct {
	// 表单验证规则: 遵循 gf 的验证规则
	Validate string
	// 可新建
	Creatable bool
	// 可更新
	Updatable bool
	// 可检索
	Searchable bool
}

// Name describes the annotation name.
func (Field) Name() string {
	return "EntRESTField"
}

// API 接口配置注解
type API struct {
	// 忽略此 schema 接口自动生成
	SkipGen bool
	// 接口 uri 前缀
	Prefix string
	// 中间件
	Middlewares []MiddlewareConfig
}

// Name describes the annotation name.
func (API) Name() string {
	return "EntRESTAPI"
}

var (
	_ schema.Annotation = (*Field)(nil)
	_ schema.Annotation = (*API)(nil)
)
