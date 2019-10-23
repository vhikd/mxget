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

func GetPlaylistFromNetEase(c *gin.Context) {
	getPlaylist(c, netease.Client())
}

func GetPlaylistFromQQ(c *gin.Context) {
	getPlaylist(c, qq.Client())
}

func GetPlaylistFromMiGu(c *gin.Context) {
	getPlaylist(c, migu.Client())
}

func GetPlaylistFromKuGou(c *gin.Context) {
	getPlaylist(c, kugou.Client())
}

func GetPlaylistFromKuWo(c *gin.Context) {
	getPlaylist(c, kuwo.Client())
}

func getPlaylist(c *gin.Context, client provider.API) {
	id := strings.TrimSpace(c.Param("id"))
	data, err := client.GetPlaylist(id)
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
