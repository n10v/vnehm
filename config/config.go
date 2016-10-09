// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Config is used for reading a config file and flags.
// Inspired from spf13/viper.
package config

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/bogem/vnehm/applescript"
	"github.com/bogem/vnehm/ui"
	"github.com/bogem/vnehm/util"
)

var (
	override = make(map[string]string)
	config   = make(map[string]string)

	configPath = path.Join(os.Getenv("HOME"), ".vnehmconfig")
	configRead bool
)

// Get has the behavior of returning the value associated with the first
// place from where it is set. Get will check value in the following order:
// flag, config file.
func Get(key string) string {
	if value, exists := override[key]; exists {
		return value
	}

	if !configRead {
		configRead = true
		read()
	}

	return config[key]
}

// read will discover and load the config file from disk.
func read() {
	configFile, err := os.Open(configPath)
	if os.IsNotExist(err) {
		ui.Error("There is no config file in your home directory.", nil)
		return
	}
	if err != nil {
		ui.Term("Couldn't open the config file", err)
	}

	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		ui.Term("Couldn't read the config file", err)
	}

	if err := yaml.Unmarshal(configBytes, config); err != nil {
		ui.Term("Couldn't unmarshal the config file", err)
	}
}

// GetPermalink returns the value associated with the key "dlFolder".
// If key "dlFolder" is blank in config, then it returns path to
// home directory.
func GetDLFolder() string {
	dlFolder := Get("dlFolder")
	if dlFolder == "" {
		ui.Warning("You didn't set a download folder. Tracks will be downloaded to your home directory.")
		return os.Getenv("HOME")
	}
	return util.SanitizePath(dlFolder)
}

// GetItunesPlaylist returns the value associated with
// the key "itunesPlaylist".
// If the OS of this computer isn't macOS, then it returns blank string.
func GetItunesPlaylist() string {
	playlist := ""
	if runtime.GOOS == "darwin" {
		playlist = Get("itunesPlaylist")

		if playlist == "" {
			ui.Warning("You didn't set an iTunes playlist. Tracks won't be added to iTunes.")
			return playlist
		}

		playlistsList := applescript.ListOfPlaylists()
		if !strings.Contains(playlistsList, playlist) {
			ui.Term("Playlist "+playlist+" doesn't exist. Please enter correct name.", nil)
		}
	}
	return playlist
}

// GetToken returns the value associated with the key "token".
func GetToken() string {
	token := Get("token")
	if token == "" {
		ui.Term("You aren't authorized. Please execute `vnehm auth` command to authorize.", nil)
	}
	return token
}

// Set sets the value for the key in the override regiser.
func Set(key, value string) {
	override[key] = value
}

// Write appends key and value to config file.
func Write(key, value string) {
	config[key] = value

	configFile, err := os.OpenFile(configPath, os.O_WRONLY, os.ModePerm)
	if os.IsNotExist(err) {
		configFile, err = os.Create(configPath)
		if err != nil {
			ui.Term("Couldn't create the config file", err)
		}
		err = nil
	}
	if err != nil {
		ui.Term("Couldn't open the config file", err)
	}
	defer configFile.Close()

	read()

	configBytes, err := yaml.Marshal(config)
	if err != nil {
		ui.Term("Coudn't marshal the config map", err)
	}

	_, err = configFile.Write(configBytes)
	if err != nil {
		ui.Term("Couldn't write to the config file", err)
	}
}
