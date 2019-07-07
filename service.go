package crud

import (
	"github.com/go-xorm/xorm"
)

// Service 通用服务接口
type Service interface {
	Update(model ModelInterface) (int64, error)
	Insert(model ModelInterface) (int64, error)
	Delete(model ModelInterface) (int64, error)
	Get(model ModelInterface) (bool, error)
	GetAll(rowsSlicePtr interface{},queryModel ModelInterface) (int64, error)
	Exists(model ModelInterface) (bool, error)
	Save(queryModel ModelInterface, model ModelInterface) (int64, error)
	Engine() *xorm.Engine
	Page(model ModelInterface, slicePtr interface{}, where string, orderBy string, page int, pageSize int) (count int64, err error)
}

type service struct {
	dao *Dao
}

//NewService 创建历史名称服务实例
func NewService(engine *xorm.Engine) Service {
	return &Service{
		dao: NewDao(engine),
	}
}

// Update 更新
func (instance *service) Update(model ModelInterface) (int64, error) {
	return instance.dao.Update(model, []string{})
}

// Insert 新增
func (instance *service) Insert(model ModelInterface) (int64, error) {
	return instance.dao.Insert(model)
}

// Delete 新增
func (instance *service) Delete(model ModelInterface) (int64, error) {
	return instance.dao.Delete(model)
}

// Get 获取
func (instance *service) Get(model ModelInterface) (bool, error) {
	return instance.dao.Get(model)
}
// GetAll 获取所有
func (instance *service)GetAll(rowsSlicePtr interface{},queryModel ModelInterface) (int64, error) {
	return instance.dao.GetAll(rowsSlicePtr,queryModel)
}

// 分页查询
func (instance *service)Page(model ModelInterface, slicePtr interface{}, where string, orderBy string, page int, pageSize int) (count int64, err error) {
	return instance.dao.Page(model, slicePtr, where, orderBy, page, pageSize)
}

// Exists 是否存在
func (instance *service) Exists(model ModelInterface) (bool, error) {
	return instance.dao.Exists(model)
}

// Engin 获取所有记录
func (instance *service) Engine() *xorm.Engine {
	return instance.dao.Engine()
}

// Save 保存
func (instance *service) Save(queryModel ModelInterface, model ModelInterface) (int64, error) {
	ok, err := instance.Get(queryModel)
	if err != nil {
		return 0, err
	}
	if ok {
		model.SetId(queryModel.GetId()) // 根据id更新
		return instance.Update(model)
	}

	return instance.Insert(model)
}