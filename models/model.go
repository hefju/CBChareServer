package models

import (
    "time"
"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_"github.com/lunny/godbc"
	"log"
)

var engine *xorm.Engine
func init(){
	var err error
	//test addr
	engine, err = xorm.NewEngine("odbc", "driver={SQL Server};Server=192.168.1.200; Database=charge; uid=sa; pwd=123;")
	//engine, err = xorm.NewEngine("odbc", "driver={SQL Server};server=.;database=charge;integrated security=SSPI;")

	if err!=nil{
		log.Fatalln("xorm create error",err)
	}
    engine.ShowSQL = true
	engine.SetMapper(core.SameMapper{})
    // engine.CreateTables(new(tp_charge_billing))
	err = engine.Sync2(new(Tp_charge_billing2))//, new(Group))
	if err!=nil{
		log.Fatalln("xorm sync error",err)
	}
}

//客户端用的读取单日充电表数据
func GetChargeListByDate(date time.Time)[]Tp_charge_billing2{
    date1 := date.Format("2006-01-02 00:00:00")
    date = date.AddDate(0, 0, 1)
    date2 := date.Format("2006-01-02 00:00:00")

    bills := make([]Tp_charge_billing2, 0)
//  engine.Where("Crt_date>='?' and Crt_date<'?'",date1,date2).Find(&bills)
	engine.Where("Crt_date>=? and Crt_date<?",date1,date2).Find(&bills)
//log.Println("bills length:",len(bills))
    return bills
}

//插入单条数据, 测试用的
func InsertBillOne(bill Tp_charge_billing2)int64{
    affected, err := engine.Omit("Sn").Insert(bill)
    if err!=nil{
        log.Fatalln("insert bill",err)
    }
    return affected
}

//插入数组数据, 正式使用
func InsertBill(bills []Tp_charge_billing2)int64{
    affected, err := engine.Omit("Sn").Insert(bills)
    if err!=nil{
        log.Fatalln("insert bill",err)
    }
    return affected
}

//新建一个表Tp_charge_billing2, 把sn的自增属性去掉, 这样就能用数组批量插入. xorm的作者讲数组插入没有问题, 难度是我不会这是xomr的autoincr属性
//折中的解决办法是建立一个用于传输的Tp_charge_billing2表
type Tp_charge_billing2 struct {
	Sn             int  `xorm:"pk"`
	Chg_id         int
	Area_id        int
	Chg_port       int
	Chg_sn         string
	Opt_model      int
	Card_type      int
	Card_id        string
	Card_money     int
	Card_money_end int
	Chg_model      int
	Chg_para       int
	Charge         int
	Chg_pw         int
	Chg_time       int
	Pw_total       int
	Pw_total_end   int
	Soc_st         int
	Soc_ed         int
	Ch_st_date     string
	Ch_ed_date     string
	St_date        string
	Ed_date        string
	Ed_code        int
	Crt_date       time.Time
}
