package smc

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/BeanWei/gfent/internal/data/entx"
	"github.com/BeanWei/gfent/internal/data/smc/ent"
	_ "github.com/BeanWei/gfent/internal/data/smc/ent/runtime"
	"github.com/gogf/gf/frame/g"
)

// Client is alias
type Client *ent.Client

var (
	masterClient Client
	slaveClient  Client
)

func newClient(drv *sql.Driver) Client {
	client := ent.NewClient(ent.Driver(drv))
	if g.DB().GetConfig().Debug {
		client.Debug()
	}

	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			v, err := next.Mutate(ctx, m)
			if err != nil {
				g.DB().GetLogger().Error(err)
				return nil, err
			}
			return v, nil
		})
	})

	if err := client.Schema.Create(context.Background()); err != nil {
		g.Log().Fatalf("[Module SMC] - failed creating schema resources: %v", err)
	}

	return client
}

// 初始化主从 ent client
func init() {
	masterDrv, er := entx.GetMasterDriver()
	if er != nil {
		g.Log().Fatalf("init smc-module master db failed: %v", er)
	}
	masterClient = newClient(masterDrv)

	slaveDrv, er := entx.GetSlaveDriver()
	if er != nil {
		g.Log().Errorf("init smc-module slave db failed: %v", er)
		slaveClient = masterClient
	} else {
		slaveClient = newClient(slaveDrv)
	}
}

// NewClient .
func NewClient(master bool) Client {
	if master {
		return masterClient
	}
	return slaveClient
}
