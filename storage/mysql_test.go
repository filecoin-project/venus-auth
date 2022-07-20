package storage

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

var mySQLStore *mysqlStore
var sqlDB *sql.DB
var mock sqlmock.Sqlmock

func init() {
	mysqlSetup()
}

func TestMysqlStore(t *testing.T) {
	t.Run("add users mysql", testMySQLAddUser)
	// TODO: More tests
	mysqlShutdown()
}

func testMySQLAddUser(t *testing.T) {
	now := time.Now()
	id := uuid.NewString()
	user := "test_user_001"

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `users` (`id`,`name`,`comment`,`stype`,`state`,`createTime`,`updateTime`,`is_deleted`) VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs(id, user, "", 0, 0, now, now, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := mySQLStore.PutUser(&User{
		Id:         id,
		Name:       user,
		UpdateTime: now,
		CreateTime: now,
	})
	if err != nil {
		t.Fatalf("add user failed:%s", err.Error())
	}
}

func mysqlSetup() {
	var err error
	sqlDB, mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	mySQLStore = &mysqlStore{db: gormDB}
}

func mysqlShutdown() {
	sqlDB.Close()
}
