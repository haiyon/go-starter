package conf

import (
	"time"
)

// ORMConfig orm configs object
type ORMConfig struct {
	Debug           bool
	Dialect         string
	DSN             string
	AutoMigrate     bool
	TablePrefix     string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime time.Duration
}

func getORMConfig() *ORMConfig {
	return &ORMConfig{
		Debug:           c.GetBool("orm.debug"),
		Dialect:         c.GetString("orm.dialect"),
		DSN:             c.GetString("orm.dsn"),
		AutoMigrate:     c.GetBool("orm.auto_migrate"),
		TablePrefix:     c.GetString("orm.table_prefix"),
		MaxIdleConn:     c.GetInt("orm.max_idle_conn"),
		MaxOpenConn:     c.GetInt("orm.max_open_conn"),
		ConnMaxLifeTime: c.GetDuration("orm.max_life_time"),
	}
}
