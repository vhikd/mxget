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

func GetArtistFromNetEase(c *gin.Context) {
	getArtist(c, netease.Client())
}

func GetArtistFromQQ(c *gin.Context) {
	getArtist(c, qq.Client())
}

func GetArtistFromMiGu(c *gin.Context) {
	getArtist(c, migu.Client())
}

func GetArtistFromKuGou(c *gin.Context) {
	getArtist(c, kugou.Client())
}

func GetArtistFromKuWo(c *gin.Context) {
	getArtist(c, kuwo.Client())
}

func getArtist(c *gin.Context, client provider.API) {
	id := strings.TrimSpace(c.Param("id"))
	data, err := client.GetArtist(id)
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
