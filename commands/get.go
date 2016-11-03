// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commands

import (
	"strconv"
	"strings"

	"github.com/bogem/vnehm/client"
	"github.com/bogem/vnehm/config"
	"github.com/bogem/vnehm/track"
	"github.com/bogem/vnehm/tracksprocessor"
	"github.com/bogem/vnehm/ui"
	"github.com/spf13/cobra"
)

var (
	getCommand = &cobra.Command{
		Use:     "get [number or url]",
		Short:   "download inputed count of likes or audios from the wall post (and add to your iTunes library)",
		Aliases: []string{"g"},
		Run:     getTracks,
	}
)

func init() {
	addCommonFlags(getCommand)
	addOffsetFlag(getCommand)
}

func getTracks(cmd *cobra.Command, args []string) {
	initializeConfig(cmd)

	tp := tracksprocessor.NewConfiguredTracksProcessor()

	var arg string
	if len(args) == 0 {
		arg = "1"
	} else {
		arg = args[0]
	}

	var downloadTracks []track.Track
	var err error
	if isVKWallURL(arg) {
		downloadTracks, err = getTracksFromVKWall(arg)
	} else if num, err := strconv.Atoi(arg); err == nil {
		downloadTracks, err = getLastTracks(uint(num))
	} else {
		ui.Term("you've entered invalid argument. Run 'vnehm get --help' for usage.", nil)
	}

	if err != nil {
		handleError(err)
	}

	tp.ProcessAll(downloadTracks)
}

func isVKWallURL(uri string) bool {
	return strings.Contains(uri, "vk.com/wall")
}

// getTracksFromVKWall returns an id of post. It knows, what uri is
// correct VK link to post.
func getTracksFromVKWall(uri string) ([]track.Track, error) {
	ui.Println("Downloading audio(s) from the post")
	index := strings.Index(uri, "wall")
	postID := uri[index+len("wall"):]
	return client.WallAudios(postID, config.Get("token"))
}

func getLastTracks(count uint) ([]track.Track, error) {
	return client.Audios(count, offset, config.Get("token"))
}

func handleError(err error) {
	switch {
	case strings.Contains(err.Error(), "403"):
		ui.Term("you're not allowed to see these tracks", nil)
	case strings.Contains(err.Error(), "404"):
		ui.Term("there are no tracks", nil)
	}
}
