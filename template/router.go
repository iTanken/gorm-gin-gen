package template

var RouterTmpl = `package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ConfigRouter(router *gin.RouterGroup) {
    {{range .}}config{{pluralize .}}Router(router)
    {{end}}
}
`
