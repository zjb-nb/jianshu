package main

import (
	"sync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	doOnce sync.Once
	dsn    string = "username:psd@(ip:port)/database?database?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms&charset=utf8mb4&parseTime=true&loc=Local"
)

func GetDB() *gorm.DB {
	var err error
	doOnce.Do(func() {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		// db.Table()
		// db.Transaction()
	})
	return db
}

func TestQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	require.NoError(t, err)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(2)
	mock.ExpectQuery("SELECT *").WillReturnRows(rows)
	type Name struct {
		Id string
	}
	var name Name
	res := gdb.First(&name)
	t.Log(res.Error)
	t.Log(res.RowsAffected)
	t.Log(name.Id)
}
