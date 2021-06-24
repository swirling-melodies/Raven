package billModels

import (
	"Raven/src/log"
	"Raven/src/web/service"
	"encoding/json"
	. "github.com/ahmetb/go-linq/v3"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"os"
	"strconv"
)

type BillChartModel struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type IBillData interface {
	NewBillData()
	BillsInitDB()
	BillsWriteToJSON()
	BillsGetYearData()
}

type BillDataByDate struct {
	Data []BillDetail
	Year int `json:"Year" form:"Year"`
}

func (data *BillDataByDate) NewBillData() {

}

func (data *BillDataByDate) BillsInitDB() {
	billsInitDB()
}

func (data *BillDataByDate) BillsWriteToJSON() {
	var f *os.File
	src := strconv.Itoa(data.Year) + ".json"
	val, err := json.MarshalIndent(data.Data, "", "	") // 第二个表示每行的前缀，这里不用，第三个是缩进符号，这里用tab
	if err != nil {
		log.Writer(log.Error, err)
	}

	if service.CheckFileIsExist(src) { //如果文件存在
		f, err = os.OpenFile(src, os.O_APPEND, 0666) //打开文件
	} else {
		f, err = os.Create(src) //创建文件

	}

	err = ioutil.WriteFile(src, val, 0777)
	service.CheckErr(err)
	f.Close()
}

func (data *BillDataByDate) BillsGetYearData() {
	billsGetYearData(data)
}

func (data *BillDataByDate) BillsGetDataByMonth() {
	billsGetDataByMonth(data)
}

func BillsGetTable(bill *BillTable) *BillTable {
	billsInitDB()
	billsGetTable(bill)
	return bill
}

func BillsGetTableOption() *BillOption {
	billsInitDB()
	return billsGetTableOption()
}

func BillsGetDiagram(bill *BillTable) ([]BillChartModel, error) {
	var data []BillChartModel

	billsInitDB()
	billsGetDiagram(bill)

	From(bill.BillDetail).GroupBy(func(i interface{}) interface{} {
		return i.(BillDetail).BillName
	}, func(i interface{}) interface{} {
		return i.(BillDetail)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(BillDetail).Account
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return BillChartModel{i.Key.(string), m}
	}).ToSlice(&data)
	return data, nil
}
