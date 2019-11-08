package app

import (
	"autossh/src/utils"
)

func showVersion() {
	utils.Logln("autossh " + Version + " Build " + Build + "。")
	utils.Logln("由 YangYongli 编写，项目地址：https://github.com/yangyongli/autossh。")
}
