package entrest

import "entgo.io/ent/schema"

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
	Prefix string
}

// Name describes the annotation name.
func (API) Name() string {
	return "EntRESTAPI"
}

var (
	_ schema.Annotation = (*Field)(nil)
	_ schema.Annotation = (*API)(nil)
)
