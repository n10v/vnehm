// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tracksprocessor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
		DownloadFolder: config.Get("dlFolder"),
		ItunesPlaylist: config.Get("itunesPlaylist"),
	}
}

func (tp TracksProcessor) ProcessAll(tracks []track.Track) {
	if len(tracks) == 0 {
		ui.Term("there are no tracks to download", nil)
	}

	var errors []string
	// Start with last track
	for i := len(tracks) - 1; i >= 0; i-- {
		track := tracks[i]
		if err := tp.Process(track); err != nil {
			errors = append(errors, track.Fullname()+": "+err.Error())
			ui.Error("there was an error while downloading "+track.Fullname(), err)
		}
		ui.Newline()
	}

	if len(errors) > 0 {
		ui.Println(ui.RedString("There were errors while downloading tracks:"))
		for _, errText := range errors {
			ui.Println(ui.RedString("  " + errText))
		}
		ui.Newline()
	}

	ui.Success("Done!")
	ui.Quit()
}

func (tp TracksProcessor) Process(t track.Track) error {
	// Download track
	trackPath := filepath.Join(tp.DownloadFolder, t.Filename())
	if _, err := os.Create(trackPath); err != nil {
		return fmt.Errorf("couldn't create track file: %v", err)
	}
	if err := downloadTrack(t, trackPath); err != nil {
		return fmt.Errorf("couldn't download track: %v", err)
	}

	// Tag track
	if err := tag(t, trackPath); err != nil {
		return fmt.Errorf("coudln't tag file: %v", err)
	}

	// Add to iTunes
	if tp.ItunesPlaylist != "" {
		ui.Println("Adding to iTunes")
		if err := applescript.AddTrackToPlaylist(trackPath, tp.ItunesPlaylist); err != nil {
			return fmt.Errorf("couldn't add track to playlist: %v", err)
		}
	}
	return nil
}

func downloadTrack(t track.Track, path string) error {
	ui.Println("Downloading " + t.Artist() + " - " + t.Title())
	return runDownloadCmd(path, t.URL())
}

func runDownloadCmd(path, url string) error {
	cmd := exec.Command("curl", "-#", "-o", path, "-L", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func tag(t track.Track, trackPath string) error {
	tag, err := id3v2.Open(trackPath)
	if err != nil {
		return err
	}
	defer tag.Close()

	tag.SetArtist(t.Artist())
	tag.SetTitle(t.Title())

	return tag.Save()
}
