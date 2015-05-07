package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/hefju/CBChareServer/setting"
    "github.com/hefju/CBChareServer/models"
    "time"
//  "github.com/hefju/LXServer/tools"
//  "strconv"
    "log"
  //  "os"
   // "io"
    "fmt"
    "strconv"
    "github.com/donnie4w/go-logger/logger"
)
func main(){


    //配置初始化
    setting.LoadConfig()
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
    router.GET("/", func(c *gin.Context) {//测试，获取数据表信息
       // log.Println("visit homepage\r\n")
        logger.Debug("visit homepage")
        c.String(http.StatusOK, "Welcome to CBChareServer...")
    })

    router.GET("/getbill", GetBill)//测试，获取昨天数据表信息

    //获取服务器时间
    router.GET("/time", func(c *gin.Context) {
        c.String(http.StatusOK,time.Now().Format("2006-01-02 15:04:05"))
    })

    router.POST("/upload",uploaddata)//批量上传数据
    router.POST("/uploadone",uploaddataone)//上传单条数据

    router.Run(":8083")
}

func GetBill(c *gin.Context) {
    date:=time.Now().AddDate(0,0,-1)
    bill:=  models.GetChargeListByDate(date)
    log.Println("getbill date:",date," length:",len(bill))
    c.JSON(200,bill)
}

//批量插入数据, 正式使用
func uploaddata(c *gin.Context) {
    var json []models.Tp_charge_billing2
    c.Bind(&json)
    count:=  models.InsertBill(json)
    logger.Debug("upload插入结果:",count)
    //log.Println("call upload/n")
    c.String(http.StatusOK,"insert:"+strconv.FormatInt(count,10))//返回结果, 插入了多少条记录
}

//测试使用插入单条数据
func uploaddataone(c *gin.Context) {
    var json models.Tp_charge_billing2
    c.Bind(&json)
    fmt.Println(json)

    count:=  models.InsertBillOne(json)
    log.Println("uploaddataone插入结果:",count)
    c.String(http.StatusOK,"insert:"+strconv.FormatInt(count,10))//返回结果, 插入了多少条记录
}
