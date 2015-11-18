require 'vnehm/applescript'
require 'vnehm/playlist'

module Vnehm

  ##
  # Playlist manager works with iTunes playlists

  module PlaylistManager

    def self.default_playlist
      default_user_playlist || music_master_library if OS.mac?
    end

    ##
    # Checks path for existence and returns it if exists

    def self.get_playlist(playlist_name)
      if AppleScript.list_of_playlists.include? playlist_name
        Playlist.new(playlist_name)
      else
        UI.term 'Такого плейлиста не существует. Введите корректное название'
      end
    end

    def self.set_playlist
      loop do
        playlist = UI.ask('Введите имя плейлиста iTunes, в который вы ' \
                          'хотите добавлять треки (нажмите Enter, чтобы ' \
                          'не добавлять аудиозаписи в плейлист, ' \
                          'а в музыкальную медиатеку iTunes)')

        # If entered nothing, unset iTunes playlist
        if playlist == ''
          Cfg[:playlist] = nil
          UI.success 'Треки будут добавляться в музыкальную медиатеку iTunes'
          break
        end

        if AppleScript.list_of_playlists.include? playlist
          Cfg[:playlist] = playlist
          UI.say 'Аудиозаписи будут добавляться в плейлист '.green +
                                                              playlist.magenta
          break
        else
          UI.error 'Такого плейлиста не существует. Введите корректное название'
        end
      end
    end


    module_function

    def default_user_playlist
      Playlist.new(Cfg[:playlist]) unless Cfg[:playlist].nil?
    end

    ##
    # Music master library is main iTunes music library

    def music_master_library
      Playlist.new(AppleScript.music_master_library)
    end

  end
end
