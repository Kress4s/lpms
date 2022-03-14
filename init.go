package main

import (
	"lpms/commom/drivers/database"
	"lpms/migrations"
)

/*

检查初始化插件工作连接是否正常使用
1. 数据库驱动
2. 数据库连接
3. redis连接
...
*/

func init() {
	if database.GetDriver() == nil {
		panic("connect database error")
	}

	// if minio.GetDriver() == nil {
	// 	panic("connect minio error")
	// }

	if err := migrations.Migrate(); err != nil {
		panic(err)
	}
}
