package {{.Models}}
{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
	{{range .Imports}}"{{.}}"{{end}}
)
{{end}}

{{range .Tables}}
{{$modelName :=Mapper .Name}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
{{end}}
}

//TableName   模型对应的表名称
func (instance *{{$modelName}} ) TableName() string {
	return "{{.Name}}"
}

// GetId 获取记录主键
func (instance *{{$modelName}}) GetId() int {
	return instance.Id
}

// SetId 获取记录主键
func (instance *{{$modelName}}) SetId(id int)  {
	instance.Id = id
}



// New{{$modelName}}ModelInstance 返回模型数据实例(crud使用)
func New{{$modelName}}ModelInstance() *ModelInstance {
	modelList := make([]{{$modelName}}, 0)
	modelInstance := &ModelInstance{
		Model:     &{{$modelName}}{},
		ModelList: &modelList,
	}
	return modelInstance

}

{{end}}
