package main

import (
	"log"

	"github.com/Unknwon/goconfig"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
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
