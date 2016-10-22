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
		Short:   "download inputed count of likes (and add to your iTunes library)",
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
	if num, err := strconv.Atoi(arg); err == nil {
		downloadTracks, err = getLastTracks(uint(num))
		if err != nil {
			handleError(err)
		}
	} else {
		ui.Term("You've entered invalid argument. Run 'vnehm get --help' for usage.", nil)
	}

	tp.ProcessAll(downloadTracks)
}

func getLastTracks(count uint) ([]track.Track, error) {
	return client.Audios(count, offset, config.Get("token"))
}

func handleError(err error) {
	switch {
	case strings.Contains(err.Error(), "403"):
		ui.Term("You're not allowed to see these tracks", nil)
	case strings.Contains(err.Error(), "404"):
		ui.Term("There are no tracks", nil)
	}
}
