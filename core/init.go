package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"wheatwaves/config"
	"wheatwaves/global"

	"gopkg.in/yaml.v2"
)

// 初始化配置
func InitConf() {
	const File = "config.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(File) // 读取文件
	if err != nil {
		panic(fmt.Errorf("yanmlConf Error: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c) // Unmarshal将字符串转化为Config结构
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success")
	// fmt.Print(c)
	global.Config = c
}
