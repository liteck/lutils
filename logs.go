package lutils

import "lutils/logs"

func INFO(i interface{}) {
	logs.INFO(i)
}

func DEBUG(i interface{}) {
	logs.DEBUG(i)
}

func NOTICE(i interface{}) {
	logs.NOTICE(i)
}

func WARNING(i interface{}) {
	logs.WARNING(i)
}

func ERROR(i interface{}) {
	logs.ERROR(i)
}

func CRITICAL(i interface{}) {
	logs.CRITICAL(i)
}

func ALERT(i interface{}) {
	logs.ALERT(i)
}
