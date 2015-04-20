package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
// "log"
    "time"
//  "github.com/hefju/LXServer/tools"
//  "strconv"
)
func main(){
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "CBChareServer...")
    })
    //获取服务器时间
    router.GET("/time", func(c *gin.Context) {
        c.String(http.StatusOK,time.Now().Format("2006-01-02 15:04:05"))
    })

    router.POST("/upload",uploaddata)


    router.Run(":8083")
}

func uploaddata(c *gin.Context) {
    c.String(http.StatusOK,"from method uploaddata")
}
