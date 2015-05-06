package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
  //  "github.com/hefju/CBChareServer/setting"
    "github.com/hefju/CBChareServer/models"
// "log"
    "time"
//  "github.com/hefju/LXServer/tools"
//  "strconv"
    "log"
   // "os"
   // "io"
    "fmt"
    "strconv"
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

//    logfile,err:=os.OpenFile("D:/w8/logs/cbserver.log",os.O_RDWR|os.O_CREATE,0)
//    if err!=nil{
//        log.Println("err")
//    }
//   log.SetOutput(logfile)

    router := gin.Default()
//    router.GET("/", func(c *gin.Context) {
//        log.Println("call index/n")
//        c.String(http.StatusOK, "CBChareServer...")
//    })
    router.GET("/", GetBill)//测试，获取数据表信息

    //获取服务器时间
    router.GET("/time", func(c *gin.Context) {
        c.String(http.StatusOK,time.Now().Format("2006-01-02 15:04:05"))
    })

    router.POST("/upload",uploaddata)
    router.POST("/uploadone",uploaddataone)


    router.Run(":8083")
}

func GetBill(c *gin.Context) {
    date:=time.Now().AddDate(0,-2,-1)
    log.Println("getbill date:",date)
    bill:=  models.GetChargeListByDate(date)
    c.JSON(200,bill)
}

func uploaddata(c *gin.Context) {
    var json []models.Tp_charge_billing
    c.Bind(&json)
    count:=  models.InsertBill(json)
    log.Println("插入结果:",count)
    //log.Println("call upload/n")
   // c.String(http.StatusOK,"from method uploaddata")
}


func uploaddataone(c *gin.Context) {
    var json models.Tp_charge_billing
    c.Bind(&json)
    fmt.Println(json)


    count:=  models.InsertBillOne(json)
    log.Println("uploaddataone插入结果:",count)
    c.String(http.StatusOK,"insert:"+strconv.FormatInt(count,10))//返回结果, 插入了多少条记录
    //log.Println("call upload/n")
    // c.String(http.StatusOK,"from method uploaddata")
}
