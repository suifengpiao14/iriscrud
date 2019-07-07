package crud

import (
	"github.com/kataras/iris"
	"fmt"
	"crypto/md5"
)

// ModelInstance crud 中传递模型和模型列表类型
type ModelInstance struct {
	Model     ModelInterface // 单个模型
	ModelList interface{}    // 模型列表
}

//ModelInstanceKey 中间件共享模型数据键名称
const ModelInstanceKey = "modelInstance"

// ModelInstanceMiddle 注册modelInstance实例中间件
func ModelInstanceMiddle(NewModelInstance func() *ModelInstance) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		modelInstance := NewModelInstance()
		ctx.Values().Set(ModelInstanceKey, modelInstance)
		ctx.Next()
	}
}

// NewModelTemplateInstance 返回模型数据实例(crud使用，每个model生成一个，作为ModelInstanceMiddle的函数参数，目前没有找到好的方式，产生对象，所以每个model生成一个方法)
func NewModelTemplateInstance() *ModelInstance {
	modelList := make([]ModelTemplate, 0)
	modelInstance := &ModelInstance{
		Model:     &ModelTemplate{},
		ModelList: &modelList,
	}
	return modelInstance

}

// ModelInterface 模型接口
type ModelInterface interface {
	TableName() string
	GetId() int
	SetId(id int)
}

// ModelTemplate model 模板
type ModelTemplate struct {
	Id int
}

// NewModelTemplate  创建ModelTemplate
func (instance *ModelTemplate) NewModelTemplate() ModelInterface {
	return &ModelTemplate{}
}

//TableName   模型对应的表名称
func (instance *ModelTemplate) TableName() string {
	return ""
}

// GetId 获取记录主键
func (instance *ModelTemplate) GetId() int {
	return instance.Id
}

// SetId 获取记录主键
func (instance *ModelTemplate) SetId(id int) {
	instance.Id = id
}

// Md5 将字符串md5化
func Md5(str string) string {

	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
