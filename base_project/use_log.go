package main

import log "github.com/cihub/seelog"

/*
	通过fmt可以输出一些程序信息，
	但是显然无法在正式环境中使用，
	log是后端开发中不可或缺的一部分，
	我们这里使用seelog来记录日志，
	大家可以在
	https://github.com/cihub/seelog
	阅读其wiki来学习其功能。

	seelog需要我们设置其运行模式，
	我们会通过配置文件来配置seelog：
	主要有：
	1. 日志级别
	2. 日志格式
	3. 输出目标
	4. 日志切分

*/

func useLogger() {
	err := initLog()
	if err != nil {
		panic(err)
	}

	// 需要在结束之前调用Flush以保证日志都写入文件
	defer log.Flush()
	log.Info("this is a info log %s", "sss")
	log.Error("this is an error log!")
	log.Debug("this is a debug log!")
}

func initLog() error {
	logger, err := log.LoggerFromConfigAsFile("./conf/log.xml")
	if err != nil {
		return err
	}
	log.ReplaceLogger(logger)
	return nil
}
