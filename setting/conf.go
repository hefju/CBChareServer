package setting

import (
	"github.com/donnie4w/go-logger/logger"
	"os"
)

func LoadConfig() {
	//logger.SetConsole(false)//默认是输出到控制台的, 所以logger.SetConsole(true) 写不写都无所谓
	os.MkdirAll("./log", 0777)                  //创建log文件夹, 用来存放日志
	logger.SetRollingDaily("./log", "test.log") //如果没有log文件夹, 需要新增文件夹
	logger.SetLevel(logger.DEBUG)

	//1.0
	//    flag.Parse()
	//    logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	//    if logErr != nil {
	//        log.Println("Fail to find", *logFile, "cServer start Failed")
	//        os.Exit(1)
	//    }
	//    log.SetOutput(logFile)
	//    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//
	//    //write log
	//    log.Println("Server abort! Cause:%v \r\n", "test log file")

	//2.创建日志文件
	//    t := time.Now()
	//    filepath := "./log_" + strings.Replace(t.String()[:10], ":", "_", 3) + ".txt"//:19 就输出到时分秒
	//    logfile, err := os.OpenFile(filepath, os.O_CREATE, 0666)
	//    if err != nil {
	//        log.Fatal("create log file failed!")
	//    }
	//    defer logfile.Close()
	//    log.SetOutput(logfile)
	//    log.Println("log.SetOutput")

}

/*获取当前文件执行的路径*/
//func GetCurrPath() string {
//    file, _ := exec.LookPath(os.Args[0])
//    path, _ := filepath.Abs(file)
//    splitstring := strings.Split(path, "\\")
//    size := len(splitstring)
//    splitstring = strings.Split(path, splitstring[size-1])
//    ret := strings.Replace(splitstring[0], "\\", "/", size-1)
//    return ret
//}
