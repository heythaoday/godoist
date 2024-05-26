package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"godoist/internal/config"
	"godoist/model/mysql"
)

type ServiceContext struct {
	Config config.Config
	Model  mysql.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  mysql.NewUserModel(sqlx.NewMysql(c.DataSource)),
	}
}
