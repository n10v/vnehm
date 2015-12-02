require 'taglib'

require 'vnehm/track'

module Vnehm

  class TrackManager

    def initialize(options)
      setup_environment(options)
    end

    def process_tracks(tracks)
      tracks.reverse_each do |track|
        dl(track)
        tag(track)
        @playlist.add_track(track) if @playlist
        UI.newline
      end
      UI.success 'Готово!'
    end

    def tracks(limit, offset)
      UI.term "Неверное значение опции 'limit'\nОно должно быть больше 0" if limit <= 0
      UI.term "Неверное значении опции 'offset'\nОно должно быть больше или равно 0" if offset < 0

      tracks = Client.tracks(limit, offset)
      return nil if tracks.empty?

      tracks.map! { |hash| Track.new(hash) }
    end

    def search(query, limit, offset)
      UI.term "Неверное значение опции 'limit'\nОно должно быть больше 0" if limit <= 0
      UI.term "Неверное значении опции 'offset'\nОно должно быть больше или равно 0" if offset < 0

      found = Client.search(query, limit, offset)
      return nil if found.empty?

      found.shift # Deleting first item - 'count'

      found.map! { |hash| Track.new(hash) }
    end

    private

    def dl(track)
      # Downloading track
      UI.say 'Загрузка ' + track.full_name
      `curl -# -o \"#{track.file_path}\" -L \"#{track.url}\"`
    end

    def tag(track)
      TagLib::MPEG::File.open(track.file_path) do |file|
        tag = file.id3v2_tag
        tag.artist = track.artist
        tag.title = track.title

        file.save
      end
    end

    def setup_environment(options)
      # Setting up download path
      temp_path = options[:to]
      dl_path = temp_path ? PathManager.get_path(temp_path) : PathManager.default_dl_path
      if dl_path
        ENV['dl_path'] = dl_path
      else
        UI.error "Вы не указали путь, куда должны скачиваться треки!"
        UI.say "Укажите его через команду  #{'vnehm configure'.yellow} " \
        " или используйте опцию #{'to ПУТЬ'.yellow}"
        UI.term
      end

      # Setting up iTunes playlist
      @playlist = nil
      if options[:dl] != 'yes' && OS.mac?
        playlist_name = options[:pl]
        @playlist = playlist_name ? PlaylistManager.get_playlist(playlist_name) : PlaylistManager.default_playlist
      end
    end

  end
end
