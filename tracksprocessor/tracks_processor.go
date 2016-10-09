// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tracksprocessor

import (
	"os"
	"os/exec"
	"path"

	"github.com/bogem/id3v2"
	"github.com/bogem/vnehm/applescript"
	"github.com/bogem/vnehm/config"
	"github.com/bogem/vnehm/track"
	"github.com/bogem/vnehm/ui"
)

type TracksProcessor struct {
	DownloadFolder string // In this folder tracks will be downloaded
	ItunesPlaylist string // In this playlist tracks will be added
}

func NewConfiguredTracksProcessor() *TracksProcessor {
	return &TracksProcessor{
		DownloadFolder: config.GetDLFolder(),
		ItunesPlaylist: config.GetItunesPlaylist(),
	}
}

func (tp TracksProcessor) ProcessAll(tracks []track.Track) {
	if len(tracks) == 0 {
		ui.Term("There are no tracks to download", nil)
	}
	// Start with last track
	for i := len(tracks) - 1; i >= 0; i-- {
		track := tracks[i]
		tp.Process(track)
		ui.Newline()
	}
	ui.Success("Done!")
	ui.Quit()
}

func (tp TracksProcessor) Process(t track.Track) {
	// Download track
	trackPath := path.Join(tp.DownloadFolder, t.Filename())
	if _, err := os.OpenFile(trackPath, os.O_CREATE, 0766); err != nil {
		ui.Term("Couldn't create track file", err)
	}
	downloadTrack(t, trackPath)

	// Tag track
	tag(t, trackPath)

	// Add to iTunes
	if tp.ItunesPlaylist != "" {
		ui.Println("Adding to iTunes")
		applescript.AddTrackToPlaylist(trackPath, tp.ItunesPlaylist)
	}
}

func downloadTrack(t track.Track, path string) {
	ui.Println("Downloading " + t.Artist() + " - " + t.Title())
	runDownloadCmd(path, t.URL())
}

func runDownloadCmd(path, url string) {
	cmd := exec.Command("curl", "-#", "-o", path, "-L", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		ui.Term("Download failed", err)
	}
}

func tag(t track.Track, trackPath string) {
	tag, err := id3v2.Open(trackPath)
	if err != nil {
		ui.Term("Couldn't open track file", err)
	}
	defer tag.Close()

	tag.SetArtist(t.Artist())
	tag.SetTitle(t.Title())

	if err := tag.Save(); err != nil {
		ui.Term("Couldn't write tag to file", err)
	}
}
