package entrest

import "entgo.io/ent/schema"

// Annotation is a builtin schema annotation for attaching
// RESTful metadata to schema objects for both codegen and runtime.
type Annotation struct {
	// 表单验证规则: 遵循 gf 的验证规则
	Validate string
	// 可新建
	Creatable bool
	// 可更新
	Updatable bool
	// 可读取
	Readable bool
}

// Name describes the annotation name.
func (Annotation) Name() string {
	return "EntREST"
}

var (
	_ schema.Annotation = (*Annotation)(nil)
)
