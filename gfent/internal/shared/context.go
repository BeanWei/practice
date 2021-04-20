package shared

import (
	"context"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Context 上下文管理
var Context = ctxShared{}

// CtxKey 上下文变量存储键名
const CtxKey = "CtxKey"

// Ctx 请求上下文结构
type Ctx struct {
	Session *ghttp.Session // 当前Session管理对象
	User    *CtxUser       // 上下文用户信息
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

// CtxUser 请求上下文中的用户信息
type CtxUser struct {
	ID       string // 用户 id
	Username string // 用户名
	Phone    string // 手机号
	Email    string // 邮箱
	Avatar   string // 头像
	Disabled bool   // 状态
	IsAdmin  bool   // 是否为管理员
}

type ctxShared struct{}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *ctxShared) Init(r *ghttp.Request, customCtx *Ctx) {
	r.SetCtxVar(CtxKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *ctxShared) Get(ctx context.Context) *Ctx {
	value := ctx.Value(CtxKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*Ctx); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *ctxShared) SetUser(ctx context.Context, ctxUser *CtxUser) {
	s.Get(ctx).User = ctxUser
}

// SetData 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *ctxShared) SetData(ctx context.Context, data g.Map) {
	s.Get(ctx).Data = data
}
