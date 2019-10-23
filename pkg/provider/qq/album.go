package qq

import (
	"errors"
	"fmt"
	"strings"

	"github.com/vhikd/mxget/pkg/provider"
	"github.com/winterssy/sreq"
)

func GetAlbum(albumMid string) (*provider.Album, error) {
	return std.GetAlbum(albumMid)
}

func (a *API) GetAlbum(albumMid string) (*provider.Album, error) {
	resp, err := a.GetAlbumRaw(albumMid)
	if err != nil {
		return nil, err
	}

	n := len(resp.Data.GetSongInfo)
	if n == 0 {
		return nil, errors.New("get album: no data")
	}

	_songs := resp.Data.GetSongInfo
	a.patchSongInfo(_songs...)
	a.patchSongURL(_songs...)
	a.patchSongLyric(_songs...)
	a.patchAlbumInfo(resp)
	songs := a.resolve(_songs)
	return &provider.Album{
		Name:   strings.TrimSpace(resp.Data.GetAlbumInfo.FAlbumName),
		PicURL: resp.Data.GetAlbumInfo.PicURL,
		Count:  n,
		Songs:  songs,
	}, nil
}

func GetAlbumRaw(albumMid string) (*AlbumResponse, error) {
	return std.GetAlbumRaw(albumMid)
}

// 获取专辑
func (a *API) GetAlbumRaw(albumMid string) (*AlbumResponse, error) {
	params := sreq.Params{
		"albummid": albumMid,
	}

	resp := new(AlbumResponse)
	err := a.Request(sreq.MethodGet, GetAlbumAPI,
		sreq.WithQuery(params),
	).JSON(resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("get album: %d", resp.Code)
	}

	return resp, nil
}
