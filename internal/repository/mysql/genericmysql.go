package mysql

import (
	"com.wisecharge/central/configs"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type GenericMysql struct {
	delegate *gorm.DB
}

func New(log *zap.Logger) (*GenericMysql, error) {
	cfg := configs.Get().MySQL

	conf := &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}
	if log.Level() == zap.DebugLevel {
		conf.Logger = logger.Default.LogMode(logger.Info)
	}

	db := &GenericMysql{}
	if delegate, err := gorm.Open(mysql.Open(cfg.DSN), conf); err != nil {
		return nil, err
	} else if pool, err := delegate.DB(); err != nil {
		return nil, err
	} else {
		pool.SetMaxIdleConns(cfg.MaxIdleConnection)
		pool.SetMaxOpenConns(cfg.MaxOpenConnection)
		pool.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Minute)

		db.delegate = delegate
		return db, nil
	}
}

func (d *GenericMysql) Shutdown() error {
	delegate, err := d.delegate.DB()
	if err != nil {
		return err
	}

	return delegate.Close()
}

func (d *GenericMysql) Delegate() *gorm.DB {
	return d.delegate
}
