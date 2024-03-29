package crud

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// ControllerInterface 控制器接口
type ControllerInterface interface {
	getModel() ModelInterface
	getModelSlice() interface{} // 这地方应该是模型数组，暂时无合理方式继承
	SetChildren(children ControllerInterface)
}

// Controller 控制器
type Controller struct {
	Ctx           iris.Context
	Service Service
	children      ControllerInterface
	modelInstance *ModelInstance // 此处使用指针，方便判断该属性是否为空
}

// BeginRequest 前置操作 实现框架Controller interface
func (instance *Controller) BeginRequest(ctx context.Context) {
	// 从中间件中读取共享变量
	modelInstanceInterface := ctx.Values().Get(ModelInstanceKey)
	if modelInstance, ok := modelInstanceInterface.(*ModelInstance); ok {
		instance.modelInstance = modelInstance
	}
}

// EndRequest 后置操作 实现框架Controller interface
func (instance *Controller) EndRequest(ctx context.Context) {
	//ctx.EndRequest()

}

// SetChildren 设置子类(该方法子类不能覆盖)
func (instance *Controller) SetChildren(children ControllerInterface) {
	instance.children = children
}

func (instance *Controller) getModel() ModelInterface {
	if instance.modelInstance != nil {
		return instance.modelInstance.Model
	} else if instance.children != nil {
		return instance.children.getModel()
	}
	return nil

}

func (instance *Controller) getModelSlice() interface{} {
	if instance.modelInstance != nil {
		return instance.modelInstance.ModelList
	} else if instance.children != nil {
		return instance.children.getModelSlice()
	}
	return nil
}

//PostList post 方法获取列表
func (instance *Controller) PostList(criteria Criteria) *ResponseBean {
	model := instance.getModel()
	where := criteria.GetWhere()
	pagination:= criteria.GetPagination()
	page := 1
	pageSize := 20
	orderBy := criteria.GetOrder()
	dataList := instance.getModelSlice()
	total, _ := instance.Service.Page(model, dataList, where.SQL, orderBy, pagination.Page, pagination.Size)
	response := &ResponseBean{
		Msg:  "ok",
		Code: 200,
		List: dataList,
		Pagination: &Pagination{
			Page:  page,
			Size:  pageSize,
			Total: int(total),
		},
	}
	return response
}

//PostAdd 增加
func (instance *Controller) PostAdd() *ResponseBean {
	model := instance.getModel()
	response := &ResponseBean{}
	if err := instance.Ctx.ReadJSON(model); err != nil {
		response.Msg = err.Error()
		response.Code = iris.StatusBadRequest
		return response
	}

	if _, err := instance.Service.Insert(model); err != nil {
		response.Msg = err.Error()
		response.Code = iris.StatusInternalServerError
		return response
	}

	response.Msg = "ok"
	response.Code = iris.StatusCreated

	return response
}

//PostUpdate 修改
func (instance *Controller) PostUpdate() *ResponseBean {
	model := instance.getModel()
	response := &ResponseBean{}
	if err := instance.Ctx.ReadJSON(model); err != nil {
		response.Msg = err.Error()
		response.Code = iris.StatusBadRequest
		return response
	}

	if _, err := instance.Service.Update(model); err != nil {
		response.Msg = err.Error()
		response.Code = iris.StatusInternalServerError
		return response
	}

	response.Msg = "ok"
	response.Code = iris.StatusOK

	return response
}

//GetDelete 修改
func (instance *Controller) GetDelete() *ResponseBean {
	model := instance.getModel()
	response := &ResponseBean{}
	if err := instance.Ctx.ReadJSON(model); err != nil {
		response.Msg = err.Error()
		response.Code = iris.StatusBadRequest
		return response
	}
	// 只获取id
	newModel := instance.getModel()
	newModel.SetId(model.GetId())

	if _, err := instance.Service.Delete(newModel); err != nil {
		response.Msg = err.Error()
		response.Code = iris.StatusInternalServerError
		return response
	}

	response.Msg = "ok"
	response.Code = iris.StatusOK

	return response
}