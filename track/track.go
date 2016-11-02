// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package track

import (
	"html"
	"strings"

	"github.com/bogem/vnehm/util"
)

type Track struct {
	artist string
	title  string

	// Properties needed for JSON unmarshalling
	JArtist   string
	JDuration int64
	JID       int64
	JTitle    string
	JURL      string
}

func (t *Track) Artist() string {
	if t.artist == "" {
		t.artist, t.title = t.name()
	}
	return t.artist
}

func (t Track) Duration() string {
	return util.DurationString(util.ParseDuration(int(t.JDuration)))
}

func (t Track) Filename() string {
	// Replace all filesystem non-friendly runes with the underscore
	toReplace := "/\\"
	replaceRunes := func(r rune) rune {
		if strings.ContainsRune(toReplace, r) {
			return '_'
		}
		return r
	}

	return strings.Map(replaceRunes, t.Fullname()) + ".mp3"
}

func (t Track) Fullname() string {
	return t.Artist() + " – " + t.Title()
}

func (t Track) ID() int64 {
	return t.JID
}

// name splits track's title to artist and title if there is one of separators
// in there and unescape them.
// E.g. if track has title "Michael Jackson - Thriller" then this function will
// return as first string "Michael Jackson" and as second string "Thriller".
func (t Track) name() (string, string) {
	artist := t.JArtist
	title := t.JTitle
	separators := [...]string{" - ", " ~ ", " – "}
	for _, sep := range separators {
		if strings.Contains(t.JTitle, sep) {
			splitted := strings.SplitN(t.JTitle, sep, 2)
			artist = splitted[0]
			title = splitted[1]
		}
	}
	return html.UnescapeString(artist), html.UnescapeString(title)
}

func (t *Track) Title() string {
	if t.title == "" {
		t.artist, t.title = t.name()
	}
	return t.title
}

func (t Track) URL() string {
	return t.JURL
}
