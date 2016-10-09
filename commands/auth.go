// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package commands

import (
	"net/url"
	"strings"

	"github.com/bogem/vnehm/client"
	"github.com/bogem/vnehm/config"
	"github.com/bogem/vnehm/ui"
	"github.com/spf13/cobra"
)

var (
	authCommand = &cobra.Command{
		Use:     "auth",
		Short:   "authorize in vnehm with your VK profile",
		Long:    "This command lets you authorize in vnehm with your VK profile. Without authorization vnehm won't work. Follow the instructions by execution of command.\n\nDon't worry if you become the message, what you mustn't copy the link from address bar. This app is only downloading and searching audios and it hasn't got any permissions to access your private data like messages, friends and etc. All permissions, which have vnehm, is access to audios. You can see it, when you authorize in your browser with the link, what give you vnehm for authorizing.",
		Aliases: []string{"a"},
		Run:     authenticate,
	}
)

func authenticate(cmd *cobra.Command, args []string) {
	showMessage()
	uri := ui.ReadInput()
	token := getToken(uri)
	config.Write("token", token)
	ui.Success("You're succesfully authorized!")
}

func showMessage() {
	ui.Println("1. Open this link in your browser: " + client.AuthURL())
	ui.Println("2. Click \"Allow\"")
	ui.Print("3. Copy link from address bar of your browser and paste here: ")
}

func getToken(uri string) string {
	hashIndex := strings.Index(uri, "#")
	fragment := uri[hashIndex+1:]
	values, err := url.ParseQuery(fragment)
	if err != nil {
		invalidURLMessage()
	}
	token := values.Get("access_token")
	if token == "" {
		invalidURLMessage()
	}
	return token
}

func invalidURLMessage() {
	ui.Error("You've entered invalind URL!", nil)
	ui.Error("Please execute `auth` command one more time and check input data.", nil)
	ui.Term("If you can't solve this problem, please write an issue on GitHub page: https://github.com/bogem/vnehm", nil)
}
