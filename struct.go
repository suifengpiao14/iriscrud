package crud

//ResponseBean 返回体
type ResponseBean struct {
	Msg        string                 `json:"msg"`
	Code       int                    `json:"code"`
	Record     interface{}            `json:"record"`
	List       interface{}            `json:"list"`
	Params     map[string]interface{} `json:"params"`
	Pagination *Pagination            `json:"pagination"`
}

//Pagination 分页器
type Pagination struct {
	Size  int `json:"size"`
	Page  int `json:"page"`
	Total int `json:"total"`
}

//Where 条件对象
type Where struct{
	SQL string
	Args *map[string]interface{}
}