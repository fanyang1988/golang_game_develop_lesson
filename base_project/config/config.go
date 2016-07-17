package config

import "github.com/BurntSushi/toml"

/*
配置文件
# an title
title = "base project"

[base]
ip = "127.0.0.1"
port = 3306

[redis]
address = ":6379"

[info]
usr = ["u1", "u2"]
*/

//BaseProjectCfg 总的config结构
type BaseProjectCfg struct {
	Title string   `toml:"title"`
	Base  BaseCfg  `toml:"base"`
	Redis RedisCfg `toml:"redis"`
	Info  InfoCfg  `toml:"info"`
}

type BaseCfg struct {
	IP   string `toml:"ip"`
	Port int    `toml:"port"`
}

type RedisCfg struct {
	Address string `toml:"address"`
}

type InfoCfg struct {
	Users []string `toml:"usr"`
}

// cfg 读取出来的cfg
var cfg BaseProjectCfg

// Get 获取cfg接口
func Get() BaseProjectCfg{
	return cfg
}

// LoadCfgFromFile 从指定文件中加载配置信息 返回是否出错
func LoadCfgFromFile(path string) error{
	_, err := toml.DecodeFile(path, &cfg)
	return err
}