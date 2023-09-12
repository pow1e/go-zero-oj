package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	username = "root"
	password = "123456"
	database = "go-zero-oj"
	dns      = username + ":" + password + "@tcp(127.0.0.1:3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	gormdb, _ := gorm.Open(mysql.Open(dns))

	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		// 相对于go.mod目录下的
		OutPath: "app/cmd/api/internal/dal/query", // curd代码的输出路径

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		//FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		//FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		//FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		//FieldWithTypeTag: true, // generate with gorm column type tag
	})

	g.UseDB(gormdb) // reuse your gorm db

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	//dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
	//	"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	//	"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	//	"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	//	"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	//	"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	//}
	// 要先于`ApplyBasic`执行
	//g.WithDataTypeMap(dataMap)

	//softDeleteField := gen.FieldType("deleted_time", "soft_delete.DeletedAt")
	// 模型自定义选项组
	//fieldOpts := []gen.ModelOpt{softDeleteField}

	//models := g.GenerateAllTable(fieldOpts...)

	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
