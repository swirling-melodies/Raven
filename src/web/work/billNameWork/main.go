package main

import (
	"github.com/WFallenDown/Raven/src/web/work/billNameWork/billNameService"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	billNameService.SetBillName()
}