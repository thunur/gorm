package model

var ModelGenTemplate = `
package model

import (
	"time"
)

//这里是为了有的表没有引用time而导致的编译错误
var (
	_ = time.Second
)
// {{.StructName}} ...
type {{.StructName}} struct {
    {{range .Fields}}{{.}}
    {{end}}
}

// TableName ...
func ( {{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
`
