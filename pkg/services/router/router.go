// pkg/services/router/router.go
package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter(host, port string, registerFunc func(*gin.Engine)) error {
	r := gin.Default()
	registerFunc(r)
	return r.Run(fmt.Sprintf("%s:%s", host, port))
}
