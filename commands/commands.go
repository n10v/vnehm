// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commands

import (
	"os"
	"runtime"
	"strings"

	"github.com/bogem/vnehm/applescript"
	"github.com/bogem/vnehm/config"
	"github.com/bogem/vnehm/ui"
	"github.com/bogem/vnehm/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var RootCmd = listCommand

// Variables used in flags
var (
	limit, offset                       uint
	dlFolder, itunesPlaylist, permalink string
)

func Execute() {
	RootCmd.AddCommand(authCommand)
	RootCmd.AddCommand(getCommand)
	RootCmd.AddCommand(searchCommand)
	RootCmd.AddCommand(versionCommand)
	RootCmd.Execute()
}

// addCommonFlags adds common flags related to download tracks.
func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&dlFolder, "dlFolder", "f", "", "filesystem path to download folder")

	if runtime.GOOS == "darwin" {
		cmd.Flags().StringVarP(&itunesPlaylist, "itunesPlaylist", "i", "", "name of iTunes playlist")
	}
}

func addLimitFlag(cmd *cobra.Command) {
	cmd.Flags().UintVarP(&limit, "limit", "l", 10, "count of tracks on each page")
}

func addOffsetFlag(cmd *cobra.Command) {
	cmd.Flags().UintVarP(&offset, "offset", "o", 0, "offset relative to first like")
}

// initializeConfig initializes a config with flags.
func initializeConfig(cmd *cobra.Command) {
	err := config.ReadInConfig()
	if err == config.ErrNotExist {
		ui.Warning("there is no config file. Read README to configure vnehm")
	} else if err != nil {
		ui.Term("", err)
	}
	checkToken()

	loadDefaultSettings()

	initializeDlFolder(cmd)
	initializeItunesPlaylist(cmd)
}

func checkToken() {
	if config.Get("token") == "" {
		ui.Term("you aren't authorized. Please execute `vnehm auth` command to authorize", nil)
	}
}

func loadDefaultSettings() {
	config.SetDefault("dlFolder", os.Getenv("HOME"))
	config.SetDefault("itunesPlaylist", "")
}

func flagChanged(fs *pflag.FlagSet, key string) bool {
	flag := fs.Lookup(key)
	if flag == nil {
		return false
	}
	return flag.Changed
}

// initializeDlFolder initializes dlFolder value. If there is no dlFolder
// set up, then dlFolder is set to HOME env variable.
func initializeDlFolder(cmd *cobra.Command) {
	var df string

	if flagChanged(cmd.Flags(), "dlFolder") {
		df = dlFolder
	} else {
		df = config.Get("dlFolder")
	}

	if df == "" {
		ui.Warning("you didn't set a download folder. Tracks will be downloaded to your home directory.")
		df = os.Getenv("HOME")
	}

	config.Set("dlFolder", util.SanitizePath(df))
}

// initializeItunesPlaylist initializes itunesPlaylist value. If there is no
// itunesPlaylist set up, then itunesPlaylist set up to blank string. Blank
// string is the sign, what tracks should not to be added to iTunes.
//
// initializeItunesPlaylist sets blank string to config, if OS is darwin
func initializeItunesPlaylist(cmd *cobra.Command) {
	var playlist string

	if runtime.GOOS == "darwin" {
		if flagChanged(cmd.Flags(), "itunesPlaylist") {
			playlist = itunesPlaylist
		} else {
			playlist = config.Get("itunesPlaylist")
		}

		if playlist == "" {
			ui.Warning("you didn't set an iTunes playlist. Tracks won't be added to iTunes.")
		} else {
			playlistsList, err := applescript.ListOfPlaylists()
			if err != nil {
				ui.Term("couldn't get list of playlists", err)
			}
			if !strings.Contains(playlistsList, playlist) {
				ui.Term("playlist "+playlist+" doesn't exist. Please enter correct name.", nil)
			}
		}
	}

	config.Set("itunesPlaylist", playlist)
}
