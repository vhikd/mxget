package migu

import (
	"errors"
	"fmt"
	"strings"

	"github.com/vhikd/mxget/pkg/provider"
	"github.com/winterssy/sreq"
)

func GetPlaylist(playlistId string) (*provider.Playlist, error) {
	return std.GetPlaylist(playlistId)
}

func (a *API) GetPlaylist(playlistId string) (*provider.Playlist, error) {
	resp, err := a.GetPlaylistRaw(playlistId)
	if err != nil {
		return nil, err
	}
	if len(resp.Resource) == 0 || len(resp.Resource[0].SongItems) == 0 {
		return nil, errors.New("get playlist: no data")
	}

	a.patchSongInfo(resp.Resource[0].SongItems...)
	a.patchSongURL(SongDefaultBR, resp.Resource[0].SongItems...)
	a.patchSongLyric(resp.Resource[0].SongItems...)
	songs := a.resolve(resp.Resource[0].SongItems)
	return &provider.Playlist{
		Name:   strings.TrimSpace(resp.Resource[0].Title),
		PicURL: resp.Resource[0].ImgItem.Img,
		Count:  len(songs),
		Songs:  songs,
	}, nil
}

func GetPlaylistRaw(playlistId string) (*PlaylistResponse, error) {
	return std.GetPlaylistRaw(playlistId)
}

// 获取歌单
func (a *API) GetPlaylistRaw(playlistId string) (*PlaylistResponse, error) {
	params := sreq.Params{
		"resourceId": playlistId,
	}

	resp := new(PlaylistResponse)
	err := a.Request(sreq.MethodGet, GetPlaylistAPI,
		sreq.WithQuery(params),
	).JSON(resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != "000000" {
		return nil, fmt.Errorf("get playlist: %s", resp.Info)
	}

	return resp, nil
}
