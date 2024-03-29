package crud
import(
	"github.com/kataras/iris"
)

// Criteria where 参数，将参数转为查询条件语句
type Criteria interface {
	GetWhere()*Where
	SetWhre(ctx iris.Context,fn func(ctx iris.Context)(sql string,args *map[string]interface{}))
	GetOrder()string
	SetOrder(ctx iris.Context,fn func(ctx iris.Context)(order string))
	GetPagination()*Pagination
	SetPagination(ctx iris.Context,fn func(ctx iris.Context)(page int,size int))
}



// criteria where 参数，将参数转为查询条件语句
type criteria struct {
	Where      *Where
	Order      string
	Pagination *Pagination
}

//NewCriteria 创建查条件器构造器
func NewCriteria() Criteria {
	return &criteria{
		Where:&Where{},
		Pagination: &Pagination{},
	}
}

//GetWhere 获取where对象
func(instance *criteria)GetWhere()*Where{
	return instance.Where
}

//SetWhre 设置查询条件
func (instance *criteria)SetWhre(ctx iris.Context,fn func(ctx iris.Context) (sql string,args *map[string]interface{})){
	instance.Where.SQL,instance.Where.Args = fn(ctx)
	return 
}

//GetOrder 获取order 字符串
func(instance *criteria)GetOrder()string{
	return instance.Order
}

//SetOrder 设置排序条件
func (instance *criteria)SetOrder(ctx iris.Context,fn func(ctx iris.Context) (order string)){
	instance.Order = fn(ctx)
	return 
}

//GetPagination 获取pagination对象
func(instance *criteria)GetPagination()*Pagination{
	return instance.Pagination
}

//SetPagination 设置分页条件
func (instance *criteria)SetPagination(ctx iris.Context,fn func(ctx iris.Context) (page int,size int)){
	instance.Pagination.Page,instance.Pagination.Size = fn(ctx)
	return 
}
