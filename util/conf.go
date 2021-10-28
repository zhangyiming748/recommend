package util

import (
	"github.com/widuu/goini"
)

const confPath = "./conf.ini"

var (
	RunMode string
	conf    *goini.Config
)

/**
 * 初始化
	init函数的主要作用：
		初始化不能采用初始化表达式初始化的变量。
		程序运行前的注册。
		实现sync.Once功能。
		其他
*/
func init() {
	initConfig()
}

func initConfig() {
	conf = goini.SetConfig(confPath)
	Infoln(confPath)
	RunMode = conf.GetValue("runmode", "mode")
}

/**
 * 获取环境变量
 */
func GetEnv() string {
	if RunMode == "" {
		initConfig()
	}
	return RunMode
}

/**
 * 根据键获取值
 */
func GetVal(section, name string) string {
	if section == "" {
		section = GetEnv()
	}
	val := conf.GetValue(section, name)
	return val
}
