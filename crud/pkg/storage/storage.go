package storage

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/golang/glog"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type StorageType string

const (
	MySQL  StorageType = "mysql"
	SqLite StorageType = "sqlite"
	Memory StorageType = "memory"
)

// StorageConfig struct include the database type, address, username, password, and database name.
//
// The database type field specifies the type of database to connect to, which can be one of the following
// three options: "mysql", "sqlite", or "in-memory". Depending on the database type.
// The username and password fields are used to authenticate the connection to the database, and the database name
// field specifies the name of the database to connect to.
// The StorageConfig struct is typically used as a parameter to the Init function, which initializes a database
// connection based on the provided configuration parameters.
type StorageConfig struct {
	// The type of database to connect to (mysql, sqlite, in-memory)
	DatabaseType StorageType
	// The hostname:port of the database server.
	Address string
	// The username to authenticate the database connection.
	Username string
	// The password to authenticate the database connection
	Password string
	// The name of the database to connect to, file path for sqlite.
	Database string

	MaxIdleConns, MaxOpenConns int
}

var (
	initialized uint32
)

// Init initializes a database connection based on the provided StorageConfig.
// The supported database types are MySQL, SQLite, and in-memory. The connection is established
// using the configuration parameters provided in StorageConfig, which include the database type,
// host, port, username, password, and database name.
//
// If the database type is MySQL or SQLite, Init creates a new connection to the specified database
// using the provided credentials. If the database type is in-memory, it creates an in-memory database
// that can be used for testing or temporary storage purposes.
//
// If the connection is successful, Init stores a Global instance of the database instance, which can be used
// to perform database operations, by calling storage.Pool(). If an error occurs during the connection
// process, Init returns an error, which should be handled by the calling function.
func Init(cfg StorageConfig) error {
	if atomic.LoadUint32(&initialized) != 0 {
		return fmt.Errorf("storage.Init called more than once")
	}

	atomic.StoreUint32(&initialized, 1)

	glog.Infof("Database type %q requested", cfg.DatabaseType)
	switch cfg.DatabaseType {
	case MySQL:
		glog.Infof("Creating mysql database connection for address %v", cfg.Address)
		dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True", cfg.Username, cfg.Password, cfg.Address, cfg.Database)
		db, err := gorm.Open(
			mysql.Open(dsn),
			&gorm.Config{
				Logger: newGormLogger(),
			},
		)
		if err != nil {
			glog.Errorf("Failed to create mysql connection, reason %v", err)
			return err
		}
		// Set Global instance for db
		global = NewCrudStorageForDB(db, cfg.MaxIdleConns, cfg.MaxOpenConns)

		glog.Infof("mysql connection ready.")
	case SqLite:
		glog.Infof("Creating sqlite database connection for database %v", cfg.Address)
		db, err := gorm.Open(
			sqlite.Open(cfg.Database),
			&gorm.Config{
				Logger: newGormLogger(),
			},
		)
		if err != nil {
			return err
		}
		// Set Global instance for db
		global = NewCrudStorageForDB(db, cfg.MaxIdleConns, cfg.MaxOpenConns)

		glog.Infof("sqlite connection ready.")
	case Memory:
		glog.Infof("Creating in-memory database")
		// Set Global instance for memory
		global = NewCrudStorageForMemory()

		glog.Infof("in-memory database ready.")
	default:
		return fmt.Errorf("unsupported database type")
	}
	return nil
}

func Pool() Store {
	if atomic.LoadUint32(&initialized) != 1 {
		glog.Fatalf("storage.Pool called before storage.Init, exiting application")
	}
	return global
}

func Shutdown() {
	if atomic.LoadUint32(&initialized) != 1 {
		glog.Errorf("storage.Shutdown called before storage.Init, skipping shutdown")
		return
	}
	global.Shutdown()
}

func newGormLogger() logger.Interface {
	return logger.New(&appLogWriter{}, logger.Config{
		SlowThreshold:        200 * time.Millisecond,
		LogLevel:             logger.Info,
		ParameterizedQueries: true,
	})
}

type appLogWriter struct{}

func (*appLogWriter) Printf(s string, v ...interface{}) {
	s = strings.Replace(s, "\n", " ", -1)
	sourceFile := v[0].(string)
	v[0] = sourceFile[strings.LastIndex(sourceFile, "/")+1:]
	glog.InfoDepth(1, fmt.Sprintf(s, v...))
}
