package main

import "fmt"

// 表操作
func table() {
	// 创建表
	err := engine.CreateTables(xUser{})
	if err != nil {
		panic(err)
	}

	// 判断表是否为空
	exist, err := engine.IsTableExist(xUser{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("isTableExist=%t\n", exist)

	// 判断表是否为空
	empty, err := engine.IsTableEmpty(xUser{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("isTableEmpty=%t\n", empty)

	// 删除表
	//err = engine.DropTables(xUser{})
	//if err != nil {
	//	 panic(err)
	//}
}

// 索引操作
func index() {
	// 创建唯一索引
	//err := engine.CreateUniques(xUser{})
	//if err != nil {
	//	panic(err)
	//}

	// 创建索引
	//err = engine.CreateIndexes(xUser{})
	//if err != nil {
	//	panic(err)
	//}
}

// 同步数据库结构
func sync() {
	// 自动检测和创建表，这个检测是根据表的名字
	// 自动检测和新增表中的字段，这个检测是根据字段名
	// 自动检测和创建索引和唯一索引，这个检测是根据索引的一个或多个字段名，而不根据索引名称
	err := engine.Sync(xUser{})
	if err != nil {
		panic(err)
	}

	// 自动检测和创建表，这个检测是根据表的名字
	// 自动检测和新增表中的字段，这个检测是根据字段名，同时对表中多余的字段给出警告信息
	// 自动检测，创建和删除索引和唯一索引，这个检测是根据索引的一个或多个字段名，而不根据索引名称。因此这里需要注意，如果在一个有大量数据的表中引入新的索引，数据库可能需要一定的时间来建立索引。
	// 自动转换 varchar 字段类型到text字段类型，自动警告其它字段类型在模型和数据库之间不一致的情况。
	// 自动警告字段的默认值，是否为空信息在模型和数据库之间不匹配的情况
	// 对 Sync 进行了改进，目前推荐使用 sync2
	err = engine.Sync2(xUser{})
	if err != nil {
		panic(err)
	}
}

// 导入导出 SQL 脚本
func script() {
	//engine.DumpAll()
	//engine.DumpAllToFile()
}
