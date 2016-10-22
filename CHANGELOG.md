# vnehm change log

## 2.2
* Add errors stack after the downloading if there were errors
* Add "ERROR: " and "WARNING: " prefixes for errors and warnings respectively
* Show warnings in cases, if you didn't set dlFolder and itunesPlaylist
* vnehm will not colorize messages, if output isn't STDOUT
* Minor performance and stability improvements

## 2.1
* If you didn't set dlFolder, `vnehm` will download tracks to home directory implicitly
* If you didn't set itunesPlaylist, `vnehm` will not add tracks to iTunes implicitly
* If there is any error in config, you will be notified before the showing of
tracks menu
* Minor performance improvements

# 2.0
* **Rewritten in Go: faster, less memory usage, more stable, easier to install (no need to install `Taglib`)**
* `vnehm` is now only on English language
* `vnehm` command used to show list of audios, not help as earlier
* Removed `configure` command. Now you should configure `vnehm` in configure file (more in README)
* ... and some other improvements

# 1.1.3
* If you input in `to` option non-existing directory, you can create it from `vnehm`

# 1.1.2
* Fixed: `get` and `dl` don't download more than one track

# 1.1.1
* Minor bug fixes and improvements

# 1.1
* `select` command is renamed to `list`
* Fixed `uninitialized constant Vnehm::TrackManager::UserManager`

# 1.0
* First release!


