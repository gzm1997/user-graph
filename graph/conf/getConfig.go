package conf

import (
	//下面这个import必不可少 因为在这个import这个库的时候 它的init函数会注册这个类型的数据库 允许使用这个数据库进行存储
	_ "github.com/cayleygraph/cayley/graph/nosql/mongo"
	"github.com/astaxie/beego/config"
	"fmt"
)

var dbUrl string
var databaseName string

func init()  {
	conf, err := config.NewConfig("ini", "./graph/conf/data.yaml")
	if err != nil {
		mongodbIp := "120.92.100.60"
		mongodbPort := 27017
		cayleyDatabaseName := "testCayley"
		dbUrl = fmt.Sprintf("mongodb://%s:%s/%s", mongodbIp, mongodbPort, cayleyDatabaseName)
	}
	dbUrl = fmt.Sprintf("mongodb://%s:%s/%s", conf.String("mongodbIp"), conf.String("mongodbPort"), conf.String("cayleyDatabaseName"))
	databaseName = conf.String("databaseName")
}

func GetDbUrl() string {
	return dbUrl
}

func GetDataBaseName() string {
	return databaseName
}