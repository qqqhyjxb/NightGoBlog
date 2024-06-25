package utils

/*
读取配置信息

*/

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}
	LoadServer(file)
	LoadDatabase(file)
}

// LoadServer 读取Server配置信息
func LoadServer(file *ini.File) {
	// MustString方法:如果取不到对应的值，就设置为默认值
	AppMode = file.Section("server").Key("AppMode").MustString("debug")   //如果取不到默认值为debug
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000") //如果取不到默认值为3000
	JwtKey = file.Section("server").Key("JwtKey").MustString("f4da6sf4")
}

// LoadDatabase 读取Database配置信息
func LoadDatabase(file *ini.File) {
	// MustString方法:如果取不到对应的值，就设置为默认值
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("nightgo")
	DbPassword = file.Section("database").Key("DbPassword").MustString("wlq2002@")
	DbName = file.Section("database").Key("DbName").MustString("ngblog")
}
