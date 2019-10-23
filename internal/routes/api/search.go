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

func SearchSongFromNetEase(c *gin.Context) {
	searchSong(c, netease.Client())
}

func SearchSongFromQQ(c *gin.Context) {
	searchSong(c, qq.Client())
}

func SearchSongFromMiGu(c *gin.Context) {
	searchSong(c, migu.Client())
}

func SearchSongFromKuGou(c *gin.Context) {
	searchSong(c, kugou.Client())
}

func SearchSongFromKuWo(c *gin.Context) {
	searchSong(c, kuwo.Client())
}

func searchSong(c *gin.Context, client provider.API) {
	keyword := strings.TrimSpace(c.Param("keyword"))
	data, err := client.SearchSong(keyword)
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
