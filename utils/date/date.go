package date

import (
	"hh_blog/utils"
	"time"
)

// 获取当前YMD日期
func GetCurrentYmd() string {
	return time.Now().Format(utils.TimeYmdLayout)
}

// 时间戳转YMD日期
func TimeSecToYmd(sec int) string {
	ymd := time.Unix(int64(sec), 0).Format(utils.TimeYmdLayout)
	return ymd
}
