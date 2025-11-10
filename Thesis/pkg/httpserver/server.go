package httpserver


import (
"github.com/gin-contrib/cors"
"github.com/gin-gonic/gin"
"time"
)


func New() *gin.Engine {
r := gin.Default()
r.Use(cors.New(cors.Config{
AllowOrigins: []string{"http://localhost:3000"},
AllowMethods: []string{"GET", "POST", "OPTIONS"},
AllowHeaders: []string{"Authorization", "Content-Type"},
ExposeHeaders: []string{"Content-Length"},
AllowCredentials: true,
MaxAge: 12 * time.Hour,
}))
return r
}