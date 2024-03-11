package main

import (
	"fmt"
	_ "pro05shopping/docs"

	"time"

	"github.com/gin-gonic/gin"
)

// @title 电商项目
// @description 电商项目
// @version 1.0
// @contact.name 多课网
// @contact.url https://www.duoke360.com

// @host localhost:8080
// @BasePath /
// func main() {
// 	r := gin.Default()
// 	registerMiddlewares(r)
// 	api.RegisterHandlers(r)
// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
// 	srv := &http.Server{
// 		Addr:    ":8000",
// 		Handler: r,
// 	}

// 	go func() {
// 		// service connections
// 		if err := srv.ListenAndServe(); err != nil {
// 			log.Printf("listen: %s\n", err)
// 		}
// 	}()
// 	graceful.ShutdownGin(srv, time.Second*5)

// }

// 注册中间件
func registerMiddlewares(r *gin.Engine) {
	r.Use(
		gin.LoggerWithFormatter(
			func(param gin.LogFormatterParams) string {

				return fmt.Sprintf(
					"%s - [%s] \"%s %s %s %d %s %s\"\n",
					param.ClientIP,
					param.TimeStamp.Format(time.RFC3339),
					param.Method,
					param.Path,
					param.Request.Proto,
					param.StatusCode,
					param.Latency,
					param.ErrorMessage,
				)
			}))
	r.Use(gin.Recovery())
}
