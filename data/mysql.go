package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"time"
)

const dsnFormat = "%v:%v@tcp(127.0.0.1:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"

type MysqlConf struct {
	username string
	passwd   string
	port     string
	dbname   string
}

func (*MysqlConf) generatedsn(c *MysqlConf) (dsn string) {
	return fmt.Sprintf(dsnFormat, c.username, c.passwd, c.port, c.dbname)
}

var sqlConfCold = MysqlConf{username: "root", passwd: "123456", port: "3306", dbname: "test_gorm"}
var sqlConfHot = MysqlConf{username: "root", passwd: "123456", port: "13306", dbname: "test_gorm"}

var dsnCold = fmt.Sprintf(dsnFormat, sqlConfCold.username, sqlConfCold.passwd, sqlConfCold.port, sqlConfCold.dbname)
var dsnHot = fmt.Sprintf(dsnFormat, sqlConfHot.username, sqlConfHot.passwd, sqlConfHot.port, sqlConfHot.dbname)

type groupMemberTable struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;not null"`
	UID        int64     `gorm:"not null;"`
	GID        int64     `gorm:"not null;"`
	CREATETIME time.Time `gorm:"autoCreateTime;not null"`
}

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
