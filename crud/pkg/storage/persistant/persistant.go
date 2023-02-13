package persistant

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type db struct {
	dbtype             string
	address, database  string
	username, password string

	maxIdleCoos, maxOpenConns int

	gormDB *gorm.DB
}

type Option func(*db)

func WithDBType(t string) Option {
	return func(d *db) {
		d.dbtype = t
	}
}

func WithAddress(a string) Option {
	return func(d *db) {
		d.address = a
	}
}

func DatabaseName(dn string) Option {
	return func(d *db) {
		d.database = dn
	}
}

func WithUsernamePassword(u, p string) Option {
	return func(d *db) {
		d.username = u
		d.password = p
	}
}

func New(opts ...Option) (*db, error) {
	db := &db{
		// Hardcoing it now, but can be configured via the option func.
		maxIdleCoos:  10,
		maxOpenConns: 50,
	}

	for _, opt := range opts {
		opt(db)
	}

	var err error
	switch db.dbtype {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True", db.username, db.password, db.address, db.database)
		db.gormDB, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			return nil, err
		}
	case "sqlite":
		db.gormDB, err = gorm.Open(sqlite.Open(db.database))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported database type requested")
	}

	sqlDB, err := db.gormDB.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(db.maxIdleCoos)
		sqlDB.SetMaxOpenConns(db.maxOpenConns)
	}
	return db, nil
}
