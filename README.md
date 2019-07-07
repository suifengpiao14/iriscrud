# 基于github.com/kataras/iris 框架和github.com/go-xorm/xorm 组件实现的通用crud包
该包是基于github.com/kataras/iris 框架和github.com/go-xorm/xorm 组件实现通用crud操作api接口
文件说明：
dao.go service 和 mysql 隔离层
service.go controller 层获取数据的接口
model_interface.go  系统的model必须实现ModelInterface 接口方法 (可用使用template_goxorm_struct.go.tpl模板自动根据数据表生成)
table_generator.go 根据model 生成数据库表，方便迁移
controller.go 通用基本crud 控制器
template_goxorm_struct.go.tpl 根据表生成实现ModelInterface 接口的model类，并生成获取模型实例、列表的方法
# 简单的使用案例
```
import (
	"../../bootstrap"
	"../../models"
	"../../services"
	"../controllers"
	"github.com/kataras/iris/mvc"
)
// 后台v1版本接口
	backendV1 := index.Party("/backend/v1/")

	// xml
	backendV1.Party("/xml", models.ModelInstanceMiddle(models.NewXmlModelInstance)).Handle(new(controllers.BaseController))
	// tag
	backendV1.Party("/tag", models.ModelInstanceMiddle(models.NewTagModelInstance)).Handle(new(controllers.BaseController))
```