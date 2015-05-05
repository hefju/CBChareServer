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
	engine, err = xorm.NewEngine("odbc", "driver={SQL Server};Server=192.168.1.200; Database=charge; uid=sa; pwd=123;")
	if err!=nil{
		log.Fatalln("xorm create error",err)
	}
    engine.ShowSQL = true
	engine.SetMapper(core.SameMapper{})
    // engine.CreateTables(new(tp_charge_billing))
	err = engine.Sync2(new(Tp_charge_billing))//, new(Group))
	if err!=nil{
		log.Fatalln("xorm sync error",err)
	}
}
func GetBilling(date time.Time)[]Tp_charge_billing{
//    date1 := date.Format("2006-01-02 00:00:00")
//    date = date.AddDate(0, 0, 1)
//    date2 := date.Format("2006-01-02 00:00:00")

    bills := make([]Tp_charge_billing, 0)
//    err := engine.Find(&everyone)
//    engine.Where("Crt_date>='?' and Crt_date<'?'",date1,date2).Find(&bills)
	engine.Limit(0,10).Find(&bills)
log.Println("bills length:",len(bills))
    return bills
}
func InsertBill(bill Tp_charge_billing)int64{
    affected, err := engine.Insert(bill)
    if err!=nil{
        log.Fatalln("insert bill",err)
    }
    return affected
}


type Tp_charge_billing struct {
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
