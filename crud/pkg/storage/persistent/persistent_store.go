package persistent

import (
	"fmt"

	"github.com/golang/glog"
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
		glog.Infof("Creating mysql database connection for address %v", db.address)
		dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True", db.username, db.password, db.address, db.database)
		db.gormDB, err = gorm.Open(
			mysql.Open(dsn),
			&gorm.Config{
				Logger: newAppLogger(),
			},
		)
		if err != nil {
			return nil, err
		}
		glog.Infof("mysql connection ready.")
	case "sqlite":
		glog.Infof("Creating sqlite database connection for database %v", db.address)
		db.gormDB, err = gorm.Open(
			sqlite.Open(db.database),
			&gorm.Config{
				Logger: newAppLogger(),
			},
		)
		if err != nil {
			return nil, err
		}
		glog.Infof("sqlite connection ready.")
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
