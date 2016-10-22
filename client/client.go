// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"encoding/json"
	"math"
	"net/url"
	"strconv"

	"github.com/bogem/vnehm/track"
	"github.com/bogem/vnehm/ui"
)

const (
	tracksLimit = 150
)

type responseObject struct {
	Response response `json:"response"`
}

type response struct {
	Tracks []track.Track `json:"items"`
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
		params.Set("v", "5.57")

		bAudios, err := getTracks(params)
		if err == ErrNotFound {
			break
		}
		if err != nil {
			return nil, err
		}

		resp := new(responseObject)
		if err := json.Unmarshal(bAudios, &resp); err != nil {
			ui.Term("could't unmarshal JSON with likes", err)
		}
		tracks = append(tracks, resp.Response.Tracks...)
	}
	return tracks, nil
}

func Search(limit, offset uint, query, token string) ([]track.Track, error) {
	params := url.Values{}
	params.Set("q", query)
	params.Set("count", strconv.Itoa(int(limit)))
	params.Set("offset", strconv.Itoa(int(offset)))
	params.Set("access_token", token)
	params.Set("v", "5.57")

	bFound, err := search(params)
	if err != nil {
		return nil, err
	}

	resp := new(responseObject)
	if err := json.Unmarshal(bFound, &resp); err != nil {
		ui.Term("couldn't unmarshal JSON with search results", err)
	}

	return resp.Response.Tracks, nil
}

const clientID = "5144754"

func AuthURL() string {
	params := url.Values{}
	params.Set("client_id", clientID)
	params.Set("redirect_uri", "https://oauth.vk.com/blank.html")
	params.Set("display", "page")
	params.Set("scope", "audio,offline")
	params.Set("response_type", "token")
	params.Set("v", "5.57")
	return "https://oauth.vk.com/authorize?" + params.Encode()
}
