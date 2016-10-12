<div align="center">
<h1>vnehm</h1>
<p><b><i>vnehm</i></b> is a console tool, which downloads and adds to your iTunes library (if you use macOS) your <b>VK</b> audios in convenient way.</p>

<a href="https://raw.githubusercontent.com/bogem/vnehm/master/Pictures/list.png" target="_blank"><img src="https://raw.github.com/bogem/vnehm/master/Pictures/list.thumb.png" alt="List"></img></a>
<a href="https://raw.githubusercontent.com/bogem/vnehm/master/Pictures/get.png" target="_blank"><img src="https://raw.github.com/bogem/vnehm/master/Pictures/get.thumb.png" alt="List"></img></a>
<a href="https://raw.githubusercontent.com/bogem/vnehm/master/Pictures/search.png" target="_blank"><img src="https://raw.github.com/bogem/vnehm/master/Pictures/search.thumb.png" alt="List"></img></a>
<a href="https://raw.githubusercontent.com/bogem/vnehm/master/Pictures/help.png" target="_blank"><img src="https://raw.github.com/bogem/vnehm/master/Pictures/help.thumb.png" alt="List"></img></a>
<p><b>(click to zoom)</b></p>
</div>

---

<div align="center">
<h2>DISCLAIMER</h2>
<b><i><p>For personal use only</p>
vnehm developer doesn't responsible for any illegal usage of this program</i></b>
</div>

---

## Description
`vnehm` is a console tool written in `Go`. It can download your VK audios and add to iTunes, **if you use `macOS`**.

`vnehm` *wasn't tested on Windows machine, so it can be very buggy on it. I'll be very thankful, if you will report any bug.*

***If you have ideas to improve `vnehm`, issues and pull requests are always welcome! Also, if you have difficulties with installation/configuration/usage of `vnehm`, don't hesitate to write an issue. I will answer as soon as possible.***

## Installation
Install via `go` command:

	$ go get -u github.com/bogem/vnehm

or you can download and install binary from [latest release](https://github.com/bogem/vnehm/releases).

## Configuration
First of all, you should configure `vnehm`:

* **Create a file `.vnehmconfig` in your home directory**

* **Write in it configuration, i.e. set three variables in YAML format:**

`dlFolder` - filesystem path to download folder, where will be downloaded all tracks.
By default, your audios are being downloaded to your home directory

`itunesPlaylist` - name of iTunes playlist, where will be added all tracks *(if you're using `macOS`)*.
 By default, your audios are **not** being added to iTunes

Example:
```
dlFolder: /Users/bogem/Music
itunesPlaylist: iPod
```

* **Execute `vnehm auth` to authorize**

Don't worry if you become the message, what you mustn't copy the link from address bar. This app is only downloading and searching audios and it hasn't got any permissions to access your private data like messages, friends and etc. All permissions, which have vnehm, is access to audios. You can see it, when you authorize in your browser with the link, what give you vnehm for authorizing.

## Usage Examples

Type `vnehm help` to list of all available commands or `vnehm help COMMAND` for specific command.

Also commands may be abbreviated to one symbol length. For example, you can input `vnehm s` instead of `vnehm search`.

#### Get list of audios and download selected

	$ vnehm

#### Download last audio

	$ vnehm get

#### Download last 3 audios

	$ vnehm get 3

#### Download second audio and don't add it to iTunes playlist

	$ vnehm get -o 1 -i ''

#### Search for audios and download them

	$ vnehm search nasa

## FAQ

**Q: How can I add track to iTunes' music library, but not to any playlist?**

**A:** It depends on language you're using on your Mac. The name of your iTunes' music library you can see here:

![iTunes music master library](https://raw.github.com/bogem/vnehm/master/Pictures/music_master_library.png)

For example, english users should use `vnehm get -i Music`, russian users - `vnehm get -i Музыка`.

## TODO
- [ ] Make tests
- [ ] Upload to `homebrew`
- [ ] Use built-in downloader instead of `curl`

## License

MIT
