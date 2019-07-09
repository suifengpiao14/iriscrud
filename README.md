# 基于github.com/kataras/iris 框架和github.com/go-xorm/xorm 组件实现的通用crud包
该包是基于github.com/kataras/iris 框架和github.com/go-xorm/xorm 组件实现通用crud操作api接口
文件说明：
dao.go service 和 mysql 隔离层
service.go controller 层获取数据的接口
model_interface.go  系统的model必须实现ModelInterface 接口方法 (可用使用template_goxorm_struct.go.tpl模板自动根据数据表生成)
table_generator.go 根据model 生成数据库表，方便迁移
controller.go 通用基本crud 控制器
template_goxorm_struct.go.tpl 根据表生成实现ModelInterface 接口的model类，并生成获取模型实例、列表的方法
criteria.go 构造sql查询条件,单独提炼成对象后方便结构化，进而考虑在openapi文档中标识后自动化
# 简单的使用案例
```
package routes

import (
	"../../bootstrap"
    "../../models"
	crud "suifengpiao14/iriscrud"
	"github.com/kataras/iris/mvc"
)

// Configure 配置路由和依赖di
func Configure(bootstrap *bootstrap.Bootstrapper) {
	service := crud.NewService()

	index := mvc.New(bootstrap.Party("/"))

	index.Register(
		service,
	)

	// 后台v1版本接口
	backendV1 := index.Party("/backend/v1/")
	// xml
	backendV1.Party("/xml", crud.ModelInstanceMiddle(models.NewXmlModelInstance)).Handle(new(crud.Controller))
	// tag
	backendV1.Party("/tag", crud.ModelInstanceMiddle(models.NewTagModelInstance)).Handle(new(crud.Controller))
	// serverVariable
	backendV1.Party("/server-variable", crud.ModelInstanceMiddle(models.NewServerVariableModelInstance)).Handle(new(crud.Controller))
	// server
	backendV1.Party("/server", crud.ModelInstanceMiddle(models.NewServerModelInstance)).Handle(new(crud.Controller))
}
```