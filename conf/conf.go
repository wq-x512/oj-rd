package conf

import (
	"fmt"

	"github.com/go-ini/ini"
)

var (
	AppMode     string
	HttpPort    string
	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbName      string
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
	MongoDBName string
	MongoDBAddr string
	MongoDBPwd  string
	MongoDBPort string
)

func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("ini load error")
	}
	loadServer(file)
	loadMysql(file)
	loadMongoDB(file)
}

func loadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()

}
func loadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
func loadMongoDB(file *ini.File) {
	MongoDBName = file.Section("MongoDB").Key("MongoDBName").String()
	MongoDBAddr = file.Section("MongoDB").Key("MongoDBAddr").String()
	MongoDBPwd = file.Section("MongoDB").Key("MongoDBPwd").String()
	MongoDBPort = file.Section("MongoDB").Key("MongoDBPort").String()
}
