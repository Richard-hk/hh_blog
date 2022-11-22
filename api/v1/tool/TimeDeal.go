package tool

import (
	"hh_blog/utils"
	"hh_blog/utils/date"
	"hh_blog/utils/rtn"
	"hh_blog/utils/validators"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UnixSec struct {
	StartUnix int    `json:"start_unix" validate:"gt=0" label:"开始时间"`
	EndUnix   int    `json:"end_unix" validate:"gt=0" label:"结束时间"`
	Glue      string `json:"glue" validate:"max=50" label:"分割字符"`
}

func TimeUnixSecDeal(c *gin.Context) {
	var unixSec UnixSec
	err := c.BindJSON(&unixSec)
	if err != nil {
		c.JSON(http.StatusOK, rtn.ReturnJson(http.StatusBadRequest, "", err))
		return
	}
	errS := validators.ValidateStructLabel(unixSec)
	if errS != nil {
		c.JSON(http.StatusOK, rtn.ReturnJson(http.StatusBadRequest, "", errS))
		return
	}
	data := buildData(unixSec)
	c.JSON(http.StatusOK, rtn.ReturnJson(http.StatusOK, data, ""))
}

// 构建数据
func buildData(unixSec UnixSec) string {
	data := date.TimeSecToYmd(unixSec.StartUnix)
	startUnix := unixSec.StartUnix + utils.Day
	endUnix := unixSec.EndUnix
	glue := strings.Replace(unixSec.Glue, "\\n", "\n", -1)
	for currUnix := startUnix; currUnix <= endUnix; currUnix += utils.Day {
		data += glue + date.TimeSecToYmd(currUnix)
	}
	return data
}
