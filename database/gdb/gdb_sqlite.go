// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
// Note:
// 1. It needs manually import: _ "github.com/mattn/go-sqlite3"
// 2. It does not support Save/Replace features.

package gdb

import (
	"database/sql"
)

type dbSqlite struct {
	*dbBase
}

func (db *dbSqlite) Open(config *ConfigNode) (*sql.DB, error) {
	var source string
	if config.LinkInfo != "" {
		source = config.LinkInfo
	} else {
		source = config.Name
	}
	if db, err := sql.Open("sqlite3", source); err == nil {
		return db, nil
	} else {
		return nil, err
	}
}

func (db *dbSqlite) getChars() (charLeft string, charRight string) {
	return "`", "`"
}

// 返回当前数据库所有的数据表名称
// TODO
func (db *dbSqlite) Tables() (tables []string, err error) {
	return
}

// 获得指定表表的数据结构，构造成map哈希表返回，其中键名为表字段名称，键值为字段数据结构.
// TODO
func (db *dbSqlite) TableFields(table string) (fields map[string]*TableField, err error) {
	return
}

// 在执行sql之前对sql进行进一步处理。
// @todo 需要增加对Save方法的支持，可使用正则来实现替换，
// @todo 将ON DUPLICATE KEY UPDATE触发器修改为两条SQL语句(INSERT OR IGNORE & UPDATE)
func (db *dbSqlite) handleSqlBeforeExec(query string) string {
	return query
}
