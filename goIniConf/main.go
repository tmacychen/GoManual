package main

import (
	"log"

	"github.com/Unknwon/goconfig"
	"os"
)

func main() {
	file,_ := os.Getwd()
	log.Println("current Path:",file)
	//当前在PyCharm的运行路径不是goconfig内
	cfg, err := goconfig.LoadConfigFile("goconfig/conf.ini")
	if err != nil {
		log.Fatal("unload the conf file")
	}
	value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "key_default")

	if err != nil {
		log.Fatal("unable to get value(%s):%s", "key_default", err)
	}
	log.Printf("%s > %s :%s ", goconfig.DEFAULT_SECTION, "key_default", value)

	isInsert := cfg.SetValue(goconfig.DEFAULT_SECTION, "key_default", "new value")
	if err != nil {
		log.Fatal("unable to set value(%s):%s", "key_default", err)
	}
	log.Printf("set the key :%s to value :%v ", "key_default", isInsert)

}
