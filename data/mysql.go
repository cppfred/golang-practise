package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"time"
)

// Auto generate mysql dsn

const dsnFormat = "%v:%v@tcp(127.0.0.1:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"

type MysqlConf struct {
	username string
	passwd   string
	port     string
	dbname   string
}

func (c *MysqlConf) generatedsn() (dsn string) {
	return fmt.Sprintf(dsnFormat, c.username, c.passwd, c.port, c.dbname)
}

func NewMysqlConf(username string, passwd string, port string, dbname string) *MysqlConf {
	return &MysqlConf{username: username, passwd: passwd, port: port, dbname: dbname}
}

var dsnCold = NewMysqlConf("root", "123456", "3306", "test_cold").generatedsn()
var dsnHot = NewMysqlConf("root", "123456", "3306", "test_hot").generatedsn()

// Table structure define

type groupMemberTable struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;not null"`
	UID        int64     `gorm:"not null;"`
	GID        int64     `gorm:"not null;"`
	CREATETIME time.Time `gorm:"autoCreateTime:milli;not null"`
}

type p2bmessageTable struct { // for both hot and cold db
	ID         uint64    `gorm:"primaryKey;autoIncrement;not null"`
	From       int64     `gorm:"not null;"`
	To         int64     `gorm:"not null;"`
	SeqID      int64     `gorm:""`
	Type       int32     `gorm:"not null;"`
	Context    string    `gorm:"not null;"`
	CREATETIME time.Time `gorm:"autoCreateTime:milli;not null"`
}

type p2pmessageTable struct { // for both hot and cold db
	ID         uint64    `gorm:"primaryKey;autoIncrement;not null"`
	From       int64     `gorm:"not null;"`
	To         int64     `gorm:"not null;"`
	Type       int32     `gorm:"not null;"`
	Context    string    `gorm:"not null;"`
	CREATETIME time.Time `gorm:"autoCreateTime:milli;not null"`
}

// SqlConn struct for MySQL Connection and its operate function
type SqlConn struct {
	db *gorm.DB
}

func NewMysqlConn(dsn string) (sql *SqlConn, err error) {
	newDbConn, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil,
			errors.Wrapf(err, "failed to open mysql connection with dsn setting (%v)", dsn)
	}
	return &SqlConn{
		db: newDbConn,
	}, nil
}

func (sqldb *SqlConn) CreateTable(args ...interface{}) {
	sqldb.db = sqldb.db.AutoMigrate(args)
}

func (sqldb *SqlConn) ListTable() {
	_ = sqldb.db
}

func (sqldb *SqlConn) SwitchTable() {
	sqldb.db = sqldb.db.Table("")
}
