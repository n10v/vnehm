// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commands

import (
	"strings"

	"github.com/bogem/vnehm/client"
	"github.com/bogem/vnehm/config"
	"github.com/bogem/vnehm/track"
	"github.com/bogem/vnehm/tracksprocessor"
	"github.com/bogem/vnehm/ui"
	"github.com/spf13/cobra"
)

var (
	searchCommand = &cobra.Command{
		Use:     "search [query]",
		Short:   "search tracks, show them, download (and add to iTunes)",
		Aliases: []string{"s"},
		Run:     searchAndShowTracks,
	}

	searchQuery string
)

func init() {
	addCommonFlags(searchCommand)
	addLimitFlag(searchCommand)
}

func searchAndShowTracks(cmd *cobra.Command, args []string) {
	initializeConfig(cmd)

	tp := tracksprocessor.NewConfiguredTracksProcessor()

	searchQuery = strings.Join(args, " ")

	tm := ui.TracksMenu{
		GetTracks: searchGetTracks,
		Limit:     limit,
	}
	downloadTracks := tm.Show()

	tp.ProcessAll(downloadTracks)
}

func searchGetTracks(offset uint) ([]track.Track, error) {
	return client.Search(limit, offset, searchQuery, config.Get("token"))
}
