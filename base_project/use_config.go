package main

import (
	"github.com/fanyang1988/golang_game_develop_lesson/base_project/config"
	"fmt"
)

/*
	2.使用toml配置文件
	我们使用
	https://github.com/BurntSushi/toml
	来解析toml文件
	toml是一种通用的配置文件格式，
	类似于windows上的ini，
	比ini强大的多，是一种简单强大的配置文件格式
	借助于toml库我们可以很方便的解析toml文件
*/

// useToml 使用toml，
// 为了其他模块可以读取配置信息，
// config结构是独立的包
func useToml(){
	err := config.LoadCfgFromFile(
		"conf/config.toml")

	if err != nil{
		fmt.Errorf(
			"load cfg err by %s\n",
			err.Error())
		panic(err)
	}

	fmt.Printf("config %v\n",
		config.Get())
}