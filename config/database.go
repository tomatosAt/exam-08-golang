package config

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tomatosAt/exam-08-golang/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client struct {
	db         MariaDBConfig
	dsn        string
	ctx        *gorm.DB
	sql        *sql.DB
	gormConfig gorm.Config
	logger     *logrus.Entry
}

func NewWithConfig(cfg Config, logger *logrus.Logger) *Client {
	return &Client{
		db:     cfg.DBMain,
		logger: logger.WithField("package", "database"),
	}
}

func (c *Client) Connect() error {
	return c.ConnectWithGormConfig(gorm.Config{})
}

func (c *Client) ConnectWithGormConfig(gormCfg gorm.Config) error {
	_ = c.Close()
	c.gormConfig = gormCfg
	c.dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		c.db.User,
		c.db.Password,
		c.db.Host,
		c.db.Port,
		c.db.Database,
		"utf8mb4",
	)
	var err error
	c.ctx, err = gorm.Open(mysql.Open(c.dsn), &gormCfg)
	if err != nil {
		return err
	}

	c.sql, err = c.ctx.DB()
	if err != nil {
		return err
	}

	if c.db.Migration {
		if err := c.ctx.AutoMigrate(
			&model.Exam{},
			&model.Choice{},
		); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Ctx() *gorm.DB {
	return c.ctx
}

func (c *Client) SqlDB() *sql.DB {
	return c.sql
}

func (c *Client) MigrateDatabase(tables []interface{}) error {
	tx := c.ctx.Begin()
	for _, t := range tables {
		if err := tx.AutoMigrate(t); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (c *Client) Close() error {
	if c.sql == nil {
		return nil
	}
	if err := c.sql.Close(); err != nil {
		return err
	}
	return nil
}
