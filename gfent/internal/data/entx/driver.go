package entx

import (
	gosql "database/sql"

	"entgo.io/ent/dialect/sql"
	"github.com/gogf/gf/frame/g"
)

func getDriver(master bool, name ...string) (drv *sql.Driver, err error) {
	var sqlDB *gosql.DB
	db := g.DB(name...)
	if master {
		sqlDB, err = db.Master()
	} else {
		sqlDB, err = db.Slave()
	}
	if err != nil {
		return nil, err
	}

	drv = sql.OpenDB(db.GetConfig().Type, sqlDB)
	return
}

// GetMasterDriver 主数据库驱动
func GetMasterDriver(name ...string) (drv *sql.Driver, err error) {
	return getDriver(true, name...)
}

// GetSlaveDriver 从数据库驱动
func GetSlaveDriver(name ...string) (drv *sql.Driver, err error) {
	return getDriver(false, name...)
}
