package crud

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

//TableGenerator 初始化数据表
type TableGenerator struct {
	models []interface{}
	engine *xorm.Engine
}

//NewTableGenerator 实例化TableGenerator
fuc NewTableGenerator(engine *xorm.Engine,models []interface{}) (instance *TableGenerator){
	instance.Models = models
	instance.engine = engine
	return

}

// SetModel 记录model 实例数组（需要建表的model实例列表）
func (instance *TableGenerator) SetModel(models []interface{}) {
	instance.Models = models
}

// SetEngine 设置mysql链接实例
func (instance *TableGenerator) SetEngine(engine *xorm.Engine) {
	instance.engine = engine
}

// Run 根据model生成表
func (instance *TableGenerator) Run() {
	engine:=instance.engine
	for _, table := range instance.models {
		if exist, _ := engine.IsTableExist(table); exist {
			_ = engine.DropTables(table)
		}
		err := engine.CreateTables(table)
		if err != nil {
			fmt.Print(err)
		}
	}
}
