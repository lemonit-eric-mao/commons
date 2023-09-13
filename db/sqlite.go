package db

import (
    "database/sql"
    "github.com/lemonit-eric-mao/commons/logger"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func init() {
    var err error
    database, err = sql.Open("sqlite3", "database.db")
    if err != nil {
        log.Fatal(err)
    }
}

func Close() error {
    return database.Close()
}

// Exec 执行增、删、改
func Exec(query string, args ...any) (sql.Result, error) {

    // 语句无参数时，直接执行
    if len(args) == 0 {
        return database.Exec(query)
    }

    // 语句有参数时，执行
    // 准备语句的过程包括语法分析、编译以及优化
    stmt, err := database.Prepare(query)
    if err != nil {
        logger.Error(err)
    }
    return stmt.Exec(args...)
}

// Query 查询多条
func Query(query string, args ...any) (*sql.Rows, error) {

    // 准备语句的过程包括语法分析、编译以及优化
    stmt, err := database.Prepare(query)
    if err != nil {
        logger.Error(err)
    }
    return stmt.Query(args...)
}

// QueryRow 查询单条
func QueryRow(query string, args ...any) *sql.Row {
    stmt, err := database.Prepare(query)
    if err != nil {
        logger.Error(err)
    }
    return stmt.QueryRow(args...)
}
