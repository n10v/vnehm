// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"math"
	"net/url"
	"strconv"

	"github.com/bogem/vnehm/track"
	"github.com/tidwall/gjson"
)

const (
	tracksLimit = 5000 // vk count argument limit
)

// convRS2TS converts an array of gjson.Result to track.Track array.
func convRS2TS(results []gjson.Result) []track.Track {
	tracks := make([]track.Track, 0, len(results))
	for _, r := range results {
		tracks = append(tracks, convR2T(r))
	}
	return tracks
}

// convR2T converts a single gjson.Result to single track.Track object.
func convR2T(result gjson.Result) track.Track {
	return track.Track{
		JArtist:   result.Get("artist").String(),
		JDuration: result.Get("duration").Int(),
		JID:       result.Get("id").Int(),
		JTitle:    result.Get("title").String(),
		JURL:      result.Get("url").String(),
	}
}

func Audios(count, offset uint, token string) ([]track.Track, error) {
	requestsCount := float64(count) / float64(tracksLimit)
	requestsCount = math.Ceil(requestsCount)

	var limit uint
	var tracks []track.Track
	params := url.Values{}
	for i := uint(0); i < uint(requestsCount); i++ {
		if count < tracksLimit {
			limit = count
		} else {
			limit = tracksLimit
		}
		count -= limit

		params.Set("count", strconv.Itoa(int(limit)))
		params.Set("offset", strconv.Itoa(int((i*tracksLimit)+offset)))
		params.Set("access_token", token)
		params.Set("v", "5.60")

		jsonAudios, err := getTracks(params)
		if err == ErrNotFound {
			break
		}
		if err != nil {
			return nil, err
		}

		trackResults := gjson.GetBytes(jsonAudios, "response.items").Array()
		tracks = append(tracks, convRS2TS(trackResults)...)
	}

	return tracks, nil
}

func Search(limit, offset uint, query, token string) ([]track.Track, error) {
	params := url.Values{}
	params.Set("q", query)
	params.Set("count", strconv.Itoa(int(limit)))
	params.Set("offset", strconv.Itoa(int(offset)))
	params.Set("access_token", token)
	params.Set("v", "5.60")

	jsonFound, err := search(params)
	if err != nil {
		return nil, err
	}

	trackResults := gjson.GetBytes(jsonFound, "response.items").Array()
	return convRS2TS(trackResults), nil
}

const clientID = "5144754"

func AuthURL() string {
	params := url.Values{}
	params.Set("client_id", clientID)
	params.Set("redirect_uri", "https://oauth.vk.com/blank.html")
	params.Set("display", "page")
	params.Set("scope", "audio,offline")
	params.Set("response_type", "token")
	params.Set("v", "5.60")
	return "https://oauth.vk.com/authorize?" + params.Encode()
}

func WallAudios(id, token string) ([]track.Track, error) {
	params := url.Values{}
	params.Set("posts", id)
	params.Set("access_token", token)
	params.Set("v", "5.60")

	jsonWallAudios, err := wallAudios(params)
	if err != nil {
		return nil, err
	}

	var trackResults []gjson.Result
	attachmentsResults := gjson.GetBytes(jsonWallAudios, `response.#.attachments`)
	for _, r := range attachmentsResults.Array()[0].Array() { // it's so ugly!
		if r.Get("type").String() == "audio" {
			trackResults = append(trackResults, r.Get("audio"))
		}
	}
	return convRS2TS(trackResults), nil
}
