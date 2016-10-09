// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commands

import (
	"github.com/bogem/vnehm/client"
	"github.com/bogem/vnehm/config"
	"github.com/bogem/vnehm/track"
	"github.com/bogem/vnehm/tracksprocessor"
	"github.com/bogem/vnehm/ui"
	"github.com/spf13/cobra"
)

var (
	listCommand = &cobra.Command{
		Use:   "vnehm",
		Short: "list audios from your page, download and add them to iTunes",
		Long:  "vnehm is a console tool, which downloads (and adds to your iTunes library) your VK audios in convenient way",
		Run:   showListOfTracks,
	}
)

func init() {
	addCommonFlags(listCommand)
	addLimitFlag(listCommand)
	addOffsetFlag(listCommand)
}

func showListOfTracks(cmd *cobra.Command, args []string) {
	initializeConfig(cmd)

	tm := ui.TracksMenu{
		GetTracks: listGetTracks,
		Limit:     limit,
		Offset:    offset,
	}
	downloadTracks := tm.Show()

	tracksprocessor.NewConfiguredTracksProcessor().ProcessAll(downloadTracks)
}

func listGetTracks(offset uint) ([]track.Track, error) {
	return client.Audios(limit, offset, config.GetToken())
}
