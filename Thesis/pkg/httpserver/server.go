package httpserver


import "github.com/gin-gonic/gin"


func New() *gin.Engine {
r := gin.Default()
// ここで CORS やミドルウェアを追加してOK
return r
}