package db

import (
	"database/sql"
	"errors"

	"github.com/kinwyb/go/err"
)

//一个简便的数据库操作封装
var ErrorNotOpen = errors.New("数据库连接失败")

//TransactionFunc 事务回调函数
type TransactionFunc func(tx TxSQL) err.Error

//错误解析
type FormatError interface {
	FormatError(e error) err.Error
}

//数据库操作接口
type SQL interface {
	//查询一条数据,每条结果以map[string]interface{}方式返回
	//返回结果,使用本package中的类型函数进行数据解析
	//eg:
	//		result := QueryRow(...)
	//		result.Error(func(error.Error){
	//			这里处理错误
	// 		}).Rows(func(map[string]interface{}) bool {
	//			return true //返回true，继续循环读取下一条
	// 		})
	//param sql string SQL
	//param args... interface{} SQL参数
	QueryRow(sql string, args ...interface{}) QueryResult
	//查询多条数据,每条结果以map[string]interface{}方式返回
	//返回结果,使用本package中的类型函数进行数据解析
	//eg:
	//		result := QueryRow(...)
	//		result.Error(func(error.Error){
	//			这里处理错误
	// 		}).Rows(func(map[string]interface{}) bool {
	//			return true //返回true，继续循环读取下一条
	// 		})
	//param sql string SQL
	//param args... interface{} SQL参数
	QueryRows(sql string, args ...interface{}) QueryResult
	//Exec 执行一条SQL
	//param sql string SQL
	//param args... interface{} SQL参数
	Exec(sql string, args ...interface{}) ExecResult
	//Count SQL语句条数统计
	//param sql string SQL
	//param args... interface{} SQL参数
	Count(sql string, args ...interface{}) (int64, err.Error)
	//ParseSQL 解析SQL
	//param sql string SQL
	//param args map[string]interface{} 参数映射
	ParseSQL(sql string, args map[string]interface{}) (string, []interface{}, err.Error)
	//RowsPage 分页查询SQL
	//返回结果,使用本package中的类型函数进行数据解析
	//eg:
	//		result := QueryRow(...)
	//		result.Error(func(error.Error){
	//			这里处理错误
	// 		}).Rows(func(map[string]interface{}) bool {
	//			return true //返回true，继续循环读取下一条
	// 		})
	//param sql string SQL
	//param page *PageObj 分页数据
	//param args... interface{} SQL参数
	QueryWithPage(sql string, page *PageObj, args ...interface{}) QueryResult
	//Transaction 事务处理
	//param t TransactionFunc 事务处理函数
	Transaction(t TransactionFunc) err.Error
	//GetDb 获取数据库对象
	GetDb() (*sql.DB, err.Error)
	//Close 关闭数据库
	Close()
}

//事物数据操作接口
type TxSQL interface {
	//Rows 查询多条数据,结果以[]map[string]interface{}方式返回
	//返回结果,使用本package中的类型函数进行数据解析
	//eg:
	//		result := QueryRow(...)
	//		result.Error(func(error.Error){
	//			这里处理错误
	// 		}).Rows(func(map[string]interface{}) bool {
	//			return true //返回true，继续循环读取下一条
	// 		})
	//param sql string SQL
	//param args... interface{} SQL参数
	QueryRows(sql string, args ...interface{}) QueryResult
	//Rows 查询多条数据,结果以[]map[string]interface{}方式返回
	//返回结果,使用本package中的类型函数进行数据解析
	//eg:
	//		result := QueryRow(...)
	//		result.Error(func(error.Error){
	//			这里处理错误
	// 		}).Rows(func(map[string]interface{}) bool {
	//			return true //返回true，继续循环读取下一条
	// 		})
	//param sql string SQL
	//param args... interface{} SQL参数
	QueryRow(sql string, args ...interface{}) QueryResult
	//Exec 执行一条SQL
	//param sql string SQL
	//param args... interface{} SQL参数
	Exec(sql string, args ...interface{}) ExecResult
	//Count SQL语句条数统计
	//param sql string SQL
	//param args... interface{} SQL参数
	Count(sql string, args ...interface{}) (int64, err.Error)
	//RowsPage 分页查询SQL
	//返回结果,使用本package中的类型函数进行数据解析
	//eg:
	//		result := QueryRow(...)
	//		result.Error(func(error.Error){
	//			这里处理错误
	// 		}).Rows(func(map[string]interface{}) bool {
	//			return true //返回true，继续循环读取下一条
	// 		})
	//param sql string SQL
	//param page *PageObj 分页数据
	//param args... interface{} SQL参数
	QueryWithPage(sql string, page *PageObj, args ...interface{}) QueryResult
	//GetTx 获取事务对象
	GetTx() *sql.Tx
}
