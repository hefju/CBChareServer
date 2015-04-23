package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
  //  "github.com/hefju/CBChareServer/setting"
// "log"
    "time"
//  "github.com/hefju/LXServer/tools"
//  "strconv"
    "log"
    "os"
   // "io"
)
func main(){

    //配置初始化
//    setting.LoadConfig()
//    os.MkdirAll("D:/w8/logs",os.ModePerm)
//    f,err:=os.Create("D:/w8/logs/cbserver.log")
//    if err!=nil{
//        log.Panic("create log file faild:",err)
//    }
//    w:=io.MultiWriter(f,os.Stdout)
//    log.SetOutput(w)
    logfile,err:=os.OpenFile("D:/w8/logs/cbserver.log",os.O_RDWR|os.O_CREATE,0)
    if err!=nil{
        log.Println("err")
    }
   // log.Println(logfile)
   log.SetOutput(logfile)

    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        log.Println("call index/n")
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
    log.Println("call upload/n")
    c.String(http.StatusOK,"from method uploaddata")
}
