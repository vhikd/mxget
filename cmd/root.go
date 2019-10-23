package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vhikd/mxget/cmd/album"
	"github.com/vhikd/mxget/cmd/artist"
	"github.com/vhikd/mxget/cmd/config"
	"github.com/vhikd/mxget/cmd/playlist"
	"github.com/vhikd/mxget/cmd/search"
	"github.com/vhikd/mxget/cmd/serve"
	"github.com/vhikd/mxget/cmd/song"
	"github.com/vhikd/mxget/internal/settings"
)

const (
	Version = "1.0.0"
)

var CmdRoot = &cobra.Command{
	Use:   "mxget",
	Short: "Show help for mxget commands.",
	Long: `
 _____ ______      ___    ___ ________  _______  _________   
|\   _ \  _   \   |\  \  /  /|\   ____\|\  ___ \|\___   ___\ 
\ \  \\\__\ \  \  \ \  \/  / | \  \___|\ \   __/\|___ \  \_| 
 \ \  \\|__| \  \  \ \    / / \ \  \  __\ \  \_|/__  \ \  \  
  \ \  \    \ \  \  /     \/   \ \  \|\  \ \  \_|\ \  \ \  \ 
   \ \__\    \ \__\/  /\   \    \ \_______\ \_______\  \ \__\
    \|__|     \|__/__/ /\ __\    \|_______|\|_______|   \|__|
                  |__|/ \|__|                                

A simple tool that help you download your favorite music, please
visit https://github.com/winterssy/mxget for more detail.
`,
	Version: Version,
}

func Execute() error {
	return CmdRoot.Execute()
}

func init() {
	cobra.OnInitialize(settings.Load)

	CmdRoot.AddCommand(search.CmdSearch)
	CmdRoot.AddCommand(song.CmdSong)
	CmdRoot.AddCommand(artist.CmdArtist)
	CmdRoot.AddCommand(album.CmdAlbum)
	CmdRoot.AddCommand(playlist.CmdPlaylist)
	CmdRoot.AddCommand(serve.CmdServe)
	CmdRoot.AddCommand(config.CmdSet)
}
