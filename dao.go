package crud

import (
	"errors"
	"github.com/go-xorm/xorm"
)

// Dao 通用dao
type Dao struct {
	engine *xorm.Engine
}

//NewDao 生成通用dao 实例
func NewDao(engine *xorm.Engine) *Dao {
	return &Dao{engine: engine}
}

// Update 更新记录
func (instance *Dao) Update(data ModelInterface, columns []string) (int64, error) {
	return instance.engine.Id(data.GetId()).MustCols(columns...).Update(data)
}

// Insert 新增记录
func (instance *Dao) Insert(data ModelInterface) (int64, error) {
	return instance.engine.Insert(data)
}

// Delete 删除记录
func (instance *Dao) Delete(data ModelInterface) (int64, error) {
	id := data.GetId()
	if id == 0 {
		return 0, errors.New("id must")
	}
	return instance.engine.Id(id).Delete(data)
}

// Get 根据条件(唯一条件) 获取单条记录
func (instance *Dao) Get(model ModelInterface) (bool, error) {
	return instance.engine.Get(model)
}

// GetAll 根据条件 获取所有记录
func (instance *Dao) GetAll(rowsSlicePtr interface{}, queryModel ModelInterface) (int64, error) {
	instance.engine.Find(rowsSlicePtr, queryModel)
	return 1,nil
	//return instance.engine.FindAndCount(dataList, queryModel)
}

//Exists 判断记录是否存在
func (instance *Dao) Exists(model ModelInterface) (bool, error) {
	return instance.engine.Exist(model)

}

//Page 分页查询
func (instance *Dao) Page(model ModelInterface, slicePtr interface{}, where string, orderBy string, page int, pageSize int) (count int64, err error) {
	// 查询数量
	count, err = instance.engine.Where(where).Count(model)
	// 出错或者总数为空直接返回
	if err != nil||count == 0  {
		return 
	}

	offset := (page - 1) * pageSize
	err = instance.engine.Where(where).OrderBy(orderBy).Limit(pageSize, offset).Find(slicePtr)
	return
}

//Engine 获取xrom 操作类
func (instance *Dao) Engine() *xorm.Engine {
	return instance.engine
}
