package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vhikd/mxget/pkg/provider"
	"github.com/vhikd/mxget/pkg/provider/kugou"
	"github.com/vhikd/mxget/pkg/provider/kuwo"
	"github.com/vhikd/mxget/pkg/provider/migu"
	"github.com/vhikd/mxget/pkg/provider/netease"
	"github.com/vhikd/mxget/pkg/provider/qq"
)

func GetSongFromNetEase(c *gin.Context) {
	getSong(c, netease.Client())
}

func GetSongFromQQ(c *gin.Context) {
	getSong(c, qq.Client())
}

func GetSongFromMiGu(c *gin.Context) {
	getSong(c, migu.Client())
}

func GetSongFromKuGou(c *gin.Context) {
	getSong(c, kugou.Client())
}

func GetSongFromKuWo(c *gin.Context) {
	getSong(c, kuwo.Client())
}

func getSong(c *gin.Context, client provider.API) {
	id := strings.TrimSpace(c.Param("id"))
	data, err := client.GetSong(id)
	if err != nil {
		c.JSON(500, &provider.Response{
			Code:     500,
			Msg:      err.Error(),
			Platform: client.Platform(),
		})
		return
	}

	c.JSON(200, &provider.Response{
		Code:     200,
		Msg:      "ok",
		Data:     data,
		Platform: client.Platform(),
	})
}
