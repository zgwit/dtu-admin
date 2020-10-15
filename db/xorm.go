package db

import (
	"git.zgwit.com/zgwit/iot-admin/conf"
	"git.zgwit.com/zgwit/iot-admin/models"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func Open() error {

	if Engine != nil {
		return nil
	}

	cfg := conf.Config.Database
	var err error
	Engine, err = xorm.NewEngine(cfg.Type, cfg.Url)
	if err != nil {
		return err
	}
	Engine.ShowSQL(cfg.ShowSQL)

	//同步表
	return Engine.Sync2(
		models.Tunnel{}, models.Link{}, models.Device{}, models.Plugin{},
		models.Model{}, models.ModelAdapter{}, models.ModelVariable{}, models.ModelBatch{}, models.ModelJob{}, models.ModelStrategy{},
		models.Element{}, models.ElementVariable{},
	)
}
